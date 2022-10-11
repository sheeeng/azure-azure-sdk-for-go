package sql

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

// SynapseLinkWorkspacesClient is the the Azure SQL Database management API provides a RESTful set of web services that
// interact with Azure SQL Database services to manage your databases. The API enables you to create, retrieve, update,
// and delete databases.
type SynapseLinkWorkspacesClient struct {
	BaseClient
}

// NewSynapseLinkWorkspacesClient creates an instance of the SynapseLinkWorkspacesClient client.
func NewSynapseLinkWorkspacesClient(subscriptionID string) SynapseLinkWorkspacesClient {
	return NewSynapseLinkWorkspacesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSynapseLinkWorkspacesClientWithBaseURI creates an instance of the SynapseLinkWorkspacesClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewSynapseLinkWorkspacesClientWithBaseURI(baseURI string, subscriptionID string) SynapseLinkWorkspacesClient {
	return SynapseLinkWorkspacesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListByDatabase gets all synapselink workspaces for a database.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// databaseName - the name of the database.
func (client SynapseLinkWorkspacesClient) ListByDatabase(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result SynapseLinkWorkspaceListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SynapseLinkWorkspacesClient.ListByDatabase")
		defer func() {
			sc := -1
			if result.slwlr.Response.Response != nil {
				sc = result.slwlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByDatabaseNextResults
	req, err := client.ListByDatabasePreparer(ctx, resourceGroupName, serverName, databaseName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SynapseLinkWorkspacesClient", "ListByDatabase", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByDatabaseSender(req)
	if err != nil {
		result.slwlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.SynapseLinkWorkspacesClient", "ListByDatabase", resp, "Failure sending request")
		return
	}

	result.slwlr, err = client.ListByDatabaseResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SynapseLinkWorkspacesClient", "ListByDatabase", resp, "Failure responding to request")
		return
	}
	if result.slwlr.hasNextLink() && result.slwlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByDatabasePreparer prepares the ListByDatabase request.
func (client SynapseLinkWorkspacesClient) ListByDatabasePreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/linkWorkspaces", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByDatabaseSender sends the ListByDatabase request. The method will close the
// http.Response Body if it receives an error.
func (client SynapseLinkWorkspacesClient) ListByDatabaseSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByDatabaseResponder handles the response to the ListByDatabase request. The method always
// closes the http.Response Body.
func (client SynapseLinkWorkspacesClient) ListByDatabaseResponder(resp *http.Response) (result SynapseLinkWorkspaceListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByDatabaseNextResults retrieves the next set of results, if any.
func (client SynapseLinkWorkspacesClient) listByDatabaseNextResults(ctx context.Context, lastResults SynapseLinkWorkspaceListResult) (result SynapseLinkWorkspaceListResult, err error) {
	req, err := lastResults.synapseLinkWorkspaceListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.SynapseLinkWorkspacesClient", "listByDatabaseNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByDatabaseSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.SynapseLinkWorkspacesClient", "listByDatabaseNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByDatabaseResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SynapseLinkWorkspacesClient", "listByDatabaseNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByDatabaseComplete enumerates all values, automatically crossing page boundaries as required.
func (client SynapseLinkWorkspacesClient) ListByDatabaseComplete(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result SynapseLinkWorkspaceListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SynapseLinkWorkspacesClient.ListByDatabase")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByDatabase(ctx, resourceGroupName, serverName, databaseName)
	return
}
