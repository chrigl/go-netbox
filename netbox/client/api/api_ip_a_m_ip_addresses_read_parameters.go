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

// NewAPIIPAMIPAddressesReadParams creates a new APIIPAMIPAddressesReadParams object
// with the default values initialized.
func NewAPIIPAMIPAddressesReadParams() *APIIPAMIPAddressesReadParams {
	var ()
	return &APIIPAMIPAddressesReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIIPAMIPAddressesReadParamsWithTimeout creates a new APIIPAMIPAddressesReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIIPAMIPAddressesReadParamsWithTimeout(timeout time.Duration) *APIIPAMIPAddressesReadParams {
	var ()
	return &APIIPAMIPAddressesReadParams{

		timeout: timeout,
	}
}

// NewAPIIPAMIPAddressesReadParamsWithContext creates a new APIIPAMIPAddressesReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIIPAMIPAddressesReadParamsWithContext(ctx context.Context) *APIIPAMIPAddressesReadParams {
	var ()
	return &APIIPAMIPAddressesReadParams{

		Context: ctx,
	}
}

// NewAPIIPAMIPAddressesReadParamsWithHTTPClient creates a new APIIPAMIPAddressesReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIIPAMIPAddressesReadParamsWithHTTPClient(client *http.Client) *APIIPAMIPAddressesReadParams {
	var ()
	return &APIIPAMIPAddressesReadParams{
		HTTPClient: client,
	}
}

/*APIIPAMIPAddressesReadParams contains all the parameters to send to the API endpoint
for the api ipam ip addresses read operation typically these are written to a http.Request
*/
type APIIPAMIPAddressesReadParams struct {

	/*ID
	  A unique integer value identifying this IP address.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the api ipam ip addresses read params
func (o *APIIPAMIPAddressesReadParams) WithTimeout(timeout time.Duration) *APIIPAMIPAddressesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the api ipam ip addresses read params
func (o *APIIPAMIPAddressesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the api ipam ip addresses read params
func (o *APIIPAMIPAddressesReadParams) WithContext(ctx context.Context) *APIIPAMIPAddressesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the api ipam ip addresses read params
func (o *APIIPAMIPAddressesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the api ipam ip addresses read params
func (o *APIIPAMIPAddressesReadParams) WithHTTPClient(client *http.Client) *APIIPAMIPAddressesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the api ipam ip addresses read params
func (o *APIIPAMIPAddressesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the api ipam ip addresses read params
func (o *APIIPAMIPAddressesReadParams) WithID(id int64) *APIIPAMIPAddressesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the api ipam ip addresses read params
func (o *APIIPAMIPAddressesReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *APIIPAMIPAddressesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
