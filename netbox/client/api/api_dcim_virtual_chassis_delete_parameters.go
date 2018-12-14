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

// NewAPIDcimVirtualChassisDeleteParams creates a new APIDcimVirtualChassisDeleteParams object
// with the default values initialized.
func NewAPIDcimVirtualChassisDeleteParams() *APIDcimVirtualChassisDeleteParams {
	var ()
	return &APIDcimVirtualChassisDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIDcimVirtualChassisDeleteParamsWithTimeout creates a new APIDcimVirtualChassisDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIDcimVirtualChassisDeleteParamsWithTimeout(timeout time.Duration) *APIDcimVirtualChassisDeleteParams {
	var ()
	return &APIDcimVirtualChassisDeleteParams{

		timeout: timeout,
	}
}

// NewAPIDcimVirtualChassisDeleteParamsWithContext creates a new APIDcimVirtualChassisDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIDcimVirtualChassisDeleteParamsWithContext(ctx context.Context) *APIDcimVirtualChassisDeleteParams {
	var ()
	return &APIDcimVirtualChassisDeleteParams{

		Context: ctx,
	}
}

// NewAPIDcimVirtualChassisDeleteParamsWithHTTPClient creates a new APIDcimVirtualChassisDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIDcimVirtualChassisDeleteParamsWithHTTPClient(client *http.Client) *APIDcimVirtualChassisDeleteParams {
	var ()
	return &APIDcimVirtualChassisDeleteParams{
		HTTPClient: client,
	}
}

/*APIDcimVirtualChassisDeleteParams contains all the parameters to send to the API endpoint
for the api dcim virtual chassis delete operation typically these are written to a http.Request
*/
type APIDcimVirtualChassisDeleteParams struct {

	/*ID
	  A unique integer value identifying this virtual chassis.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the api dcim virtual chassis delete params
func (o *APIDcimVirtualChassisDeleteParams) WithTimeout(timeout time.Duration) *APIDcimVirtualChassisDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the api dcim virtual chassis delete params
func (o *APIDcimVirtualChassisDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the api dcim virtual chassis delete params
func (o *APIDcimVirtualChassisDeleteParams) WithContext(ctx context.Context) *APIDcimVirtualChassisDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the api dcim virtual chassis delete params
func (o *APIDcimVirtualChassisDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the api dcim virtual chassis delete params
func (o *APIDcimVirtualChassisDeleteParams) WithHTTPClient(client *http.Client) *APIDcimVirtualChassisDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the api dcim virtual chassis delete params
func (o *APIDcimVirtualChassisDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the api dcim virtual chassis delete params
func (o *APIDcimVirtualChassisDeleteParams) WithID(id int64) *APIDcimVirtualChassisDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the api dcim virtual chassis delete params
func (o *APIDcimVirtualChassisDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *APIDcimVirtualChassisDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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