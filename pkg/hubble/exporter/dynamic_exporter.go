// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package exporter

import (
	"context"
	"errors"

	"github.com/sirupsen/logrus"

	v1 "github.com/cilium/cilium/pkg/hubble/api/v1"
	"github.com/cilium/cilium/pkg/hubble/exporter/exporteroption"
	"github.com/cilium/cilium/pkg/lock"
	"github.com/cilium/cilium/pkg/time"
)

var reloadInterval = 5 * time.Second

// DynamicExporter is a wrapper of the hubble exporter that supports dynamic configuration reload
// for a set of exporters.
type DynamicExporter struct {
	FlowLogExporter

	logger        logrus.FieldLogger
	watcher       *configWatcher
	maxFileSizeMB int
	maxBackups    int

	// mutex protects from concurrent modification of managedExporters by config
	// reloader when hubble events are processed
	mutex            lock.RWMutex
	managedExporters map[string]*managedExporter
}

// NewDynamicExporter initializes a dynamic exporter.
//
// The actual config watching must be started by invoking watch().
//
// NOTE: Stopped instances cannot be restarted and should be re-created.
func NewDynamicExporter(logger logrus.FieldLogger, configFilePath string, maxFileSizeMB, maxBackups int) *DynamicExporter {
	dynamicExporter := &DynamicExporter{
		logger:           logger,
		managedExporters: make(map[string]*managedExporter),
		maxFileSizeMB:    maxFileSizeMB,
		maxBackups:       maxBackups,
	}

	registerMetrics(dynamicExporter)

	watcher := NewConfigWatcher(configFilePath, dynamicExporter.onConfigReload)
	dynamicExporter.watcher = watcher
	return dynamicExporter
}

// Watch starts watching the exporter configuration file at regular intervals and initiate a reload
// whenever the config changes. It blocks until the context is cancelled.
func (d *DynamicExporter) Watch(ctx context.Context) error {
	return d.watcher.watch(ctx, reloadInterval)
}

// OnDecodedEvent implemebnts FlowLogExporter.
// It distributes events across all managed exporters.
func (d *DynamicExporter) OnDecodedEvent(ctx context.Context, event *v1.Event) (bool, error) {
	select {
	case <-ctx.Done():
		return false, nil
	default:
	}

	d.mutex.RLock()
	defer d.mutex.RUnlock()

	var errs error
	for _, me := range d.managedExporters {
		if me.config.End == nil || me.config.End.After(time.Now()) {
			_, err := me.exporter.OnDecodedEvent(ctx, event)
			errs = errors.Join(errs, err)
		}
	}
	return false, errs
}

// Stop implements FlowLogExporter.
// It stops all managed flow log exporters.
func (d *DynamicExporter) Stop() error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	var errs error
	for _, me := range d.managedExporters {
		errs = errors.Join(errs, me.exporter.Stop())
	}

	return errs
}

func (d *DynamicExporter) onConfigReload(hash uint64, config DynamicExportersConfig) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	configuredFlowLogNames := make(map[string]interface{})
	for _, flowlog := range config.FlowLogs {
		configuredFlowLogNames[flowlog.Name] = struct{}{}
		var label string
		if _, ok := d.managedExporters[flowlog.Name]; ok {
			label = "update"
		} else {
			label = "add"
		}
		if d.applyUpdatedConfig(flowlog) {
			DynamicExporterReconfigurations.WithLabelValues(label).Inc()
		}
	}

	for flowLogName := range d.managedExporters {
		if _, ok := configuredFlowLogNames[flowLogName]; !ok {
			if d.removeExporter(flowLogName) {
				DynamicExporterReconfigurations.WithLabelValues("remove").Inc()
			}
		}
	}

	d.updateLastAppliedConfigGauges(hash)
}

func (d *DynamicExporter) newExporter(flowlog *FlowLogConfig) (*exporter, error) {
	exporterOpts := []exporteroption.Option{
		exporteroption.WithPath(flowlog.FilePath),
		exporteroption.WithMaxSizeMB(d.maxFileSizeMB),
		exporteroption.WithMaxBackups(d.maxBackups),
		exporteroption.WithAllowList(d.logger, flowlog.IncludeFilters),
		exporteroption.WithDenyList(d.logger, flowlog.ExcludeFilters),
		exporteroption.WithFieldMask(flowlog.FieldMask),
	}

	return NewExporter(d.logger.WithField("flowLogName", flowlog.Name), exporterOpts...)
}

func (d *DynamicExporter) applyUpdatedConfig(flowlog *FlowLogConfig) bool {
	m, ok := d.managedExporters[flowlog.Name]
	if ok && m.config.equals(flowlog) {
		return false
	}

	exporter, err := d.newExporter(flowlog)
	if err != nil {
		d.logger.Errorf("Failed to apply flowlog for name %s: %v", flowlog.Name, err)
		return false
	}

	d.removeExporter(flowlog.Name)
	d.managedExporters[flowlog.Name] = &managedExporter{
		config:   flowlog,
		exporter: exporter,
	}
	return true
}

func (d *DynamicExporter) removeExporter(name string) bool {
	m, ok := d.managedExporters[name]
	if !ok {
		return false
	}
	if err := m.exporter.Stop(); err != nil {
		d.logger.Errorf("failed to stop exporter: %w", err)
	}
	delete(d.managedExporters, name)
	return true
}

func (d *DynamicExporter) updateLastAppliedConfigGauges(hash uint64) {
	DynamicExporterConfigHash.WithLabelValues().Set(float64(hash))
	DynamicExporterConfigLastApplied.WithLabelValues().SetToCurrentTime()
}

type managedExporter struct {
	config   *FlowLogConfig
	exporter FlowLogExporter
}
