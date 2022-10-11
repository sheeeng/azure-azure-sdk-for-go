package insightsapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2022-06-01-preview/insights"
	"github.com/Azure/go-autorest/autorest"
)

// AutoscaleSettingsClientAPI contains the set of methods on the AutoscaleSettingsClient type.
type AutoscaleSettingsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, autoscaleSettingName string, parameters insights.AutoscaleSettingResource) (result insights.AutoscaleSettingResource, err error)
	Delete(ctx context.Context, resourceGroupName string, autoscaleSettingName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, autoscaleSettingName string) (result insights.AutoscaleSettingResource, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result insights.AutoscaleSettingResourceCollectionPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result insights.AutoscaleSettingResourceCollectionIterator, err error)
	ListBySubscription(ctx context.Context) (result insights.AutoscaleSettingResourceCollectionPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result insights.AutoscaleSettingResourceCollectionIterator, err error)
	Update(ctx context.Context, resourceGroupName string, autoscaleSettingName string, autoscaleSettingResource insights.AutoscaleSettingResourcePatch) (result insights.AutoscaleSettingResource, err error)
}

var _ AutoscaleSettingsClientAPI = (*insights.AutoscaleSettingsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result insights.OperationListResult, err error)
}

var _ OperationsClientAPI = (*insights.OperationsClient)(nil)

// AlertRuleIncidentsClientAPI contains the set of methods on the AlertRuleIncidentsClient type.
type AlertRuleIncidentsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, ruleName string, incidentName string) (result insights.Incident, err error)
	ListByAlertRule(ctx context.Context, resourceGroupName string, ruleName string) (result insights.IncidentListResult, err error)
}

var _ AlertRuleIncidentsClientAPI = (*insights.AlertRuleIncidentsClient)(nil)

// AlertRulesClientAPI contains the set of methods on the AlertRulesClient type.
type AlertRulesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, ruleName string, parameters insights.AlertRuleResource) (result insights.AlertRuleResource, err error)
	Delete(ctx context.Context, resourceGroupName string, ruleName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, ruleName string) (result insights.AlertRuleResource, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result insights.AlertRuleResourceCollection, err error)
	ListBySubscription(ctx context.Context) (result insights.AlertRuleResourceCollection, err error)
	Update(ctx context.Context, resourceGroupName string, ruleName string, alertRulesResource insights.AlertRuleResourcePatch) (result insights.AlertRuleResource, err error)
}

var _ AlertRulesClientAPI = (*insights.AlertRulesClient)(nil)

// LogProfilesClientAPI contains the set of methods on the LogProfilesClient type.
type LogProfilesClientAPI interface {
	CreateOrUpdate(ctx context.Context, logProfileName string, parameters insights.LogProfileResource) (result insights.LogProfileResource, err error)
	Delete(ctx context.Context, logProfileName string) (result autorest.Response, err error)
	Get(ctx context.Context, logProfileName string) (result insights.LogProfileResource, err error)
	List(ctx context.Context) (result insights.LogProfileCollection, err error)
	Update(ctx context.Context, logProfileName string, logProfilesResource insights.LogProfileResourcePatch) (result insights.LogProfileResource, err error)
}

var _ LogProfilesClientAPI = (*insights.LogProfilesClient)(nil)

// DiagnosticSettingsClientAPI contains the set of methods on the DiagnosticSettingsClient type.
type DiagnosticSettingsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceURI string, parameters insights.DiagnosticSettingsResource, name string) (result insights.DiagnosticSettingsResource, err error)
	Delete(ctx context.Context, resourceURI string, name string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceURI string, name string) (result insights.DiagnosticSettingsResource, err error)
	List(ctx context.Context, resourceURI string) (result insights.DiagnosticSettingsResourceCollection, err error)
}

var _ DiagnosticSettingsClientAPI = (*insights.DiagnosticSettingsClient)(nil)

// DiagnosticSettingsCategoryClientAPI contains the set of methods on the DiagnosticSettingsCategoryClient type.
type DiagnosticSettingsCategoryClientAPI interface {
	Get(ctx context.Context, resourceURI string, name string) (result insights.DiagnosticSettingsCategoryResource, err error)
	List(ctx context.Context, resourceURI string) (result insights.DiagnosticSettingsCategoryResourceCollection, err error)
}

