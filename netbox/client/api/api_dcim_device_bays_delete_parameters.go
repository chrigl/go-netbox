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

// NewAPIDcimDeviceBaysDeleteParams creates a new APIDcimDeviceBaysDeleteParams object
// with the default values initialized.
func NewAPIDcimDeviceBaysDeleteParams() *APIDcimDeviceBaysDeleteParams {
	var ()
	return &APIDcimDeviceBaysDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIDcimDeviceBaysDeleteParamsWithTimeout creates a new APIDcimDeviceBaysDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIDcimDeviceBaysDeleteParamsWithTimeout(timeout time.Duration) *APIDcimDeviceBaysDeleteParams {
	var ()
	return &APIDcimDeviceBaysDeleteParams{

		timeout: timeout,
	}
}

// NewAPIDcimDeviceBaysDeleteParamsWithContext creates a new APIDcimDeviceBaysDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIDcimDeviceBaysDeleteParamsWithContext(ctx context.Context) *APIDcimDeviceBaysDeleteParams {
	var ()
	return &APIDcimDeviceBaysDeleteParams{

		Context: ctx,
	}
}

// NewAPIDcimDeviceBaysDeleteParamsWithHTTPClient creates a new APIDcimDeviceBaysDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIDcimDeviceBaysDeleteParamsWithHTTPClient(client *http.Client) *APIDcimDeviceBaysDeleteParams {
	var ()
	return &APIDcimDeviceBaysDeleteParams{
		HTTPClient: client,
	}
}

/*APIDcimDeviceBaysDeleteParams contains all the parameters to send to the API endpoint
for the api dcim device bays delete operation typically these are written to a http.Request
*/
type APIDcimDeviceBaysDeleteParams struct {

	/*ID
	  A unique integer value identifying this device bay.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the api dcim device bays delete params
func (o *APIDcimDeviceBaysDeleteParams) WithTimeout(timeout time.Duration) *APIDcimDeviceBaysDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the api dcim device bays delete params
func (o *APIDcimDeviceBaysDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the api dcim device bays delete params
func (o *APIDcimDeviceBaysDeleteParams) WithContext(ctx context.Context) *APIDcimDeviceBaysDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the api dcim device bays delete params
func (o *APIDcimDeviceBaysDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the api dcim device bays delete params
func (o *APIDcimDeviceBaysDeleteParams) WithHTTPClient(client *http.Client) *APIDcimDeviceBaysDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the api dcim device bays delete params
func (o *APIDcimDeviceBaysDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the api dcim device bays delete params
func (o *APIDcimDeviceBaysDeleteParams) WithID(id int64) *APIDcimDeviceBaysDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the api dcim device bays delete params
func (o *APIDcimDeviceBaysDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *APIDcimDeviceBaysDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
