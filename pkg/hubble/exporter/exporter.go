// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package exporter

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/cilium/lumberjack/v2"
	"github.com/sirupsen/logrus"

	flowpb "github.com/cilium/cilium/api/v1/flow"
	observerpb "github.com/cilium/cilium/api/v1/observer"
	v1 "github.com/cilium/cilium/pkg/hubble/api/v1"
	"github.com/cilium/cilium/pkg/hubble/exporter/exporteroption"
	"github.com/cilium/cilium/pkg/hubble/filters"
	nodeTypes "github.com/cilium/cilium/pkg/node/types"
)

// exporter is an implementation of OnDecodedEvent interface that writes Hubble events to a file.
type exporter struct {
	FlowLogExporter

	logger  logrus.FieldLogger
	encoder *json.Encoder
	writer  io.WriteCloser
	flow    *flowpb.Flow

	opts exporteroption.Options
}

// NewExporter initializes an exporter.
// NOTE: Stopped instances cannot be restarted and should be re-created.
func NewExporter(logger logrus.FieldLogger, options ...exporteroption.Option) (*exporter, error) {
	opts := exporteroption.Default // start with defaults
	for _, opt := range options {
		if err := opt(&opts); err != nil {
			return nil, fmt.Errorf("failed to apply option: %w", err)
		}
	}
	logger.WithField("options", opts).Info("Configuring Hubble event exporter")
	var writer io.WriteCloser
	// If hubble-export-file-path is set to "stdout", use os.Stdout as the writer.
	if opts.Path == "stdout" {
		writer = &noopWriteCloser{w: os.Stdout}
	} else {
		writer = &lumberjack.Logger{
			Filename:   opts.Path,
			MaxSize:    opts.MaxSizeMB,
			MaxBackups: opts.MaxBackups,
			Compress:   opts.Compress,
		}
	}
	return newExporter(logger, writer, opts)
}

// newExporter let's you supply your own WriteCloser for tests.
func newExporter(logger logrus.FieldLogger, writer io.WriteCloser, opts exporteroption.Options) (*exporter, error) {
	encoder := json.NewEncoder(writer)
	var flow *flowpb.Flow
	if opts.FieldMask.Active() {
		flow = new(flowpb.Flow)
		opts.FieldMask.Alloc(flow.ProtoReflect())
	}
	return &exporter{
		logger:  logger,
		encoder: encoder,
		writer:  writer,
		flow:    flow,
		opts:    opts,
	}, nil
}

// eventToExportEvent converts Event to ExportEvent.
func (e *exporter) eventToExportEvent(event *v1.Event) *observerpb.ExportEvent {
	switch ev := event.Event.(type) {
	case *flowpb.Flow:
		if e.opts.FieldMask.Active() {
			e.opts.FieldMask.Copy(e.flow.ProtoReflect(), ev.ProtoReflect())
			ev = e.flow
		}
		return &observerpb.ExportEvent{
			Time:     ev.GetTime(),
			NodeName: ev.GetNodeName(),
			ResponseTypes: &observerpb.ExportEvent_Flow{
				Flow: ev,
			},
		}
	case *flowpb.LostEvent:
		return &observerpb.ExportEvent{
			Time:     event.Timestamp,
			NodeName: nodeTypes.GetName(),
			ResponseTypes: &observerpb.ExportEvent_LostEvents{
				LostEvents: ev,
			},
		}
	case *flowpb.AgentEvent:
		return &observerpb.ExportEvent{
			Time:     event.Timestamp,
			NodeName: nodeTypes.GetName(),
			ResponseTypes: &observerpb.ExportEvent_AgentEvent{
				AgentEvent: ev,
			},
		}
	case *flowpb.DebugEvent:
		return &observerpb.ExportEvent{
			Time:     event.Timestamp,
			NodeName: nodeTypes.GetName(),
			ResponseTypes: &observerpb.ExportEvent_DebugEvent{
				DebugEvent: ev,
			},
		}
	default:
		return nil
	}
}

func (e *exporter) Stop() error {
	if e.writer == nil {
		// Already stoppped
		return nil
	}
	err := e.writer.Close()
	e.writer = nil
	return err
}

// OnDecodedEvent implements the observeroption.OnDecodedEvent interface.
//
// It takes care of applying filters on the received event, and if allowed, proceeds to invoke the
// registered OnExportEvent hooks. If none of the hooks return true (abort signal) the event is then
// wrapped in observerpb.ExportEvent before being json-encoded and written to its underlying writer.
func (e *exporter) OnDecodedEvent(ctx context.Context, ev *v1.Event) (bool, error) {
	select {
	case <-ctx.Done():
		return false, nil
	default:
	}
	if !filters.Apply(e.opts.AllowFilters(), e.opts.DenyFilters(), ev) {
		return false, nil
	}
	res := e.eventToExportEvent(ev)
	if res == nil {
		return false, nil
	}
	return false, e.encoder.Encode(res)
}

var _ io.WriteCloser = (*noopWriteCloser)(nil)

// nopWriteCloser returns a io.ReadWrite with a no-op Close method wrapping
// the provided io.Writer w.
type noopWriteCloser struct {
	w io.Writer
}

func (nwc *noopWriteCloser) Write(p []byte) (int, error) {
	return nwc.w.Write(p)
}

func (nwc *noopWriteCloser) Close() error {
	return nil
}