var _ DiagnosticSettingsCategoryClientAPI = (*insights.DiagnosticSettingsCategoryClient)(nil)

// ActionGroupsClientAPI contains the set of methods on the ActionGroupsClient type.
type ActionGroupsClientAPI interface {
	CreateNotificationsAtActionGroupResourceLevel(ctx context.Context, resourceGroupName string, actionGroupName string, notificationRequest insights.NotificationRequestBody) (result insights.ActionGroupsCreateNotificationsAtActionGroupResourceLevelFuture, err error)
	CreateNotificationsAtResourceGroupLevel(ctx context.Context, resourceGroupName string, notificationRequest insights.NotificationRequestBody) (result insights.ActionGroupsCreateNotificationsAtResourceGroupLevelFuture, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, actionGroupName string, actionGroup insights.ActionGroupResource) (result insights.ActionGroupResource, err error)
	Delete(ctx context.Context, resourceGroupName string, actionGroupName string) (result autorest.Response, err error)
	EnableReceiver(ctx context.Context, resourceGroupName string, actionGroupName string, enableRequest insights.EnableRequest) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, actionGroupName string) (result insights.ActionGroupResource, err error)
	GetTestNotifications(ctx context.Context, notificationID string) (result insights.TestNotificationDetailsResponse, err error)
	GetTestNotificationsAtActionGroupResourceLevel(ctx context.Context, resourceGroupName string, actionGroupName string, notificationID string) (result insights.TestNotificationDetailsResponse, err error)
	GetTestNotificationsAtResourceGroupLevel(ctx context.Context, resourceGroupName string, notificationID string) (result insights.TestNotificationDetailsResponse, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result insights.ActionGroupList, err error)
	ListBySubscriptionID(ctx context.Context) (result insights.ActionGroupList, err error)
	PostTestNotifications(ctx context.Context, notificationRequest insights.NotificationRequestBody) (result insights.ActionGroupsPostTestNotificationsFuture, err error)
	Update(ctx context.Context, resourceGroupName string, actionGroupName string, actionGroupPatch insights.ActionGroupPatchBody) (result insights.ActionGroupResource, err error)
}

var _ ActionGroupsClientAPI = (*insights.ActionGroupsClient)(nil)

// ActivityLogsClientAPI contains the set of methods on the ActivityLogsClient type.
type ActivityLogsClientAPI interface {
	List(ctx context.Context, filter string, selectParameter string) (result insights.EventDataCollectionPage, err error)
	ListComplete(ctx context.Context, filter string, selectParameter string) (result insights.EventDataCollectionIterator, err error)
}

var _ ActivityLogsClientAPI = (*insights.ActivityLogsClient)(nil)

// EventCategoriesClientAPI contains the set of methods on the EventCategoriesClient type.
type EventCategoriesClientAPI interface {
	List(ctx context.Context) (result insights.EventCategoryCollection, err error)
}

var _ EventCategoriesClientAPI = (*insights.EventCategoriesClient)(nil)

// TenantActivityLogsClientAPI contains the set of methods on the TenantActivityLogsClient type.
type TenantActivityLogsClientAPI interface {
	List(ctx context.Context, filter string, selectParameter string) (result insights.EventDataCollectionPage, err error)
	ListComplete(ctx context.Context, filter string, selectParameter string) (result insights.EventDataCollectionIterator, err error)
}

var _ TenantActivityLogsClientAPI = (*insights.TenantActivityLogsClient)(nil)

// MetricDefinitionsClientAPI contains the set of methods on the MetricDefinitionsClient type.
type MetricDefinitionsClientAPI interface {
	List(ctx context.Context, resourceURI string, metricnamespace string) (result insights.MetricDefinitionCollection, err error)
}

var _ MetricDefinitionsClientAPI = (*insights.MetricDefinitionsClient)(nil)

