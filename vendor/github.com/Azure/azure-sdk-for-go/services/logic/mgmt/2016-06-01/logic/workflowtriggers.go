package logic

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// WorkflowTriggersClient is the REST API for Azure Logic Apps.
type WorkflowTriggersClient struct {
	BaseClient
}

// NewWorkflowTriggersClient creates an instance of the WorkflowTriggersClient client.
func NewWorkflowTriggersClient(subscriptionID string) WorkflowTriggersClient {
	return NewWorkflowTriggersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWorkflowTriggersClientWithBaseURI creates an instance of the WorkflowTriggersClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewWorkflowTriggersClientWithBaseURI(baseURI string, subscriptionID string) WorkflowTriggersClient {
	return WorkflowTriggersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets a workflow trigger.
// Parameters:
// resourceGroupName - the resource group name.
// workflowName - the workflow name.
// triggerName - the workflow trigger name.
func (client WorkflowTriggersClient) Get(ctx context.Context, resourceGroupName string, workflowName string, triggerName string) (result WorkflowTrigger, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkflowTriggersClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, workflowName, triggerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowTriggersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.WorkflowTriggersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowTriggersClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client WorkflowTriggersClient) GetPreparer(ctx context.Context, resourceGroupName string, workflowName string, triggerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"triggerName":       autorest.Encode("path", triggerName),
		"workflowName":      autorest.Encode("path", workflowName),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/triggers/{triggerName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowTriggersClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client WorkflowTriggersClient) GetResponder(resp *http.Response) (result WorkflowTrigger, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetSchemaJSON get the trigger schema as JSON.
// Parameters:
// resourceGroupName - the resource group name.
// workflowName - the workflow name.
// triggerName - the workflow trigger name.
func (client WorkflowTriggersClient) GetSchemaJSON(ctx context.Context, resourceGroupName string, workflowName string, triggerName string) (result JSONSchema, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkflowTriggersClient.GetSchemaJSON")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetSchemaJSONPreparer(ctx, resourceGroupName, workflowName, triggerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowTriggersClient", "GetSchemaJSON", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSchemaJSONSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.WorkflowTriggersClient", "GetSchemaJSON", resp, "Failure sending request")
		return
	}

	result, err = client.GetSchemaJSONResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowTriggersClient", "GetSchemaJSON", resp, "Failure responding to request")
	}

	return
}

// GetSchemaJSONPreparer prepares the GetSchemaJSON request.
func (client WorkflowTriggersClient) GetSchemaJSONPreparer(ctx context.Context, resourceGroupName string, workflowName string, triggerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"triggerName":       autorest.Encode("path", triggerName),
		"workflowName":      autorest.Encode("path", workflowName),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/triggers/{triggerName}/schemas/json", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSchemaJSONSender sends the GetSchemaJSON request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowTriggersClient) GetSchemaJSONSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetSchemaJSONResponder handles the response to the GetSchemaJSON request. The method always
// closes the http.Response Body.
func (client WorkflowTriggersClient) GetSchemaJSONResponder(resp *http.Response) (result JSONSchema, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets a list of workflow triggers.
// Parameters:
// resourceGroupName - the resource group name.
// workflowName - the workflow name.
// top - the number of items to be included in the result.
// filter - the filter to apply on the operation.
func (client WorkflowTriggersClient) List(ctx context.Context, resourceGroupName string, workflowName string, top *int32, filter string) (result WorkflowTriggerListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkflowTriggersClient.List")
		defer func() {
			sc := -1
			if result.wtlr.Response.Response != nil {
				sc = result.wtlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, workflowName, top, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowTriggersClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.wtlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.WorkflowTriggersClient", "List", resp, "Failure sending request")
		return
	}

	result.wtlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowTriggersClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client WorkflowTriggersClient) ListPreparer(ctx context.Context, resourceGroupName string, workflowName string, top *int32, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workflowName":      autorest.Encode("path", workflowName),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/triggers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowTriggersClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client WorkflowTriggersClient) ListResponder(resp *http.Response) (result WorkflowTriggerListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client WorkflowTriggersClient) listNextResults(ctx context.Context, lastResults WorkflowTriggerListResult) (result WorkflowTriggerListResult, err error) {
	req, err := lastResults.workflowTriggerListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic.WorkflowTriggersClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic.WorkflowTriggersClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowTriggersClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client WorkflowTriggersClient) ListComplete(ctx context.Context, resourceGroupName string, workflowName string, top *int32, filter string) (result WorkflowTriggerListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkflowTriggersClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, workflowName, top, filter)
	return
}

// ListCallbackURL get the callback URL for a workflow trigger.
// Parameters:
// resourceGroupName - the resource group name.
// workflowName - the workflow name.
// triggerName - the workflow trigger name.
func (client WorkflowTriggersClient) ListCallbackURL(ctx context.Context, resourceGroupName string, workflowName string, triggerName string) (result WorkflowTriggerCallbackURL, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkflowTriggersClient.ListCallbackURL")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListCallbackURLPreparer(ctx, resourceGroupName, workflowName, triggerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowTriggersClient", "ListCallbackURL", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListCallbackURLSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.WorkflowTriggersClient", "ListCallbackURL", resp, "Failure sending request")
		return
	}

	result, err = client.ListCallbackURLResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowTriggersClient", "ListCallbackURL", resp, "Failure responding to request")
	}

	return
}

// ListCallbackURLPreparer prepares the ListCallbackURL request.
func (client WorkflowTriggersClient) ListCallbackURLPreparer(ctx context.Context, resourceGroupName string, workflowName string, triggerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"triggerName":       autorest.Encode("path", triggerName),
		"workflowName":      autorest.Encode("path", workflowName),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/triggers/{triggerName}/listCallbackUrl", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListCallbackURLSender sends the ListCallbackURL request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowTriggersClient) ListCallbackURLSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListCallbackURLResponder handles the response to the ListCallbackURL request. The method always
// closes the http.Response Body.
func (client WorkflowTriggersClient) ListCallbackURLResponder(resp *http.Response) (result WorkflowTriggerCallbackURL, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Reset resets a workflow trigger.
// Parameters:
// resourceGroupName - the resource group name.
// workflowName - the workflow name.
// triggerName - the workflow trigger name.
func (client WorkflowTriggersClient) Reset(ctx context.Context, resourceGroupName string, workflowName string, triggerName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkflowTriggersClient.Reset")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ResetPreparer(ctx, resourceGroupName, workflowName, triggerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowTriggersClient", "Reset", nil, "Failure preparing request")
		return
	}

	resp, err := client.ResetSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "logic.WorkflowTriggersClient", "Reset", resp, "Failure sending request")
		return
	}

	result, err = client.ResetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowTriggersClient", "Reset", resp, "Failure responding to request")
	}

	return
}

