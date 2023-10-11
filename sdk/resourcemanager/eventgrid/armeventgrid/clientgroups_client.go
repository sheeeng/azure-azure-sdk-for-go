//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armeventgrid

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ClientGroupsClient contains the methods for the ClientGroups group.
// Don't use this type directly, use NewClientGroupsClient() instead.
type ClientGroupsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewClientGroupsClient creates a new instance of ClientGroupsClient with the specified values.
//   - subscriptionID - Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientGroupsClient, error) {
	cl, err := arm.NewClient(moduleName+".ClientGroupsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ClientGroupsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update a client group with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - namespaceName - Name of the namespace.
//   - clientGroupName - The client group name.
//   - clientGroupInfo - Client group information.
//   - options - ClientGroupsClientBeginCreateOrUpdateOptions contains the optional parameters for the ClientGroupsClient.BeginCreateOrUpdate
//     method.
func (client *ClientGroupsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, clientGroupName string, clientGroupInfo ClientGroup, options *ClientGroupsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ClientGroupsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, namespaceName, clientGroupName, clientGroupInfo, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ClientGroupsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ClientGroupsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create or update a client group with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
func (client *ClientGroupsClient) createOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, clientGroupName string, clientGroupInfo ClientGroup, options *ClientGroupsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, namespaceName, clientGroupName, clientGroupInfo, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ClientGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, clientGroupName string, clientGroupInfo ClientGroup, options *ClientGroupsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/namespaces/{namespaceName}/clientGroups/{clientGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if clientGroupName == "" {
		return nil, errors.New("parameter clientGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clientGroupName}", url.PathEscape(clientGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, clientGroupInfo)
}

// BeginDelete - Delete an existing client group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - namespaceName - Name of the namespace.
//   - clientGroupName - Name of the client group.
//   - options - ClientGroupsClientBeginDeleteOptions contains the optional parameters for the ClientGroupsClient.BeginDelete
//     method.
func (client *ClientGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, namespaceName string, clientGroupName string, options *ClientGroupsClientBeginDeleteOptions) (*runtime.Poller[ClientGroupsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, namespaceName, clientGroupName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ClientGroupsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ClientGroupsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete an existing client group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
func (client *ClientGroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, namespaceName string, clientGroupName string, options *ClientGroupsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, namespaceName, clientGroupName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ClientGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, clientGroupName string, options *ClientGroupsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/namespaces/{namespaceName}/clientGroups/{clientGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if clientGroupName == "" {
		return nil, errors.New("parameter clientGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clientGroupName}", url.PathEscape(clientGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get properties of a client group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - namespaceName - Name of the namespace.
//   - clientGroupName - Name of the client group.
//   - options - ClientGroupsClientGetOptions contains the optional parameters for the ClientGroupsClient.Get method.
func (client *ClientGroupsClient) Get(ctx context.Context, resourceGroupName string, namespaceName string, clientGroupName string, options *ClientGroupsClientGetOptions) (ClientGroupsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, namespaceName, clientGroupName, options)
	if err != nil {
		return ClientGroupsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClientGroupsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ClientGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, clientGroupName string, options *ClientGroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/namespaces/{namespaceName}/clientGroups/{clientGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if clientGroupName == "" {
		return nil, errors.New("parameter clientGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clientGroupName}", url.PathEscape(clientGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ClientGroupsClient) getHandleResponse(resp *http.Response) (ClientGroupsClientGetResponse, error) {
	result := ClientGroupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClientGroup); err != nil {
		return ClientGroupsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByNamespacePager - Get all the client groups under a namespace.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - namespaceName - Name of the namespace.
//   - options - ClientGroupsClientListByNamespaceOptions contains the optional parameters for the ClientGroupsClient.NewListByNamespacePager
//     method.
func (client *ClientGroupsClient) NewListByNamespacePager(resourceGroupName string, namespaceName string, options *ClientGroupsClientListByNamespaceOptions) *runtime.Pager[ClientGroupsClientListByNamespaceResponse] {
	return runtime.NewPager(runtime.PagingHandler[ClientGroupsClientListByNamespaceResponse]{
		More: func(page ClientGroupsClientListByNamespaceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ClientGroupsClientListByNamespaceResponse) (ClientGroupsClientListByNamespaceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByNamespaceCreateRequest(ctx, resourceGroupName, namespaceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ClientGroupsClientListByNamespaceResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ClientGroupsClientListByNamespaceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ClientGroupsClientListByNamespaceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByNamespaceHandleResponse(resp)
		},
	})
}

// listByNamespaceCreateRequest creates the ListByNamespace request.
func (client *ClientGroupsClient) listByNamespaceCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, options *ClientGroupsClientListByNamespaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/namespaces/{namespaceName}/clientGroups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByNamespaceHandleResponse handles the ListByNamespace response.
func (client *ClientGroupsClient) listByNamespaceHandleResponse(resp *http.Response) (ClientGroupsClientListByNamespaceResponse, error) {
	result := ClientGroupsClientListByNamespaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClientGroupsListResult); err != nil {
		return ClientGroupsClientListByNamespaceResponse{}, err
	}
	return result, nil
}