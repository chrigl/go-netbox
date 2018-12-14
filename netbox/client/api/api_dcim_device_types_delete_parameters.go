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

// NewAPIDcimDeviceTypesDeleteParams creates a new APIDcimDeviceTypesDeleteParams object
// with the default values initialized.
func NewAPIDcimDeviceTypesDeleteParams() *APIDcimDeviceTypesDeleteParams {
	var ()
	return &APIDcimDeviceTypesDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIDcimDeviceTypesDeleteParamsWithTimeout creates a new APIDcimDeviceTypesDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIDcimDeviceTypesDeleteParamsWithTimeout(timeout time.Duration) *APIDcimDeviceTypesDeleteParams {
	var ()
	return &APIDcimDeviceTypesDeleteParams{

		timeout: timeout,
	}
}

// NewAPIDcimDeviceTypesDeleteParamsWithContext creates a new APIDcimDeviceTypesDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIDcimDeviceTypesDeleteParamsWithContext(ctx context.Context) *APIDcimDeviceTypesDeleteParams {
	var ()
	return &APIDcimDeviceTypesDeleteParams{

		Context: ctx,
	}
}

// NewAPIDcimDeviceTypesDeleteParamsWithHTTPClient creates a new APIDcimDeviceTypesDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIDcimDeviceTypesDeleteParamsWithHTTPClient(client *http.Client) *APIDcimDeviceTypesDeleteParams {
	var ()
	return &APIDcimDeviceTypesDeleteParams{
		HTTPClient: client,
	}
}

/*APIDcimDeviceTypesDeleteParams contains all the parameters to send to the API endpoint
for the api dcim device types delete operation typically these are written to a http.Request
*/
type APIDcimDeviceTypesDeleteParams struct {

	/*ID
	  A unique integer value identifying this device type.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the api dcim device types delete params
func (o *APIDcimDeviceTypesDeleteParams) WithTimeout(timeout time.Duration) *APIDcimDeviceTypesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the api dcim device types delete params
func (o *APIDcimDeviceTypesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the api dcim device types delete params
func (o *APIDcimDeviceTypesDeleteParams) WithContext(ctx context.Context) *APIDcimDeviceTypesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the api dcim device types delete params
func (o *APIDcimDeviceTypesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the api dcim device types delete params
func (o *APIDcimDeviceTypesDeleteParams) WithHTTPClient(client *http.Client) *APIDcimDeviceTypesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the api dcim device types delete params
func (o *APIDcimDeviceTypesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the api dcim device types delete params
func (o *APIDcimDeviceTypesDeleteParams) WithID(id int64) *APIDcimDeviceTypesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the api dcim device types delete params
func (o *APIDcimDeviceTypesDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *APIDcimDeviceTypesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