// ResetPreparer prepares the Reset request.
func (client WorkflowTriggersClient) ResetPreparer(ctx context.Context, resourceGroupName string, workflowName string, triggerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"triggerName":       autorest.Encode("path", triggerName),
		"workflowName":      autorest.Encode("path", workflowName),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/triggers/{triggerName}/reset", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ResetSender sends the Reset request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowTriggersClient) ResetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ResetResponder handles the response to the Reset request. The method always
// closes the http.Response Body.
func (client WorkflowTriggersClient) ResetResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Run runs a workflow trigger.
// Parameters:
// resourceGroupName - the resource group name.
// workflowName - the workflow name.
// triggerName - the workflow trigger name.
func (client WorkflowTriggersClient) Run(ctx context.Context, resourceGroupName string, workflowName string, triggerName string) (result SetObject, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkflowTriggersClient.Run")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RunPreparer(ctx, resourceGroupName, workflowName, triggerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowTriggersClient", "Run", nil, "Failure preparing request")
		return
	}

	resp, err := client.RunSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.WorkflowTriggersClient", "Run", resp, "Failure sending request")
		return
	}

	result, err = client.RunResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowTriggersClient", "Run", resp, "Failure responding to request")
	}

	return
}

// RunPreparer prepares the Run request.
func (client WorkflowTriggersClient) RunPreparer(ctx context.Context, resourceGroupName string, workflowName string, triggerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"triggerName":       autorest.Encode("path", triggerName),
		"workflowName":      autorest.Encode("path", workflowName),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/triggers/{triggerName}/run", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RunSender sends the Run request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowTriggersClient) RunSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// RunResponder handles the response to the Run request. The method always
// closes the http.Response Body.
func (client WorkflowTriggersClient) RunResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// SetState sets the state of a workflow trigger.
// Parameters:
// resourceGroupName - the resource group name.
// workflowName - the workflow name.
// triggerName - the workflow trigger name.
// setState - the workflow trigger state.
func (client WorkflowTriggersClient) SetState(ctx context.Context, resourceGroupName string, workflowName string, triggerName string, setState SetTriggerStateActionDefinition) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkflowTriggersClient.SetState")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: setState,
			Constraints: []validation.Constraint{{Target: "setState.Source", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("logic.WorkflowTriggersClient", "SetState", err.Error())
	}

	req, err := client.SetStatePreparer(ctx, resourceGroupName, workflowName, triggerName, setState)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowTriggersClient", "SetState", nil, "Failure preparing request")
		return
	}

	resp, err := client.SetStateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "logic.WorkflowTriggersClient", "SetState", resp, "Failure sending request")
		return
	}

	result, err = client.SetStateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowTriggersClient", "SetState", resp, "Failure responding to request")
	}

	return
}

// SetStatePreparer prepares the SetState request.
func (client WorkflowTriggersClient) SetStatePreparer(ctx context.Context, resourceGroupName string, workflowName string, triggerName string, setState SetTriggerStateActionDefinition) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"triggerName":       autorest.Encode("path", triggerName),
		"workflowName":      autorest.Encode("path", workflowName),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/triggers/{triggerName}/setState", pathParameters),
		autorest.WithJSON(setState),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// SetStateSender sends the SetState request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowTriggersClient) SetStateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// SetStateResponder handles the response to the SetState request. The method always
// closes the http.Response Body.
func (client WorkflowTriggersClient) SetStateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
