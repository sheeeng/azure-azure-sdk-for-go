//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armvoiceservices

// CommunicationsGatewaysClientCreateOrUpdateResponse contains the response from method CommunicationsGatewaysClient.BeginCreateOrUpdate.
type CommunicationsGatewaysClientCreateOrUpdateResponse struct {
	CommunicationsGateway
}

// CommunicationsGatewaysClientDeleteResponse contains the response from method CommunicationsGatewaysClient.BeginDelete.
type CommunicationsGatewaysClientDeleteResponse struct {
	// placeholder for future response values
}

// CommunicationsGatewaysClientGetResponse contains the response from method CommunicationsGatewaysClient.Get.
type CommunicationsGatewaysClientGetResponse struct {
	CommunicationsGateway
}

// CommunicationsGatewaysClientListByResourceGroupResponse contains the response from method CommunicationsGatewaysClient.NewListByResourceGroupPager.
type CommunicationsGatewaysClientListByResourceGroupResponse struct {
	CommunicationsGatewayListResult
}

// CommunicationsGatewaysClientListBySubscriptionResponse contains the response from method CommunicationsGatewaysClient.NewListBySubscriptionPager.
type CommunicationsGatewaysClientListBySubscriptionResponse struct {
	CommunicationsGatewayListResult
}

// CommunicationsGatewaysClientUpdateResponse contains the response from method CommunicationsGatewaysClient.Update.
type CommunicationsGatewaysClientUpdateResponse struct {
	CommunicationsGateway
}

// NameAvailabilityClientCheckLocalResponse contains the response from method NameAvailabilityClient.CheckLocal.
type NameAvailabilityClientCheckLocalResponse struct {
	CheckNameAvailabilityResponse
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	OperationListResult
}

// TestLinesClientCreateOrUpdateResponse contains the response from method TestLinesClient.BeginCreateOrUpdate.
type TestLinesClientCreateOrUpdateResponse struct {
	TestLine
}

// TestLinesClientDeleteResponse contains the response from method TestLinesClient.BeginDelete.
type TestLinesClientDeleteResponse struct {
	// placeholder for future response values
}

// TestLinesClientGetResponse contains the response from method TestLinesClient.Get.
type TestLinesClientGetResponse struct {
	TestLine
}

// TestLinesClientListByCommunicationsGatewayResponse contains the response from method TestLinesClient.NewListByCommunicationsGatewayPager.
type TestLinesClientListByCommunicationsGatewayResponse struct {
	TestLineListResult
}

// TestLinesClientUpdateResponse contains the response from method TestLinesClient.Update.
type TestLinesClientUpdateResponse struct {
	TestLine
}