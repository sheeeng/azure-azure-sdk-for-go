//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armloadtestservice

import "time"

// CheckQuotaAvailabilityResponse - Check quota availability response object.
type CheckQuotaAvailabilityResponse struct {
	// Check quota availability response properties.
	Properties *CheckQuotaAvailabilityResponseProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// CheckQuotaAvailabilityResponseProperties - Check quota availability response properties.
type CheckQuotaAvailabilityResponseProperties struct {
	// Message indicating additional details to add to quota support request.
	AvailabilityStatus *string `json:"availabilityStatus,omitempty"`

	// True/False indicating whether the quota request be granted based on availability.
	IsAvailable *bool `json:"isAvailable,omitempty"`
}

// EncryptionProperties - Key and identity details for Customer Managed Key encryption of load test resource
type EncryptionProperties struct {
	// All identity configuration for Customer-managed key settings defining which identity should be used to auth to Key Vault.
	Identity *EncryptionPropertiesIdentity `json:"identity,omitempty"`

	// key encryption key Url, versioned. Ex: https://contosovault.vault.azure.net/keys/contosokek/562a4bb76b524a1493a6afe8e536ee78
	// or https://contosovault.vault.azure.net/keys/contosokek.
	KeyURL *string `json:"keyUrl,omitempty"`
}

// EncryptionPropertiesIdentity - All identity configuration for Customer-managed key settings defining which identity should
// be used to auth to Key Vault.
type EncryptionPropertiesIdentity struct {
	// user assigned identity to use for accessing key encryption key Url. Ex: /subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/
	// /providers/Microsoft.ManagedIdentity/userAssignedIdentities/myId
	ResourceID *string `json:"resourceId,omitempty"`

	// Managed identity type to use for accessing encryption key Url
	Type *Type `json:"type,omitempty"`
}

// EndpointDependency - A domain name and connection details used to access a dependency.
type EndpointDependency struct {
	// READ-ONLY; Human-readable supplemental information about the dependency and when it is applicable.
	Description *string `json:"description,omitempty" azure:"ro"`

	// READ-ONLY; The domain name of the dependency. Domain names may be fully qualified or may contain a * wildcard.
	DomainName *string `json:"domainName,omitempty" azure:"ro"`

	// READ-ONLY; The list of connection details for this endpoint.
	EndpointDetails []*EndpointDetail `json:"endpointDetails,omitempty" azure:"ro"`
}

// EndpointDetail - Details about the connection between the Batch service and the endpoint.
type EndpointDetail struct {
	// READ-ONLY; The port an endpoint is connected to.
	Port *int32 `json:"port,omitempty" azure:"ro"`
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info interface{} `json:"info,omitempty" azure:"ro"`

	// READ-ONLY; The additional info type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ErrorDetail - The error detail.
type ErrorDetail struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo `json:"additionalInfo,omitempty" azure:"ro"`

	// READ-ONLY; The error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; The error details.
	Details []*ErrorDetail `json:"details,omitempty" azure:"ro"`

	// READ-ONLY; The error message.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; The error target.
	Target *string `json:"target,omitempty" azure:"ro"`
}

// ErrorResponse - Common error response for all Azure Resource Manager APIs to return error details for failed operations.
// (This also follows the OData error response format.).
type ErrorResponse struct {
	// The error object.
	Error *ErrorDetail `json:"error,omitempty"`
}

// LoadTestProperties - LoadTest resource properties.
type LoadTestProperties struct {
	// Description of the resource.
	Description *string `json:"description,omitempty"`

	// CMK Encryption property.
	Encryption *EncryptionProperties `json:"encryption,omitempty"`

	// READ-ONLY; Resource data plane URI.
	DataPlaneURI *string `json:"dataPlaneURI,omitempty" azure:"ro"`

	// READ-ONLY; Resource provisioning state.
	ProvisioningState *ResourceState `json:"provisioningState,omitempty" azure:"ro"`
}

// LoadTestResource - LoadTest details
type LoadTestResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// The type of identity used for the resource.
	Identity *ManagedServiceIdentity `json:"identity,omitempty"`

	// Load Test resource properties
	Properties *LoadTestProperties `json:"properties,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// LoadTestResourcePageList - List of resources page result.
type LoadTestResourcePageList struct {
	// Link to next page of resources.
	NextLink *string `json:"nextLink,omitempty"`

	// List of resources in current page.
	Value []*LoadTestResource `json:"value,omitempty"`
}

// LoadTestResourcePatchRequestBody - LoadTest resource patch request body.
type LoadTestResourcePatchRequestBody struct {
	// The type of identity used for the resource.
	Identity *ManagedServiceIdentity `json:"identity,omitempty"`

	// Load Test resource properties
	Properties *LoadTestResourcePatchRequestBodyProperties `json:"properties,omitempty"`

	// Resource tags.
	Tags interface{} `json:"tags,omitempty"`
}

// LoadTestResourcePatchRequestBodyProperties - Load Test resource properties
type LoadTestResourcePatchRequestBodyProperties struct {
	// Description of the resource.
	Description *string `json:"description,omitempty"`

	// CMK Encryption property.
	Encryption *EncryptionProperties `json:"encryption,omitempty"`
}

// LoadTestsClientBeginCreateOrUpdateOptions contains the optional parameters for the LoadTestsClient.BeginCreateOrUpdate
// method.
type LoadTestsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// LoadTestsClientBeginDeleteOptions contains the optional parameters for the LoadTestsClient.BeginDelete method.
type LoadTestsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// LoadTestsClientBeginUpdateOptions contains the optional parameters for the LoadTestsClient.BeginUpdate method.
type LoadTestsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// LoadTestsClientGetOptions contains the optional parameters for the LoadTestsClient.Get method.
type LoadTestsClientGetOptions struct {
	// placeholder for future optional parameters
}

// LoadTestsClientListByResourceGroupOptions contains the optional parameters for the LoadTestsClient.ListByResourceGroup
// method.
type LoadTestsClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// LoadTestsClientListBySubscriptionOptions contains the optional parameters for the LoadTestsClient.ListBySubscription method.
type LoadTestsClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// LoadTestsClientListOutboundNetworkDependenciesEndpointsOptions contains the optional parameters for the LoadTestsClient.ListOutboundNetworkDependenciesEndpoints
// method.
type LoadTestsClientListOutboundNetworkDependenciesEndpointsOptions struct {
	// placeholder for future optional parameters
}

// ManagedServiceIdentity - Managed service identity (system assigned and/or user assigned identities)
type ManagedServiceIdentity struct {
	// REQUIRED; Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
	Type *ManagedServiceIdentityType `json:"type,omitempty"`

	// The set of user assigned identities associated with the resource. The userAssignedIdentities dictionary keys will be ARM
	// resource ids in the form:
	// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}.
	// The dictionary values can be empty objects ({}) in
	// requests.
	UserAssignedIdentities map[string]*UserAssignedIdentity `json:"userAssignedIdentities,omitempty"`

	// READ-ONLY; The service principal ID of the system assigned identity. This property will only be provided for a system assigned
	// identity.
	PrincipalID *string `json:"principalId,omitempty" azure:"ro"`

	// READ-ONLY; The tenant ID of the system assigned identity. This property will only be provided for a system assigned identity.
	TenantID *string `json:"tenantId,omitempty" azure:"ro"`
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay `json:"display,omitempty"`

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType `json:"actionType,omitempty" azure:"ro"`

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
	// operations.
	IsDataAction *bool `json:"isDataAction,omitempty" azure:"ro"`

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin `json:"origin,omitempty" azure:"ro"`
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string `json:"description,omitempty" azure:"ro"`

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
	// Machine", "Restart Virtual Machine".
	Operation *string `json:"operation,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
	// Compute".
	Provider *string `json:"provider,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
	// Schedule Collections".
	Resource *string `json:"resource,omitempty" azure:"ro"`
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation `json:"value,omitempty" azure:"ro"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// OutboundEnvironmentEndpoint - A collection of related endpoints from the same service for which the Batch service requires
// outbound access.
type OutboundEnvironmentEndpoint struct {
	// READ-ONLY; The type of service that Azure Load Testing connects to.
	Category *string `json:"category,omitempty" azure:"ro"`

	// READ-ONLY; The endpoints for this service to which the Batch service makes outbound calls.
	Endpoints []*EndpointDependency `json:"endpoints,omitempty" azure:"ro"`
}

// OutboundEnvironmentEndpointCollection - Values returned by the List operation.
type OutboundEnvironmentEndpointCollection struct {
	// The continuation token.
	NextLink *string `json:"nextLink,omitempty"`

	// READ-ONLY; The collection of outbound network dependency endpoints returned by the listing operation.
	Value []*OutboundEnvironmentEndpoint `json:"value,omitempty" azure:"ro"`
}

// QuotaBucketRequest - Request object of new quota for a quota bucket.
type QuotaBucketRequest struct {
	// New quota request request properties.
	Properties *QuotaBucketRequestProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// QuotaBucketRequestProperties - New quota request request properties.
type QuotaBucketRequestProperties struct {
	// Current quota limit of the quota bucket.
	CurrentQuota *int32 `json:"currentQuota,omitempty"`

	// Current quota usage of the quota bucket.
	CurrentUsage *int32 `json:"currentUsage,omitempty"`

	// Dimensions for new quota request.
	Dimensions *QuotaBucketRequestPropertiesDimensions `json:"dimensions,omitempty"`

	// New quota limit of the quota bucket.
	NewQuota *int32 `json:"newQuota,omitempty"`
}

// QuotaBucketRequestPropertiesDimensions - Dimensions for new quota request.
type QuotaBucketRequestPropertiesDimensions struct {
	// Location dimension for new quota request of the quota bucket.
	Location *string `json:"location,omitempty"`

	// Subscription Id dimension for new quota request of the quota bucket.
	SubscriptionID *string `json:"subscriptionId,omitempty"`
}

// QuotaResource - Quota bucket details object.
type QuotaResource struct {
	// Quota bucket resource properties.
	Properties *QuotaResourceProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// QuotaResourceList - List of quota bucket objects. It contains a URL link to get the next set of results.
type QuotaResourceList struct {
	// READ-ONLY; URL to get the next set of quota bucket objects results (if there are any).
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; List of quota bucket objects provided by the loadtestservice.
	Value []*QuotaResource `json:"value,omitempty" azure:"ro"`
}

// QuotaResourceProperties - Quota bucket resource properties.
type QuotaResourceProperties struct {
	// Current quota limit of the quota bucket.
	Limit *int32 `json:"limit,omitempty"`

	// Current quota usage of the quota bucket.
	Usage *int32 `json:"usage,omitempty"`

	// READ-ONLY; Resource provisioning state.
	ProvisioningState *ResourceState `json:"provisioningState,omitempty" azure:"ro"`
}

// QuotasClientCheckAvailabilityOptions contains the optional parameters for the QuotasClient.CheckAvailability method.
type QuotasClientCheckAvailabilityOptions struct {
	// placeholder for future optional parameters
}

// QuotasClientGetOptions contains the optional parameters for the QuotasClient.Get method.
type QuotasClientGetOptions struct {
	// placeholder for future optional parameters
}

// QuotasClientListOptions contains the optional parameters for the QuotasClient.List method.
type QuotasClientListOptions struct {
	// placeholder for future optional parameters
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// The type of identity that created the resource.
	CreatedByType *CreatedByType `json:"createdByType,omitempty"`

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`

	// The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType `json:"lastModifiedByType,omitempty"`
}

// TrackedResource - The resource model definition for an Azure Resource Manager tracked top level resource which has 'tags'
// and a 'location'
type TrackedResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// UserAssignedIdentity - User assigned identity properties
type UserAssignedIdentity struct {
	// READ-ONLY; The client ID of the assigned identity.
	ClientID *string `json:"clientId,omitempty" azure:"ro"`

	// READ-ONLY; The principal ID of the assigned identity.
	PrincipalID *string `json:"principalId,omitempty" azure:"ro"`
}