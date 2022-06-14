// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
//
// Package network implements the Azure ARM Network service API version .
//
// Network Client
package network

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

const (
	// DefaultBaseURI is the default URI used for the service Network
	DefaultBaseURI = "https://management.azure.com"
)

// BaseClient is the base client for Network.
type BaseClient struct {
	autorest.Client
	BaseURI        string
	SubscriptionID string
}

// New creates an instance of the BaseClient client.
func New(subscriptionID string) BaseClient {
	return NewWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWithBaseURI creates an instance of the BaseClient client using a custom endpoint.  Use this when interacting with
// an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return BaseClient{
		Client:         autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI:        baseURI,
		SubscriptionID: subscriptionID,
	}
}

// CheckDNSNameAvailability checks whether a domain name in the cloudapp.azure.com zone is available for use.
// Parameters:
// location - the location of the domain name.
// domainNameLabel - the domain name to be verified. It must conform to the following regular expression:
// ^[a-z][a-z0-9-]{1,61}[a-z0-9]$.
func (client BaseClient) CheckDNSNameAvailability(ctx context.Context, location string, domainNameLabel string) (result DNSNameAvailabilityResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.CheckDNSNameAvailability")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CheckDNSNameAvailabilityPreparer(ctx, location, domainNameLabel)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.BaseClient", "CheckDNSNameAvailability", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckDNSNameAvailabilitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.BaseClient", "CheckDNSNameAvailability", resp, "Failure sending request")
		return
	}

	result, err = client.CheckDNSNameAvailabilityResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.BaseClient", "CheckDNSNameAvailability", resp, "Failure responding to request")
		return
	}

	return
}

// CheckDNSNameAvailabilityPreparer prepares the CheckDNSNameAvailability request.
func (client BaseClient) CheckDNSNameAvailabilityPreparer(ctx context.Context, location string, domainNameLabel string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version":     APIVersion,
		"domainNameLabel": autorest.Encode("query", domainNameLabel),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/CheckDnsNameAvailability", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CheckDNSNameAvailabilitySender sends the CheckDNSNameAvailability request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) CheckDNSNameAvailabilitySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CheckDNSNameAvailabilityResponder handles the response to the CheckDNSNameAvailability request. The method always
// closes the http.Response Body.
func (client BaseClient) CheckDNSNameAvailabilityResponder(resp *http.Response) (result DNSNameAvailabilityResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Generatevirtualwanvpnserverconfigurationvpnprofile generates a unique VPN profile for P2S clients for VirtualWan and
// associated VpnServerConfiguration combination in the specified resource group.
// Parameters:
// resourceGroupName - the resource group name.
// virtualWANName - the name of the VirtualWAN whose associated VpnServerConfigurations is needed.
// vpnClientParams - parameters supplied to the generate VirtualWan VPN profile generation operation.
func (client BaseClient) Generatevirtualwanvpnserverconfigurationvpnprofile(ctx context.Context, resourceGroupName string, virtualWANName string, vpnClientParams VirtualWanVpnProfileParameters) (result GeneratevirtualwanvpnserverconfigurationvpnprofileFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.Generatevirtualwanvpnserverconfigurationvpnprofile")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GeneratevirtualwanvpnserverconfigurationvpnprofilePreparer(ctx, resourceGroupName, virtualWANName, vpnClientParams)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.BaseClient", "Generatevirtualwanvpnserverconfigurationvpnprofile", nil, "Failure preparing request")
		return
	}

	result, err = client.GeneratevirtualwanvpnserverconfigurationvpnprofileSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.BaseClient", "Generatevirtualwanvpnserverconfigurationvpnprofile", result.Response(), "Failure sending request")
		return
	}

	return
}

// GeneratevirtualwanvpnserverconfigurationvpnprofilePreparer prepares the Generatevirtualwanvpnserverconfigurationvpnprofile request.
func (client BaseClient) GeneratevirtualwanvpnserverconfigurationvpnprofilePreparer(ctx context.Context, resourceGroupName string, virtualWANName string, vpnClientParams VirtualWanVpnProfileParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"virtualWANName":    autorest.Encode("path", virtualWANName),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{virtualWANName}/GenerateVpnProfile", pathParameters),
		autorest.WithJSON(vpnClientParams),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GeneratevirtualwanvpnserverconfigurationvpnprofileSender sends the Generatevirtualwanvpnserverconfigurationvpnprofile request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) GeneratevirtualwanvpnserverconfigurationvpnprofileSender(req *http.Request) (future GeneratevirtualwanvpnserverconfigurationvpnprofileFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// GeneratevirtualwanvpnserverconfigurationvpnprofileResponder handles the response to the Generatevirtualwanvpnserverconfigurationvpnprofile request. The method always
// closes the http.Response Body.
func (client BaseClient) GeneratevirtualwanvpnserverconfigurationvpnprofileResponder(resp *http.Response) (result VpnProfileResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// SupportedSecurityProviders gives the supported security providers for the virtual wan.
// Parameters:
// resourceGroupName - the resource group name.
// virtualWANName - the name of the VirtualWAN for which supported security providers are needed.
func (client BaseClient) SupportedSecurityProviders(ctx context.Context, resourceGroupName string, virtualWANName string) (result VirtualWanSecurityProviders, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.SupportedSecurityProviders")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.SupportedSecurityProvidersPreparer(ctx, resourceGroupName, virtualWANName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.BaseClient", "SupportedSecurityProviders", nil, "Failure preparing request")
		return
	}

	resp, err := client.SupportedSecurityProvidersSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.BaseClient", "SupportedSecurityProviders", resp, "Failure sending request")
		return
	}

	result, err = client.SupportedSecurityProvidersResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.BaseClient", "SupportedSecurityProviders", resp, "Failure responding to request")
		return
	}

	return
}

// SupportedSecurityProvidersPreparer prepares the SupportedSecurityProviders request.
func (client BaseClient) SupportedSecurityProvidersPreparer(ctx context.Context, resourceGroupName string, virtualWANName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"virtualWANName":    autorest.Encode("path", virtualWANName),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{virtualWANName}/supportedSecurityProviders", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// SupportedSecurityProvidersSender sends the SupportedSecurityProviders request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) SupportedSecurityProvidersSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// SupportedSecurityProvidersResponder handles the response to the SupportedSecurityProviders request. The method always
// closes the http.Response Body.
func (client BaseClient) SupportedSecurityProvidersResponder(resp *http.Response) (result VirtualWanSecurityProviders, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