// MetricsClientAPI contains the set of methods on the MetricsClient type.
type MetricsClientAPI interface {
	List(ctx context.Context, resourceURI string, timespan string, interval *string, metricnames string, aggregation string, top *int32, orderby string, filter string, resultType insights.ResultType, metricnamespace string) (result insights.Response, err error)
}

var _ MetricsClientAPI = (*insights.MetricsClient)(nil)

// BaselinesClientAPI contains the set of methods on the BaselinesClient type.
type BaselinesClientAPI interface {
	List(ctx context.Context, resourceURI string, metricnames string, metricnamespace string, timespan string, interval *string, aggregation string, sensitivities string, filter string, resultType insights.ResultType) (result insights.MetricBaselinesResponse, err error)
}

var _ BaselinesClientAPI = (*insights.BaselinesClient)(nil)

// MetricAlertsClientAPI contains the set of methods on the MetricAlertsClient type.
type MetricAlertsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, ruleName string, parameters insights.MetricAlertResource) (result insights.MetricAlertResource, err error)
	Delete(ctx context.Context, resourceGroupName string, ruleName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, ruleName string) (result insights.MetricAlertResource, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result insights.MetricAlertResourceCollection, err error)
	ListBySubscription(ctx context.Context) (result insights.MetricAlertResourceCollection, err error)
	Update(ctx context.Context, resourceGroupName string, ruleName string, parameters insights.MetricAlertResourcePatch) (result insights.MetricAlertResource, err error)
}

var _ MetricAlertsClientAPI = (*insights.MetricAlertsClient)(nil)

// MetricAlertsStatusClientAPI contains the set of methods on the MetricAlertsStatusClient type.
type MetricAlertsStatusClientAPI interface {
	List(ctx context.Context, resourceGroupName string, ruleName string) (result insights.MetricAlertStatusCollection, err error)
	ListByName(ctx context.Context, resourceGroupName string, ruleName string, statusName string) (result insights.MetricAlertStatusCollection, err error)
}

var _ MetricAlertsStatusClientAPI = (*insights.MetricAlertsStatusClient)(nil)

// ScheduledQueryRulesClientAPI contains the set of methods on the ScheduledQueryRulesClient type.
type ScheduledQueryRulesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, ruleName string, parameters insights.ScheduledQueryRuleResource) (result insights.ScheduledQueryRuleResource, err error)
	Delete(ctx context.Context, resourceGroupName string, ruleName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, ruleName string) (result insights.ScheduledQueryRuleResource, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result insights.ScheduledQueryRuleResourceCollectionPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result insights.ScheduledQueryRuleResourceCollectionIterator, err error)
	ListBySubscription(ctx context.Context) (result insights.ScheduledQueryRuleResourceCollectionPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result insights.ScheduledQueryRuleResourceCollectionIterator, err error)
	Update(ctx context.Context, resourceGroupName string, ruleName string, parameters insights.ScheduledQueryRuleResourcePatch) (result insights.ScheduledQueryRuleResource, err error)
}

var _ ScheduledQueryRulesClientAPI = (*insights.ScheduledQueryRulesClient)(nil)

// MetricNamespacesClientAPI contains the set of methods on the MetricNamespacesClient type.
type MetricNamespacesClientAPI interface {
	List(ctx context.Context, resourceURI string, startTime string) (result insights.MetricNamespaceCollection, err error)
}

var _ MetricNamespacesClientAPI = (*insights.MetricNamespacesClient)(nil)

// VMInsightsClientAPI contains the set of methods on the VMInsightsClient type.
type VMInsightsClientAPI interface {
	GetOnboardingStatus(ctx context.Context, resourceURI string) (result insights.VMInsightsOnboardingStatus, err error)
}

var _ VMInsightsClientAPI = (*insights.VMInsightsClient)(nil)

