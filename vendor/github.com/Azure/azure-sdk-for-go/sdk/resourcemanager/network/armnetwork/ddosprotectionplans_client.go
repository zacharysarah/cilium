//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// DdosProtectionPlansClient contains the methods for the DdosProtectionPlans group.
// Don't use this type directly, use NewDdosProtectionPlansClient() instead.
type DdosProtectionPlansClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDdosProtectionPlansClient creates a new instance of DdosProtectionPlansClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDdosProtectionPlansClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DdosProtectionPlansClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &DdosProtectionPlansClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a DDoS protection plan.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - The name of the resource group.
// ddosProtectionPlanName - The name of the DDoS protection plan.
// parameters - Parameters supplied to the create or update operation.
// options - DdosProtectionPlansClientBeginCreateOrUpdateOptions contains the optional parameters for the DdosProtectionPlansClient.BeginCreateOrUpdate
// method.
func (client *DdosProtectionPlansClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, parameters DdosProtectionPlan, options *DdosProtectionPlansClientBeginCreateOrUpdateOptions) (*runtime.Poller[DdosProtectionPlansClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, ddosProtectionPlanName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[DdosProtectionPlansClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[DdosProtectionPlansClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a DDoS protection plan.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
func (client *DdosProtectionPlansClient) createOrUpdate(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, parameters DdosProtectionPlan, options *DdosProtectionPlansClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, ddosProtectionPlanName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DdosProtectionPlansClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, parameters DdosProtectionPlan, options *DdosProtectionPlansClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosProtectionPlans/{ddosProtectionPlanName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ddosProtectionPlanName == "" {
		return nil, errors.New("parameter ddosProtectionPlanName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ddosProtectionPlanName}", url.PathEscape(ddosProtectionPlanName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the specified DDoS protection plan.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - The name of the resource group.
// ddosProtectionPlanName - The name of the DDoS protection plan.
// options - DdosProtectionPlansClientBeginDeleteOptions contains the optional parameters for the DdosProtectionPlansClient.BeginDelete
// method.
func (client *DdosProtectionPlansClient) BeginDelete(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, options *DdosProtectionPlansClientBeginDeleteOptions) (*runtime.Poller[DdosProtectionPlansClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, ddosProtectionPlanName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[DdosProtectionPlansClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[DdosProtectionPlansClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the specified DDoS protection plan.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
func (client *DdosProtectionPlansClient) deleteOperation(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, options *DdosProtectionPlansClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, ddosProtectionPlanName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DdosProtectionPlansClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, options *DdosProtectionPlansClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosProtectionPlans/{ddosProtectionPlanName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ddosProtectionPlanName == "" {
		return nil, errors.New("parameter ddosProtectionPlanName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ddosProtectionPlanName}", url.PathEscape(ddosProtectionPlanName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets information about the specified DDoS protection plan.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - The name of the resource group.
// ddosProtectionPlanName - The name of the DDoS protection plan.
// options - DdosProtectionPlansClientGetOptions contains the optional parameters for the DdosProtectionPlansClient.Get method.
func (client *DdosProtectionPlansClient) Get(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, options *DdosProtectionPlansClientGetOptions) (DdosProtectionPlansClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, ddosProtectionPlanName, options)
	if err != nil {
		return DdosProtectionPlansClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DdosProtectionPlansClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DdosProtectionPlansClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DdosProtectionPlansClient) getCreateRequest(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, options *DdosProtectionPlansClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosProtectionPlans/{ddosProtectionPlanName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ddosProtectionPlanName == "" {
		return nil, errors.New("parameter ddosProtectionPlanName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ddosProtectionPlanName}", url.PathEscape(ddosProtectionPlanName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DdosProtectionPlansClient) getHandleResponse(resp *http.Response) (DdosProtectionPlansClientGetResponse, error) {
	result := DdosProtectionPlansClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DdosProtectionPlan); err != nil {
		return DdosProtectionPlansClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all DDoS protection plans in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// options - DdosProtectionPlansClientListOptions contains the optional parameters for the DdosProtectionPlansClient.List
// method.
func (client *DdosProtectionPlansClient) NewListPager(options *DdosProtectionPlansClientListOptions) *runtime.Pager[DdosProtectionPlansClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[DdosProtectionPlansClientListResponse]{
		More: func(page DdosProtectionPlansClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DdosProtectionPlansClientListResponse) (DdosProtectionPlansClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DdosProtectionPlansClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DdosProtectionPlansClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DdosProtectionPlansClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *DdosProtectionPlansClient) listCreateRequest(ctx context.Context, options *DdosProtectionPlansClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/ddosProtectionPlans"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DdosProtectionPlansClient) listHandleResponse(resp *http.Response) (DdosProtectionPlansClientListResponse, error) {
	result := DdosProtectionPlansClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DdosProtectionPlanListResult); err != nil {
		return DdosProtectionPlansClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets all the DDoS protection plans in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - The name of the resource group.
// options - DdosProtectionPlansClientListByResourceGroupOptions contains the optional parameters for the DdosProtectionPlansClient.ListByResourceGroup
// method.
func (client *DdosProtectionPlansClient) NewListByResourceGroupPager(resourceGroupName string, options *DdosProtectionPlansClientListByResourceGroupOptions) *runtime.Pager[DdosProtectionPlansClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[DdosProtectionPlansClientListByResourceGroupResponse]{
		More: func(page DdosProtectionPlansClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DdosProtectionPlansClientListByResourceGroupResponse) (DdosProtectionPlansClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DdosProtectionPlansClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DdosProtectionPlansClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DdosProtectionPlansClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *DdosProtectionPlansClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *DdosProtectionPlansClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosProtectionPlans"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *DdosProtectionPlansClient) listByResourceGroupHandleResponse(resp *http.Response) (DdosProtectionPlansClientListByResourceGroupResponse, error) {
	result := DdosProtectionPlansClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DdosProtectionPlanListResult); err != nil {
		return DdosProtectionPlansClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// UpdateTags - Update a DDoS protection plan tags.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - The name of the resource group.
// ddosProtectionPlanName - The name of the DDoS protection plan.
// parameters - Parameters supplied to the update DDoS protection plan resource tags.
// options - DdosProtectionPlansClientUpdateTagsOptions contains the optional parameters for the DdosProtectionPlansClient.UpdateTags
// method.
func (client *DdosProtectionPlansClient) UpdateTags(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, parameters TagsObject, options *DdosProtectionPlansClientUpdateTagsOptions) (DdosProtectionPlansClientUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, ddosProtectionPlanName, parameters, options)
	if err != nil {
		return DdosProtectionPlansClientUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DdosProtectionPlansClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DdosProtectionPlansClientUpdateTagsResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *DdosProtectionPlansClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, parameters TagsObject, options *DdosProtectionPlansClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosProtectionPlans/{ddosProtectionPlanName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ddosProtectionPlanName == "" {
		return nil, errors.New("parameter ddosProtectionPlanName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ddosProtectionPlanName}", url.PathEscape(ddosProtectionPlanName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *DdosProtectionPlansClient) updateTagsHandleResponse(resp *http.Response) (DdosProtectionPlansClientUpdateTagsResponse, error) {
	result := DdosProtectionPlansClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DdosProtectionPlan); err != nil {
		return DdosProtectionPlansClientUpdateTagsResponse{}, err
	}
	return result, nil
}
