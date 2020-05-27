package qnamaker

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

// AlterationsClient is the an API for QnAMaker Service
type AlterationsClient struct {
	BaseClient
}

// NewAlterationsClient creates an instance of the AlterationsClient client.
func NewAlterationsClient(endpoint string) AlterationsClient {
	return AlterationsClient{New(endpoint)}
}

// Get sends the get request.
func (client AlterationsClient) Get(ctx context.Context) (result WordAlterationsDTO, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AlterationsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "qnamaker.AlterationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "qnamaker.AlterationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "qnamaker.AlterationsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client AlterationsClient) GetPreparer(ctx context.Context) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{Endpoint}/qnamaker/v4.0", urlParameters),
		autorest.WithPath("/alterations"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AlterationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AlterationsClient) GetResponder(resp *http.Response) (result WordAlterationsDTO, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Replace sends the replace request.
// Parameters:
// wordAlterations - new alterations data.
func (client AlterationsClient) Replace(ctx context.Context, wordAlterations WordAlterationsDTO) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AlterationsClient.Replace")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: wordAlterations,
			Constraints: []validation.Constraint{{Target: "wordAlterations.WordAlterations", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("qnamaker.AlterationsClient", "Replace", err.Error())
	}

	req, err := client.ReplacePreparer(ctx, wordAlterations)
	if err != nil {
		err = autorest.NewErrorWithError(err, "qnamaker.AlterationsClient", "Replace", nil, "Failure preparing request")
		return
	}

	resp, err := client.ReplaceSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "qnamaker.AlterationsClient", "Replace", resp, "Failure sending request")
		return
	}

	result, err = client.ReplaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "qnamaker.AlterationsClient", "Replace", resp, "Failure responding to request")
	}

	return
}

// ReplacePreparer prepares the Replace request.
func (client AlterationsClient) ReplacePreparer(ctx context.Context, wordAlterations WordAlterationsDTO) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithCustomBaseURL("{Endpoint}/qnamaker/v4.0", urlParameters),
		autorest.WithPath("/alterations"),
		autorest.WithJSON(wordAlterations))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ReplaceSender sends the Replace request. The method will close the
// http.Response Body if it receives an error.
func (client AlterationsClient) ReplaceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ReplaceResponder handles the response to the Replace request. The method always
// closes the http.Response Body.
func (client AlterationsClient) ReplaceResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}
