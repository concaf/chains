// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright 2021 The Sigstore Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package timestamp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/sigstore/rekor/pkg/generated/models"
)

// NewTimestampResponseParams creates a new TimestampResponseParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTimestampResponseParams() *TimestampResponseParams {
	return &TimestampResponseParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTimestampResponseParamsWithTimeout creates a new TimestampResponseParams object
// with the ability to set a timeout on a request.
func NewTimestampResponseParamsWithTimeout(timeout time.Duration) *TimestampResponseParams {
	return &TimestampResponseParams{
		timeout: timeout,
	}
}

// NewTimestampResponseParamsWithContext creates a new TimestampResponseParams object
// with the ability to set a context for a request.
func NewTimestampResponseParamsWithContext(ctx context.Context) *TimestampResponseParams {
	return &TimestampResponseParams{
		Context: ctx,
	}
}

// NewTimestampResponseParamsWithHTTPClient creates a new TimestampResponseParams object
// with the ability to set a custom HTTPClient for a request.
func NewTimestampResponseParamsWithHTTPClient(client *http.Client) *TimestampResponseParams {
	return &TimestampResponseParams{
		HTTPClient: client,
	}
}

/* TimestampResponseParams contains all the parameters to send to the API endpoint
   for the timestamp response operation.

   Typically these are written to a http.Request.
*/
type TimestampResponseParams struct {

	// Query.
	Query *models.TimestampRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the timestamp response params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TimestampResponseParams) WithDefaults() *TimestampResponseParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the timestamp response params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TimestampResponseParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the timestamp response params
func (o *TimestampResponseParams) WithTimeout(timeout time.Duration) *TimestampResponseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the timestamp response params
func (o *TimestampResponseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the timestamp response params
func (o *TimestampResponseParams) WithContext(ctx context.Context) *TimestampResponseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the timestamp response params
func (o *TimestampResponseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the timestamp response params
func (o *TimestampResponseParams) WithHTTPClient(client *http.Client) *TimestampResponseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the timestamp response params
func (o *TimestampResponseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQuery adds the query to the timestamp response params
func (o *TimestampResponseParams) WithQuery(query *models.TimestampRequest) *TimestampResponseParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the timestamp response params
func (o *TimestampResponseParams) SetQuery(query *models.TimestampRequest) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *TimestampResponseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Query != nil {
		if err := r.SetBodyParam(o.Query); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}