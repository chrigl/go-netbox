// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewAPIIPAMAggregatesDeleteParams creates a new APIIPAMAggregatesDeleteParams object
// with the default values initialized.
func NewAPIIPAMAggregatesDeleteParams() *APIIPAMAggregatesDeleteParams {
	var ()
	return &APIIPAMAggregatesDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIIPAMAggregatesDeleteParamsWithTimeout creates a new APIIPAMAggregatesDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIIPAMAggregatesDeleteParamsWithTimeout(timeout time.Duration) *APIIPAMAggregatesDeleteParams {
	var ()
	return &APIIPAMAggregatesDeleteParams{

		timeout: timeout,
	}
}

// NewAPIIPAMAggregatesDeleteParamsWithContext creates a new APIIPAMAggregatesDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIIPAMAggregatesDeleteParamsWithContext(ctx context.Context) *APIIPAMAggregatesDeleteParams {
	var ()
	return &APIIPAMAggregatesDeleteParams{

		Context: ctx,
	}
}

// NewAPIIPAMAggregatesDeleteParamsWithHTTPClient creates a new APIIPAMAggregatesDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIIPAMAggregatesDeleteParamsWithHTTPClient(client *http.Client) *APIIPAMAggregatesDeleteParams {
	var ()
	return &APIIPAMAggregatesDeleteParams{
		HTTPClient: client,
	}
}

/*APIIPAMAggregatesDeleteParams contains all the parameters to send to the API endpoint
for the api ipam aggregates delete operation typically these are written to a http.Request
*/
type APIIPAMAggregatesDeleteParams struct {

	/*ID
	  A unique integer value identifying this aggregate.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the api ipam aggregates delete params
func (o *APIIPAMAggregatesDeleteParams) WithTimeout(timeout time.Duration) *APIIPAMAggregatesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the api ipam aggregates delete params
func (o *APIIPAMAggregatesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the api ipam aggregates delete params
func (o *APIIPAMAggregatesDeleteParams) WithContext(ctx context.Context) *APIIPAMAggregatesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the api ipam aggregates delete params
func (o *APIIPAMAggregatesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the api ipam aggregates delete params
func (o *APIIPAMAggregatesDeleteParams) WithHTTPClient(client *http.Client) *APIIPAMAggregatesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the api ipam aggregates delete params
func (o *APIIPAMAggregatesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the api ipam aggregates delete params
func (o *APIIPAMAggregatesDeleteParams) WithID(id int64) *APIIPAMAggregatesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the api ipam aggregates delete params
func (o *APIIPAMAggregatesDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *APIIPAMAggregatesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}