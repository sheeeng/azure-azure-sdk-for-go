//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhdinsightcontainers

// AvailableClusterPoolVersionsClientListByLocationOptions contains the optional parameters for the AvailableClusterPoolVersionsClient.NewListByLocationPager
// method.
type AvailableClusterPoolVersionsClientListByLocationOptions struct {
	// placeholder for future optional parameters
}

// AvailableClusterVersionsClientListByLocationOptions contains the optional parameters for the AvailableClusterVersionsClient.NewListByLocationPager
// method.
type AvailableClusterVersionsClientListByLocationOptions struct {
	// placeholder for future optional parameters
}

// ClusterAvailableUpgradesClientListOptions contains the optional parameters for the ClusterAvailableUpgradesClient.NewListPager
// method.
type ClusterAvailableUpgradesClientListOptions struct {
	// placeholder for future optional parameters
}

// ClusterJobsClientBeginRunJobOptions contains the optional parameters for the ClusterJobsClient.BeginRunJob method.
type ClusterJobsClientBeginRunJobOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ClusterJobsClientListOptions contains the optional parameters for the ClusterJobsClient.NewListPager method.
type ClusterJobsClientListOptions struct {
	// The system query option to filter job returned in the response. Allowed value is 'jobName eq {jobName}' or 'jarName eq
	// {jarName}'.
	Filter *string
}

// ClusterLibrariesClientBeginManageLibrariesOptions contains the optional parameters for the ClusterLibrariesClient.BeginManageLibraries
// method.
type ClusterLibrariesClientBeginManageLibrariesOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ClusterLibrariesClientListOptions contains the optional parameters for the ClusterLibrariesClient.NewListPager method.
type ClusterLibrariesClientListOptions struct {
	// placeholder for future optional parameters
}

// ClusterPoolAvailableUpgradesClientListOptions contains the optional parameters for the ClusterPoolAvailableUpgradesClient.NewListPager
// method.
type ClusterPoolAvailableUpgradesClientListOptions struct {
	// placeholder for future optional parameters
}

// ClusterPoolUpgradeHistoriesClientListOptions contains the optional parameters for the ClusterPoolUpgradeHistoriesClient.NewListPager
// method.
type ClusterPoolUpgradeHistoriesClientListOptions struct {
	// placeholder for future optional parameters
}

// ClusterPoolsClientBeginCreateOrUpdateOptions contains the optional parameters for the ClusterPoolsClient.BeginCreateOrUpdate
// method.
type ClusterPoolsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ClusterPoolsClientBeginDeleteOptions contains the optional parameters for the ClusterPoolsClient.BeginDelete method.
type ClusterPoolsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ClusterPoolsClientBeginUpdateTagsOptions contains the optional parameters for the ClusterPoolsClient.BeginUpdateTags method.
type ClusterPoolsClientBeginUpdateTagsOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ClusterPoolsClientBeginUpgradeOptions contains the optional parameters for the ClusterPoolsClient.BeginUpgrade method.
type ClusterPoolsClientBeginUpgradeOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ClusterPoolsClientGetOptions contains the optional parameters for the ClusterPoolsClient.Get method.
type ClusterPoolsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ClusterPoolsClientListByResourceGroupOptions contains the optional parameters for the ClusterPoolsClient.NewListByResourceGroupPager
// method.
type ClusterPoolsClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// ClusterPoolsClientListBySubscriptionOptions contains the optional parameters for the ClusterPoolsClient.NewListBySubscriptionPager
// method.
type ClusterPoolsClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// ClusterUpgradeHistoriesClientListOptions contains the optional parameters for the ClusterUpgradeHistoriesClient.NewListPager
// method.
type ClusterUpgradeHistoriesClientListOptions struct {
	// placeholder for future optional parameters
}

// ClustersClientBeginCreateOptions contains the optional parameters for the ClustersClient.BeginCreate method.
type ClustersClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ClustersClientBeginDeleteOptions contains the optional parameters for the ClustersClient.BeginDelete method.
type ClustersClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ClustersClientBeginResizeOptions contains the optional parameters for the ClustersClient.BeginResize method.
type ClustersClientBeginResizeOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ClustersClientBeginUpdateOptions contains the optional parameters for the ClustersClient.BeginUpdate method.
type ClustersClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ClustersClientBeginUpgradeManualRollbackOptions contains the optional parameters for the ClustersClient.BeginUpgradeManualRollback
// method.
type ClustersClientBeginUpgradeManualRollbackOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ClustersClientBeginUpgradeOptions contains the optional parameters for the ClustersClient.BeginUpgrade method.
type ClustersClientBeginUpgradeOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ClustersClientGetInstanceViewOptions contains the optional parameters for the ClustersClient.GetInstanceView method.
type ClustersClientGetInstanceViewOptions struct {
	// placeholder for future optional parameters
}

// ClustersClientGetOptions contains the optional parameters for the ClustersClient.Get method.
type ClustersClientGetOptions struct {
	// placeholder for future optional parameters
}

// ClustersClientListByClusterPoolNameOptions contains the optional parameters for the ClustersClient.NewListByClusterPoolNamePager
// method.
type ClustersClientListByClusterPoolNameOptions struct {
	// placeholder for future optional parameters
}

// ClustersClientListInstanceViewsOptions contains the optional parameters for the ClustersClient.NewListInstanceViewsPager
// method.
type ClustersClientListInstanceViewsOptions struct {
	// placeholder for future optional parameters
}

// ClustersClientListServiceConfigsOptions contains the optional parameters for the ClustersClient.NewListServiceConfigsPager
// method.
type ClustersClientListServiceConfigsOptions struct {
	// placeholder for future optional parameters
}

// LocationsClientCheckNameAvailabilityOptions contains the optional parameters for the LocationsClient.CheckNameAvailability
// method.
type LocationsClientCheckNameAvailabilityOptions struct {
	// placeholder for future optional parameters
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}
