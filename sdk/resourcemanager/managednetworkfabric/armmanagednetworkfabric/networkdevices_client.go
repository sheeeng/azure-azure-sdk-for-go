//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmanagednetworkfabric

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// NetworkDevicesClient contains the methods for the NetworkDevices group.
// Don't use this type directly, use NewNetworkDevicesClient() instead.
type NetworkDevicesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewNetworkDevicesClient creates a new instance of NetworkDevicesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewNetworkDevicesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*NetworkDevicesClient, error) {
	cl, err := arm.NewClient(moduleName+".NetworkDevicesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &NetworkDevicesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Create a Network Device resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - networkDeviceName - Name of the Network Device.
//   - body - Request payload.
//   - options - NetworkDevicesClientBeginCreateOptions contains the optional parameters for the NetworkDevicesClient.BeginCreate
//     method.
func (client *NetworkDevicesClient) BeginCreate(ctx context.Context, resourceGroupName string, networkDeviceName string, body NetworkDevice, options *NetworkDevicesClientBeginCreateOptions) (*runtime.Poller[NetworkDevicesClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, networkDeviceName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkDevicesClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[NetworkDevicesClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - Create a Network Device resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *NetworkDevicesClient) create(ctx context.Context, resourceGroupName string, networkDeviceName string, body NetworkDevice, options *NetworkDevicesClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, networkDeviceName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *NetworkDevicesClient) createCreateRequest(ctx context.Context, resourceGroupName string, networkDeviceName string, body NetworkDevice, options *NetworkDevicesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/networkDevices/{networkDeviceName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkDeviceName == "" {
		return nil, errors.New("parameter networkDeviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkDeviceName}", url.PathEscape(networkDeviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginDelete - Delete the Network Device resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - networkDeviceName - Name of the Network Device.
//   - options - NetworkDevicesClientBeginDeleteOptions contains the optional parameters for the NetworkDevicesClient.BeginDelete
//     method.
func (client *NetworkDevicesClient) BeginDelete(ctx context.Context, resourceGroupName string, networkDeviceName string, options *NetworkDevicesClientBeginDeleteOptions) (*runtime.Poller[NetworkDevicesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, networkDeviceName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkDevicesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[NetworkDevicesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete the Network Device resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *NetworkDevicesClient) deleteOperation(ctx context.Context, resourceGroupName string, networkDeviceName string, options *NetworkDevicesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, networkDeviceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *NetworkDevicesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, networkDeviceName string, options *NetworkDevicesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/networkDevices/{networkDeviceName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkDeviceName == "" {
		return nil, errors.New("parameter networkDeviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkDeviceName}", url.PathEscape(networkDeviceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the Network Device resource details.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - networkDeviceName - Name of the Network Device.
//   - options - NetworkDevicesClientGetOptions contains the optional parameters for the NetworkDevicesClient.Get method.
func (client *NetworkDevicesClient) Get(ctx context.Context, resourceGroupName string, networkDeviceName string, options *NetworkDevicesClientGetOptions) (NetworkDevicesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkDeviceName, options)
	if err != nil {
		return NetworkDevicesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NetworkDevicesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return NetworkDevicesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *NetworkDevicesClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkDeviceName string, options *NetworkDevicesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/networkDevices/{networkDeviceName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkDeviceName == "" {
		return nil, errors.New("parameter networkDeviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkDeviceName}", url.PathEscape(networkDeviceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *NetworkDevicesClient) getHandleResponse(resp *http.Response) (NetworkDevicesClientGetResponse, error) {
	result := NetworkDevicesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkDevice); err != nil {
		return NetworkDevicesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List all the Network Device resources in a given resource group.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - NetworkDevicesClientListByResourceGroupOptions contains the optional parameters for the NetworkDevicesClient.NewListByResourceGroupPager
//     method.
func (client *NetworkDevicesClient) NewListByResourceGroupPager(resourceGroupName string, options *NetworkDevicesClientListByResourceGroupOptions) *runtime.Pager[NetworkDevicesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[NetworkDevicesClientListByResourceGroupResponse]{
		More: func(page NetworkDevicesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NetworkDevicesClientListByResourceGroupResponse) (NetworkDevicesClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return NetworkDevicesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return NetworkDevicesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return NetworkDevicesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *NetworkDevicesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *NetworkDevicesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/networkDevices"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *NetworkDevicesClient) listByResourceGroupHandleResponse(resp *http.Response) (NetworkDevicesClientListByResourceGroupResponse, error) {
	result := NetworkDevicesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkDevicesListResult); err != nil {
		return NetworkDevicesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List all the Network Device resources in a given subscription.
//
// Generated from API version 2023-06-15
//   - options - NetworkDevicesClientListBySubscriptionOptions contains the optional parameters for the NetworkDevicesClient.NewListBySubscriptionPager
//     method.
func (client *NetworkDevicesClient) NewListBySubscriptionPager(options *NetworkDevicesClientListBySubscriptionOptions) *runtime.Pager[NetworkDevicesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[NetworkDevicesClientListBySubscriptionResponse]{
		More: func(page NetworkDevicesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NetworkDevicesClientListBySubscriptionResponse) (NetworkDevicesClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return NetworkDevicesClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return NetworkDevicesClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return NetworkDevicesClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *NetworkDevicesClient) listBySubscriptionCreateRequest(ctx context.Context, options *NetworkDevicesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ManagedNetworkFabric/networkDevices"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *NetworkDevicesClient) listBySubscriptionHandleResponse(resp *http.Response) (NetworkDevicesClientListBySubscriptionResponse, error) {
	result := NetworkDevicesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkDevicesListResult); err != nil {
		return NetworkDevicesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginReboot - Reboot the Network Device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - networkDeviceName - Name of the Network Device.
//   - body - Request payload.
//   - options - NetworkDevicesClientBeginRebootOptions contains the optional parameters for the NetworkDevicesClient.BeginReboot
//     method.
func (client *NetworkDevicesClient) BeginReboot(ctx context.Context, resourceGroupName string, networkDeviceName string, body RebootProperties, options *NetworkDevicesClientBeginRebootOptions) (*runtime.Poller[NetworkDevicesClientRebootResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.reboot(ctx, resourceGroupName, networkDeviceName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkDevicesClientRebootResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[NetworkDevicesClientRebootResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Reboot - Reboot the Network Device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *NetworkDevicesClient) reboot(ctx context.Context, resourceGroupName string, networkDeviceName string, body RebootProperties, options *NetworkDevicesClientBeginRebootOptions) (*http.Response, error) {
	req, err := client.rebootCreateRequest(ctx, resourceGroupName, networkDeviceName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// rebootCreateRequest creates the Reboot request.
func (client *NetworkDevicesClient) rebootCreateRequest(ctx context.Context, resourceGroupName string, networkDeviceName string, body RebootProperties, options *NetworkDevicesClientBeginRebootOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/networkDevices/{networkDeviceName}/reboot"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkDeviceName == "" {
		return nil, errors.New("parameter networkDeviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkDeviceName}", url.PathEscape(networkDeviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginRefreshConfiguration - Refreshes the configuration the Network Device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - networkDeviceName - Name of the Network Device.
//   - options - NetworkDevicesClientBeginRefreshConfigurationOptions contains the optional parameters for the NetworkDevicesClient.BeginRefreshConfiguration
//     method.
func (client *NetworkDevicesClient) BeginRefreshConfiguration(ctx context.Context, resourceGroupName string, networkDeviceName string, options *NetworkDevicesClientBeginRefreshConfigurationOptions) (*runtime.Poller[NetworkDevicesClientRefreshConfigurationResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.refreshConfiguration(ctx, resourceGroupName, networkDeviceName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkDevicesClientRefreshConfigurationResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[NetworkDevicesClientRefreshConfigurationResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// RefreshConfiguration - Refreshes the configuration the Network Device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *NetworkDevicesClient) refreshConfiguration(ctx context.Context, resourceGroupName string, networkDeviceName string, options *NetworkDevicesClientBeginRefreshConfigurationOptions) (*http.Response, error) {
	req, err := client.refreshConfigurationCreateRequest(ctx, resourceGroupName, networkDeviceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// refreshConfigurationCreateRequest creates the RefreshConfiguration request.
func (client *NetworkDevicesClient) refreshConfigurationCreateRequest(ctx context.Context, resourceGroupName string, networkDeviceName string, options *NetworkDevicesClientBeginRefreshConfigurationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/networkDevices/{networkDeviceName}/refreshConfiguration"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkDeviceName == "" {
		return nil, errors.New("parameter networkDeviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkDeviceName}", url.PathEscape(networkDeviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginUpdate - Update certain properties of the Network Device resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - networkDeviceName - Name of the Network Device.
//   - body - Network Device properties to update.
//   - options - NetworkDevicesClientBeginUpdateOptions contains the optional parameters for the NetworkDevicesClient.BeginUpdate
//     method.
func (client *NetworkDevicesClient) BeginUpdate(ctx context.Context, resourceGroupName string, networkDeviceName string, body NetworkDevicePatchParameters, options *NetworkDevicesClientBeginUpdateOptions) (*runtime.Poller[NetworkDevicesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, networkDeviceName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkDevicesClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[NetworkDevicesClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Update certain properties of the Network Device resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *NetworkDevicesClient) update(ctx context.Context, resourceGroupName string, networkDeviceName string, body NetworkDevicePatchParameters, options *NetworkDevicesClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, networkDeviceName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *NetworkDevicesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, networkDeviceName string, body NetworkDevicePatchParameters, options *NetworkDevicesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/networkDevices/{networkDeviceName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkDeviceName == "" {
		return nil, errors.New("parameter networkDeviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkDeviceName}", url.PathEscape(networkDeviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginUpdateAdministrativeState - Updates the Administrative state of the Network Device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - networkDeviceName - Name of the Network Device.
//   - body - Request payload.
//   - options - NetworkDevicesClientBeginUpdateAdministrativeStateOptions contains the optional parameters for the NetworkDevicesClient.BeginUpdateAdministrativeState
//     method.
func (client *NetworkDevicesClient) BeginUpdateAdministrativeState(ctx context.Context, resourceGroupName string, networkDeviceName string, body UpdateDeviceAdministrativeState, options *NetworkDevicesClientBeginUpdateAdministrativeStateOptions) (*runtime.Poller[NetworkDevicesClientUpdateAdministrativeStateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updateAdministrativeState(ctx, resourceGroupName, networkDeviceName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkDevicesClientUpdateAdministrativeStateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[NetworkDevicesClientUpdateAdministrativeStateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// UpdateAdministrativeState - Updates the Administrative state of the Network Device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *NetworkDevicesClient) updateAdministrativeState(ctx context.Context, resourceGroupName string, networkDeviceName string, body UpdateDeviceAdministrativeState, options *NetworkDevicesClientBeginUpdateAdministrativeStateOptions) (*http.Response, error) {
	req, err := client.updateAdministrativeStateCreateRequest(ctx, resourceGroupName, networkDeviceName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateAdministrativeStateCreateRequest creates the UpdateAdministrativeState request.
func (client *NetworkDevicesClient) updateAdministrativeStateCreateRequest(ctx context.Context, resourceGroupName string, networkDeviceName string, body UpdateDeviceAdministrativeState, options *NetworkDevicesClientBeginUpdateAdministrativeStateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/networkDevices/{networkDeviceName}/updateAdministrativeState"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkDeviceName == "" {
		return nil, errors.New("parameter networkDeviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkDeviceName}", url.PathEscape(networkDeviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginUpgrade - Upgrades the version of the Network Device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - networkDeviceName - Name of the Network Device.
//   - body - Request payload.
//   - options - NetworkDevicesClientBeginUpgradeOptions contains the optional parameters for the NetworkDevicesClient.BeginUpgrade
//     method.
func (client *NetworkDevicesClient) BeginUpgrade(ctx context.Context, resourceGroupName string, networkDeviceName string, body UpdateVersion, options *NetworkDevicesClientBeginUpgradeOptions) (*runtime.Poller[NetworkDevicesClientUpgradeResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.upgrade(ctx, resourceGroupName, networkDeviceName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkDevicesClientUpgradeResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[NetworkDevicesClientUpgradeResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Upgrade - Upgrades the version of the Network Device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *NetworkDevicesClient) upgrade(ctx context.Context, resourceGroupName string, networkDeviceName string, body UpdateVersion, options *NetworkDevicesClientBeginUpgradeOptions) (*http.Response, error) {
	req, err := client.upgradeCreateRequest(ctx, resourceGroupName, networkDeviceName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// upgradeCreateRequest creates the Upgrade request.
func (client *NetworkDevicesClient) upgradeCreateRequest(ctx context.Context, resourceGroupName string, networkDeviceName string, body UpdateVersion, options *NetworkDevicesClientBeginUpgradeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/networkDevices/{networkDeviceName}/upgrade"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkDeviceName == "" {
		return nil, errors.New("parameter networkDeviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkDeviceName}", url.PathEscape(networkDeviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}