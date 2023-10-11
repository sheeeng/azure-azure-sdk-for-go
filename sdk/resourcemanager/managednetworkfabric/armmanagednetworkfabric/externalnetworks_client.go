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

// ExternalNetworksClient contains the methods for the ExternalNetworks group.
// Don't use this type directly, use NewExternalNetworksClient() instead.
type ExternalNetworksClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewExternalNetworksClient creates a new instance of ExternalNetworksClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewExternalNetworksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ExternalNetworksClient, error) {
	cl, err := arm.NewClient(moduleName+".ExternalNetworksClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ExternalNetworksClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Creates ExternalNetwork PUT method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - l3IsolationDomainName - Name of the L3 Isolation Domain.
//   - externalNetworkName - Name of the External Network.
//   - body - Request payload.
//   - options - ExternalNetworksClientBeginCreateOptions contains the optional parameters for the ExternalNetworksClient.BeginCreate
//     method.
func (client *ExternalNetworksClient) BeginCreate(ctx context.Context, resourceGroupName string, l3IsolationDomainName string, externalNetworkName string, body ExternalNetwork, options *ExternalNetworksClientBeginCreateOptions) (*runtime.Poller[ExternalNetworksClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, l3IsolationDomainName, externalNetworkName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ExternalNetworksClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ExternalNetworksClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - Creates ExternalNetwork PUT method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *ExternalNetworksClient) create(ctx context.Context, resourceGroupName string, l3IsolationDomainName string, externalNetworkName string, body ExternalNetwork, options *ExternalNetworksClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, l3IsolationDomainName, externalNetworkName, body, options)
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
func (client *ExternalNetworksClient) createCreateRequest(ctx context.Context, resourceGroupName string, l3IsolationDomainName string, externalNetworkName string, body ExternalNetwork, options *ExternalNetworksClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains/{l3IsolationDomainName}/externalNetworks/{externalNetworkName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if l3IsolationDomainName == "" {
		return nil, errors.New("parameter l3IsolationDomainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{l3IsolationDomainName}", url.PathEscape(l3IsolationDomainName))
	if externalNetworkName == "" {
		return nil, errors.New("parameter externalNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{externalNetworkName}", url.PathEscape(externalNetworkName))
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

// BeginDelete - Implements ExternalNetworks DELETE method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - l3IsolationDomainName - Name of the L3 Isolation Domain.
//   - externalNetworkName - Name of the External Network.
//   - options - ExternalNetworksClientBeginDeleteOptions contains the optional parameters for the ExternalNetworksClient.BeginDelete
//     method.
func (client *ExternalNetworksClient) BeginDelete(ctx context.Context, resourceGroupName string, l3IsolationDomainName string, externalNetworkName string, options *ExternalNetworksClientBeginDeleteOptions) (*runtime.Poller[ExternalNetworksClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, l3IsolationDomainName, externalNetworkName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ExternalNetworksClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ExternalNetworksClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Implements ExternalNetworks DELETE method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *ExternalNetworksClient) deleteOperation(ctx context.Context, resourceGroupName string, l3IsolationDomainName string, externalNetworkName string, options *ExternalNetworksClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, l3IsolationDomainName, externalNetworkName, options)
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
func (client *ExternalNetworksClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, l3IsolationDomainName string, externalNetworkName string, options *ExternalNetworksClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains/{l3IsolationDomainName}/externalNetworks/{externalNetworkName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if l3IsolationDomainName == "" {
		return nil, errors.New("parameter l3IsolationDomainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{l3IsolationDomainName}", url.PathEscape(l3IsolationDomainName))
	if externalNetworkName == "" {
		return nil, errors.New("parameter externalNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{externalNetworkName}", url.PathEscape(externalNetworkName))
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

// Get - Implements ExternalNetworks GET method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - l3IsolationDomainName - Name of the L3 Isolation Domain.
//   - externalNetworkName - Name of the External Network.
//   - options - ExternalNetworksClientGetOptions contains the optional parameters for the ExternalNetworksClient.Get method.
func (client *ExternalNetworksClient) Get(ctx context.Context, resourceGroupName string, l3IsolationDomainName string, externalNetworkName string, options *ExternalNetworksClientGetOptions) (ExternalNetworksClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, l3IsolationDomainName, externalNetworkName, options)
	if err != nil {
		return ExternalNetworksClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExternalNetworksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExternalNetworksClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExternalNetworksClient) getCreateRequest(ctx context.Context, resourceGroupName string, l3IsolationDomainName string, externalNetworkName string, options *ExternalNetworksClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains/{l3IsolationDomainName}/externalNetworks/{externalNetworkName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if l3IsolationDomainName == "" {
		return nil, errors.New("parameter l3IsolationDomainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{l3IsolationDomainName}", url.PathEscape(l3IsolationDomainName))
	if externalNetworkName == "" {
		return nil, errors.New("parameter externalNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{externalNetworkName}", url.PathEscape(externalNetworkName))
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
func (client *ExternalNetworksClient) getHandleResponse(resp *http.Response) (ExternalNetworksClientGetResponse, error) {
	result := ExternalNetworksClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExternalNetwork); err != nil {
		return ExternalNetworksClientGetResponse{}, err
	}
	return result, nil
}

// NewListByL3IsolationDomainPager - Implements External Networks list by resource group GET method.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - l3IsolationDomainName - Name of the L3 Isolation Domain.
//   - options - ExternalNetworksClientListByL3IsolationDomainOptions contains the optional parameters for the ExternalNetworksClient.NewListByL3IsolationDomainPager
//     method.
func (client *ExternalNetworksClient) NewListByL3IsolationDomainPager(resourceGroupName string, l3IsolationDomainName string, options *ExternalNetworksClientListByL3IsolationDomainOptions) *runtime.Pager[ExternalNetworksClientListByL3IsolationDomainResponse] {
	return runtime.NewPager(runtime.PagingHandler[ExternalNetworksClientListByL3IsolationDomainResponse]{
		More: func(page ExternalNetworksClientListByL3IsolationDomainResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ExternalNetworksClientListByL3IsolationDomainResponse) (ExternalNetworksClientListByL3IsolationDomainResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByL3IsolationDomainCreateRequest(ctx, resourceGroupName, l3IsolationDomainName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ExternalNetworksClientListByL3IsolationDomainResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ExternalNetworksClientListByL3IsolationDomainResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ExternalNetworksClientListByL3IsolationDomainResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByL3IsolationDomainHandleResponse(resp)
		},
	})
}

// listByL3IsolationDomainCreateRequest creates the ListByL3IsolationDomain request.
func (client *ExternalNetworksClient) listByL3IsolationDomainCreateRequest(ctx context.Context, resourceGroupName string, l3IsolationDomainName string, options *ExternalNetworksClientListByL3IsolationDomainOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains/{l3IsolationDomainName}/externalNetworks"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if l3IsolationDomainName == "" {
		return nil, errors.New("parameter l3IsolationDomainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{l3IsolationDomainName}", url.PathEscape(l3IsolationDomainName))
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

// listByL3IsolationDomainHandleResponse handles the ListByL3IsolationDomain response.
func (client *ExternalNetworksClient) listByL3IsolationDomainHandleResponse(resp *http.Response) (ExternalNetworksClientListByL3IsolationDomainResponse, error) {
	result := ExternalNetworksClientListByL3IsolationDomainResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExternalNetworksList); err != nil {
		return ExternalNetworksClientListByL3IsolationDomainResponse{}, err
	}
	return result, nil
}

// BeginUpdate - API to update certain properties of the ExternalNetworks resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - l3IsolationDomainName - Name of the L3 Isolation Domain.
//   - externalNetworkName - Name of the External Network.
//   - body - ExternalNetwork properties to update. Only annotations are supported.
//   - options - ExternalNetworksClientBeginUpdateOptions contains the optional parameters for the ExternalNetworksClient.BeginUpdate
//     method.
func (client *ExternalNetworksClient) BeginUpdate(ctx context.Context, resourceGroupName string, l3IsolationDomainName string, externalNetworkName string, body ExternalNetworkPatch, options *ExternalNetworksClientBeginUpdateOptions) (*runtime.Poller[ExternalNetworksClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, l3IsolationDomainName, externalNetworkName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ExternalNetworksClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ExternalNetworksClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - API to update certain properties of the ExternalNetworks resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *ExternalNetworksClient) update(ctx context.Context, resourceGroupName string, l3IsolationDomainName string, externalNetworkName string, body ExternalNetworkPatch, options *ExternalNetworksClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, l3IsolationDomainName, externalNetworkName, body, options)
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
func (client *ExternalNetworksClient) updateCreateRequest(ctx context.Context, resourceGroupName string, l3IsolationDomainName string, externalNetworkName string, body ExternalNetworkPatch, options *ExternalNetworksClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains/{l3IsolationDomainName}/externalNetworks/{externalNetworkName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if l3IsolationDomainName == "" {
		return nil, errors.New("parameter l3IsolationDomainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{l3IsolationDomainName}", url.PathEscape(l3IsolationDomainName))
	if externalNetworkName == "" {
		return nil, errors.New("parameter externalNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{externalNetworkName}", url.PathEscape(externalNetworkName))
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

// BeginUpdateAdministrativeState - Executes update operation to enable or disable administrative State for externalNetwork.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - l3IsolationDomainName - Name of the L3 Isolation Domain.
//   - externalNetworkName - Name of the External Network.
//   - body - Request payload.
//   - options - ExternalNetworksClientBeginUpdateAdministrativeStateOptions contains the optional parameters for the ExternalNetworksClient.BeginUpdateAdministrativeState
//     method.
func (client *ExternalNetworksClient) BeginUpdateAdministrativeState(ctx context.Context, resourceGroupName string, l3IsolationDomainName string, externalNetworkName string, body UpdateAdministrativeState, options *ExternalNetworksClientBeginUpdateAdministrativeStateOptions) (*runtime.Poller[ExternalNetworksClientUpdateAdministrativeStateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updateAdministrativeState(ctx, resourceGroupName, l3IsolationDomainName, externalNetworkName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ExternalNetworksClientUpdateAdministrativeStateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ExternalNetworksClientUpdateAdministrativeStateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// UpdateAdministrativeState - Executes update operation to enable or disable administrative State for externalNetwork.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *ExternalNetworksClient) updateAdministrativeState(ctx context.Context, resourceGroupName string, l3IsolationDomainName string, externalNetworkName string, body UpdateAdministrativeState, options *ExternalNetworksClientBeginUpdateAdministrativeStateOptions) (*http.Response, error) {
	req, err := client.updateAdministrativeStateCreateRequest(ctx, resourceGroupName, l3IsolationDomainName, externalNetworkName, body, options)
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
func (client *ExternalNetworksClient) updateAdministrativeStateCreateRequest(ctx context.Context, resourceGroupName string, l3IsolationDomainName string, externalNetworkName string, body UpdateAdministrativeState, options *ExternalNetworksClientBeginUpdateAdministrativeStateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains/{l3IsolationDomainName}/externalNetworks/{externalNetworkName}/updateAdministrativeState"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if l3IsolationDomainName == "" {
		return nil, errors.New("parameter l3IsolationDomainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{l3IsolationDomainName}", url.PathEscape(l3IsolationDomainName))
	if externalNetworkName == "" {
		return nil, errors.New("parameter externalNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{externalNetworkName}", url.PathEscape(externalNetworkName))
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

// BeginUpdateStaticRouteBfdAdministrativeState - Update Static Route BFD for external Network.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - l3IsolationDomainName - Name of the L3 Isolation Domain.
//   - externalNetworkName - Name of the External Network.
//   - body - Request payload.
//   - options - ExternalNetworksClientBeginUpdateStaticRouteBfdAdministrativeStateOptions contains the optional parameters for
//     the ExternalNetworksClient.BeginUpdateStaticRouteBfdAdministrativeState method.
func (client *ExternalNetworksClient) BeginUpdateStaticRouteBfdAdministrativeState(ctx context.Context, resourceGroupName string, l3IsolationDomainName string, externalNetworkName string, body UpdateAdministrativeState, options *ExternalNetworksClientBeginUpdateStaticRouteBfdAdministrativeStateOptions) (*runtime.Poller[ExternalNetworksClientUpdateStaticRouteBfdAdministrativeStateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updateStaticRouteBfdAdministrativeState(ctx, resourceGroupName, l3IsolationDomainName, externalNetworkName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ExternalNetworksClientUpdateStaticRouteBfdAdministrativeStateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ExternalNetworksClientUpdateStaticRouteBfdAdministrativeStateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// UpdateStaticRouteBfdAdministrativeState - Update Static Route BFD for external Network.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *ExternalNetworksClient) updateStaticRouteBfdAdministrativeState(ctx context.Context, resourceGroupName string, l3IsolationDomainName string, externalNetworkName string, body UpdateAdministrativeState, options *ExternalNetworksClientBeginUpdateStaticRouteBfdAdministrativeStateOptions) (*http.Response, error) {
	req, err := client.updateStaticRouteBfdAdministrativeStateCreateRequest(ctx, resourceGroupName, l3IsolationDomainName, externalNetworkName, body, options)
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

// updateStaticRouteBfdAdministrativeStateCreateRequest creates the UpdateStaticRouteBfdAdministrativeState request.
func (client *ExternalNetworksClient) updateStaticRouteBfdAdministrativeStateCreateRequest(ctx context.Context, resourceGroupName string, l3IsolationDomainName string, externalNetworkName string, body UpdateAdministrativeState, options *ExternalNetworksClientBeginUpdateStaticRouteBfdAdministrativeStateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains/{l3IsolationDomainName}/externalNetworks/{externalNetworkName}/updateStaticRouteBfdAdministrativeState"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if l3IsolationDomainName == "" {
		return nil, errors.New("parameter l3IsolationDomainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{l3IsolationDomainName}", url.PathEscape(l3IsolationDomainName))
	if externalNetworkName == "" {
		return nil, errors.New("parameter externalNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{externalNetworkName}", url.PathEscape(externalNetworkName))
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