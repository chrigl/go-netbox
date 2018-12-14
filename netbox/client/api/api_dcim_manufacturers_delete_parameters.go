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

// NewAPIDcimManufacturersDeleteParams creates a new APIDcimManufacturersDeleteParams object
// with the default values initialized.
func NewAPIDcimManufacturersDeleteParams() *APIDcimManufacturersDeleteParams {
	var ()
	return &APIDcimManufacturersDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIDcimManufacturersDeleteParamsWithTimeout creates a new APIDcimManufacturersDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIDcimManufacturersDeleteParamsWithTimeout(timeout time.Duration) *APIDcimManufacturersDeleteParams {
	var ()
	return &APIDcimManufacturersDeleteParams{

		timeout: timeout,
	}
}

// NewAPIDcimManufacturersDeleteParamsWithContext creates a new APIDcimManufacturersDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIDcimManufacturersDeleteParamsWithContext(ctx context.Context) *APIDcimManufacturersDeleteParams {
	var ()
	return &APIDcimManufacturersDeleteParams{

		Context: ctx,
	}
}

// NewAPIDcimManufacturersDeleteParamsWithHTTPClient creates a new APIDcimManufacturersDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIDcimManufacturersDeleteParamsWithHTTPClient(client *http.Client) *APIDcimManufacturersDeleteParams {
	var ()
	return &APIDcimManufacturersDeleteParams{
		HTTPClient: client,
	}
}

/*APIDcimManufacturersDeleteParams contains all the parameters to send to the API endpoint
for the api dcim manufacturers delete operation typically these are written to a http.Request
*/
type APIDcimManufacturersDeleteParams struct {

	/*ID
	  A unique integer value identifying this manufacturer.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the api dcim manufacturers delete params
func (o *APIDcimManufacturersDeleteParams) WithTimeout(timeout time.Duration) *APIDcimManufacturersDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the api dcim manufacturers delete params
func (o *APIDcimManufacturersDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the api dcim manufacturers delete params
func (o *APIDcimManufacturersDeleteParams) WithContext(ctx context.Context) *APIDcimManufacturersDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the api dcim manufacturers delete params
func (o *APIDcimManufacturersDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the api dcim manufacturers delete params
func (o *APIDcimManufacturersDeleteParams) WithHTTPClient(client *http.Client) *APIDcimManufacturersDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the api dcim manufacturers delete params
func (o *APIDcimManufacturersDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the api dcim manufacturers delete params
func (o *APIDcimManufacturersDeleteParams) WithID(id int64) *APIDcimManufacturersDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the api dcim manufacturers delete params
func (o *APIDcimManufacturersDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *APIDcimManufacturersDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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