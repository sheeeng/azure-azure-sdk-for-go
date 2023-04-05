//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmediaservices

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	credential     azcore.TokenCredential
	options        *arm.ClientOptions
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - The unique identifier for a Microsoft Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	_, err := arm.NewClient(moduleName+".ClientFactory", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID, credential: credential,
		options: options.Clone(),
	}, nil
}

func (c *ClientFactory) NewAccountFiltersClient() *AccountFiltersClient {
	subClient, _ := NewAccountFiltersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewClient() *Client {
	subClient, _ := NewClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPrivateLinkResourcesClient() *PrivateLinkResourcesClient {
	subClient, _ := NewPrivateLinkResourcesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPrivateEndpointConnectionsClient() *PrivateEndpointConnectionsClient {
	subClient, _ := NewPrivateEndpointConnectionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewLocationsClient() *LocationsClient {
	subClient, _ := NewLocationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOperationStatusesClient() *OperationStatusesClient {
	subClient, _ := NewOperationStatusesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOperationResultsClient() *OperationResultsClient {
	subClient, _ := NewOperationResultsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAssetsClient() *AssetsClient {
	subClient, _ := NewAssetsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAssetFiltersClient() *AssetFiltersClient {
	subClient, _ := NewAssetFiltersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewTracksClient() *TracksClient {
	subClient, _ := NewTracksClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAssetTrackOperationStatusesClient() *AssetTrackOperationStatusesClient {
	subClient, _ := NewAssetTrackOperationStatusesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAssetTrackOperationResultsClient() *AssetTrackOperationResultsClient {
	subClient, _ := NewAssetTrackOperationResultsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewContentKeyPoliciesClient() *ContentKeyPoliciesClient {
	subClient, _ := NewContentKeyPoliciesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewTransformsClient() *TransformsClient {
	subClient, _ := NewTransformsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewJobsClient() *JobsClient {
	subClient, _ := NewJobsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewStreamingPoliciesClient() *StreamingPoliciesClient {
	subClient, _ := NewStreamingPoliciesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewStreamingLocatorsClient() *StreamingLocatorsClient {
	subClient, _ := NewStreamingLocatorsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewLiveEventsClient() *LiveEventsClient {
	subClient, _ := NewLiveEventsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewLiveOutputsClient() *LiveOutputsClient {
	subClient, _ := NewLiveOutputsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewStreamingEndpointsClient() *StreamingEndpointsClient {
	subClient, _ := NewStreamingEndpointsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}