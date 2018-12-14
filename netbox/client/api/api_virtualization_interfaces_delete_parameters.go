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

// NewAPIVirtualizationInterfacesDeleteParams creates a new APIVirtualizationInterfacesDeleteParams object
// with the default values initialized.
func NewAPIVirtualizationInterfacesDeleteParams() *APIVirtualizationInterfacesDeleteParams {
	var ()
	return &APIVirtualizationInterfacesDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIVirtualizationInterfacesDeleteParamsWithTimeout creates a new APIVirtualizationInterfacesDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIVirtualizationInterfacesDeleteParamsWithTimeout(timeout time.Duration) *APIVirtualizationInterfacesDeleteParams {
	var ()
	return &APIVirtualizationInterfacesDeleteParams{

		timeout: timeout,
	}
}

// NewAPIVirtualizationInterfacesDeleteParamsWithContext creates a new APIVirtualizationInterfacesDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIVirtualizationInterfacesDeleteParamsWithContext(ctx context.Context) *APIVirtualizationInterfacesDeleteParams {
	var ()
	return &APIVirtualizationInterfacesDeleteParams{

		Context: ctx,
	}
}

// NewAPIVirtualizationInterfacesDeleteParamsWithHTTPClient creates a new APIVirtualizationInterfacesDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIVirtualizationInterfacesDeleteParamsWithHTTPClient(client *http.Client) *APIVirtualizationInterfacesDeleteParams {
	var ()
	return &APIVirtualizationInterfacesDeleteParams{
		HTTPClient: client,
	}
}

/*APIVirtualizationInterfacesDeleteParams contains all the parameters to send to the API endpoint
for the api virtualization interfaces delete operation typically these are written to a http.Request
*/
type APIVirtualizationInterfacesDeleteParams struct {

	/*ID
	  A unique integer value identifying this interface.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the api virtualization interfaces delete params
func (o *APIVirtualizationInterfacesDeleteParams) WithTimeout(timeout time.Duration) *APIVirtualizationInterfacesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the api virtualization interfaces delete params
func (o *APIVirtualizationInterfacesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the api virtualization interfaces delete params
func (o *APIVirtualizationInterfacesDeleteParams) WithContext(ctx context.Context) *APIVirtualizationInterfacesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the api virtualization interfaces delete params
func (o *APIVirtualizationInterfacesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the api virtualization interfaces delete params
func (o *APIVirtualizationInterfacesDeleteParams) WithHTTPClient(client *http.Client) *APIVirtualizationInterfacesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the api virtualization interfaces delete params
func (o *APIVirtualizationInterfacesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the api virtualization interfaces delete params
func (o *APIVirtualizationInterfacesDeleteParams) WithID(id int64) *APIVirtualizationInterfacesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the api virtualization interfaces delete params
func (o *APIVirtualizationInterfacesDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *APIVirtualizationInterfacesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