// PrivateLinkScopesClientAPI contains the set of methods on the PrivateLinkScopesClient type.
type PrivateLinkScopesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, scopeName string, azureMonitorPrivateLinkScopePayload insights.AzureMonitorPrivateLinkScope) (result insights.AzureMonitorPrivateLinkScope, err error)
	Delete(ctx context.Context, resourceGroupName string, scopeName string) (result insights.PrivateLinkScopesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, scopeName string) (result insights.AzureMonitorPrivateLinkScope, err error)
	List(ctx context.Context) (result insights.AzureMonitorPrivateLinkScopeListResultPage, err error)
	ListComplete(ctx context.Context) (result insights.AzureMonitorPrivateLinkScopeListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result insights.AzureMonitorPrivateLinkScopeListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result insights.AzureMonitorPrivateLinkScopeListResultIterator, err error)
	UpdateTags(ctx context.Context, resourceGroupName string, scopeName string, privateLinkScopeTags insights.TagsResource) (result insights.AzureMonitorPrivateLinkScope, err error)
}

var _ PrivateLinkScopesClientAPI = (*insights.PrivateLinkScopesClient)(nil)

// PrivateLinkScopeOperationStatusClientAPI contains the set of methods on the PrivateLinkScopeOperationStatusClient type.
type PrivateLinkScopeOperationStatusClientAPI interface {
	Get(ctx context.Context, asyncOperationID string, resourceGroupName string) (result insights.OperationStatus, err error)
}

var _ PrivateLinkScopeOperationStatusClientAPI = (*insights.PrivateLinkScopeOperationStatusClient)(nil)

// PrivateLinkResourcesClientAPI contains the set of methods on the PrivateLinkResourcesClient type.
type PrivateLinkResourcesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, scopeName string, groupName string) (result insights.PrivateLinkResource, err error)
	ListByPrivateLinkScope(ctx context.Context, resourceGroupName string, scopeName string) (result insights.PrivateLinkResourceListResultPage, err error)
	ListByPrivateLinkScopeComplete(ctx context.Context, resourceGroupName string, scopeName string) (result insights.PrivateLinkResourceListResultIterator, err error)
}

var _ PrivateLinkResourcesClientAPI = (*insights.PrivateLinkResourcesClient)(nil)

// PrivateEndpointConnectionsClientAPI contains the set of methods on the PrivateEndpointConnectionsClient type.
type PrivateEndpointConnectionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, scopeName string, privateEndpointConnectionName string, parameters insights.PrivateEndpointConnection) (result insights.PrivateEndpointConnectionsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, scopeName string, privateEndpointConnectionName string) (result insights.PrivateEndpointConnectionsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, scopeName string, privateEndpointConnectionName string) (result insights.PrivateEndpointConnection, err error)
	ListByPrivateLinkScope(ctx context.Context, resourceGroupName string, scopeName string) (result insights.PrivateEndpointConnectionListResultPage, err error)
	ListByPrivateLinkScopeComplete(ctx context.Context, resourceGroupName string, scopeName string) (result insights.PrivateEndpointConnectionListResultIterator, err error)
}

var _ PrivateEndpointConnectionsClientAPI = (*insights.PrivateEndpointConnectionsClient)(nil)

// PrivateLinkScopedResourcesClientAPI contains the set of methods on the PrivateLinkScopedResourcesClient type.
type PrivateLinkScopedResourcesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, scopeName string, name string, parameters insights.ScopedResource) (result insights.PrivateLinkScopedResourcesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, scopeName string, name string) (result insights.PrivateLinkScopedResourcesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, scopeName string, name string) (result insights.ScopedResource, err error)
	ListByPrivateLinkScope(ctx context.Context, resourceGroupName string, scopeName string) (result insights.ScopedResourceListResultPage, err error)
	ListByPrivateLinkScopeComplete(ctx context.Context, resourceGroupName string, scopeName string) (result insights.ScopedResourceListResultIterator, err error)
}

var _ PrivateLinkScopedResourcesClientAPI = (*insights.PrivateLinkScopedResourcesClient)(nil)

// ActivityLogAlertsClientAPI contains the set of methods on the ActivityLogAlertsClient type.
type ActivityLogAlertsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, activityLogAlertName string, activityLogAlertRule insights.ActivityLogAlertResource) (result insights.ActivityLogAlertResource, err error)
	Delete(ctx context.Context, resourceGroupName string, activityLogAlertName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, activityLogAlertName string) (result insights.ActivityLogAlertResource, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result insights.AlertRuleListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result insights.AlertRuleListIterator, err error)
	ListBySubscriptionID(ctx context.Context) (result insights.AlertRuleListPage, err error)
	ListBySubscriptionIDComplete(ctx context.Context) (result insights.AlertRuleListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, activityLogAlertName string, activityLogAlertRulePatch insights.AlertRulePatchObject) (result insights.ActivityLogAlertResource, err error)
}

var _ ActivityLogAlertsClientAPI = (*insights.ActivityLogAlertsClient)(nil)

// DataCollectionEndpointsClientAPI contains the set of methods on the DataCollectionEndpointsClient type.
type DataCollectionEndpointsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, dataCollectionEndpointName string, body *insights.DataCollectionEndpointResource) (result insights.DataCollectionEndpointResource, err error)
	Delete(ctx context.Context, resourceGroupName string, dataCollectionEndpointName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, dataCollectionEndpointName string) (result insights.DataCollectionEndpointResource, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result insights.DataCollectionEndpointResourceListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result insights.DataCollectionEndpointResourceListResultIterator, err error)
	ListBySubscription(ctx context.Context) (result insights.DataCollectionEndpointResourceListResultPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result insights.DataCollectionEndpointResourceListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, dataCollectionEndpointName string, body *insights.ResourceForUpdate) (result insights.DataCollectionEndpointResource, err error)
}

var _ DataCollectionEndpointsClientAPI = (*insights.DataCollectionEndpointsClient)(nil)

// DataCollectionRuleAssociationsClientAPI contains the set of methods on the DataCollectionRuleAssociationsClient type.
type DataCollectionRuleAssociationsClientAPI interface {
	Create(ctx context.Context, resourceURI string, associationName string, body *insights.DataCollectionRuleAssociationProxyOnlyResource) (result insights.DataCollectionRuleAssociationProxyOnlyResource, err error)
	Delete(ctx context.Context, resourceURI string, associationName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceURI string, associationName string) (result insights.DataCollectionRuleAssociationProxyOnlyResource, err error)
	ListByDataCollectionEndpoint(ctx context.Context, resourceGroupName string, dataCollectionEndpointName string) (result insights.DataCollectionRuleAssociationProxyOnlyResourceListResultPage, err error)
	ListByDataCollectionEndpointComplete(ctx context.Context, resourceGroupName string, dataCollectionEndpointName string) (result insights.DataCollectionRuleAssociationProxyOnlyResourceListResultIterator, err error)
	ListByResource(ctx context.Context, resourceURI string) (result insights.DataCollectionRuleAssociationProxyOnlyResourceListResultPage, err error)
	ListByResourceComplete(ctx context.Context, resourceURI string) (result insights.DataCollectionRuleAssociationProxyOnlyResourceListResultIterator, err error)
	ListByRule(ctx context.Context, resourceGroupName string, dataCollectionRuleName string) (result insights.DataCollectionRuleAssociationProxyOnlyResourceListResultPage, err error)
	ListByRuleComplete(ctx context.Context, resourceGroupName string, dataCollectionRuleName string) (result insights.DataCollectionRuleAssociationProxyOnlyResourceListResultIterator, err error)
}

var _ DataCollectionRuleAssociationsClientAPI = (*insights.DataCollectionRuleAssociationsClient)(nil)

// DataCollectionRulesClientAPI contains the set of methods on the DataCollectionRulesClient type.
type DataCollectionRulesClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, dataCollectionRuleName string, body *insights.DataCollectionRuleResource) (result insights.DataCollectionRuleResource, err error)
	Delete(ctx context.Context, resourceGroupName string, dataCollectionRuleName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, dataCollectionRuleName string) (result insights.DataCollectionRuleResource, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result insights.DataCollectionRuleResourceListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result insights.DataCollectionRuleResourceListResultIterator, err error)
	ListBySubscription(ctx context.Context) (result insights.DataCollectionRuleResourceListResultPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result insights.DataCollectionRuleResourceListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, dataCollectionRuleName string, body *insights.ResourceForUpdate) (result insights.DataCollectionRuleResource, err error)
}

var _ DataCollectionRulesClientAPI = (*insights.DataCollectionRulesClient)(nil)
