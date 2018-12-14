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

// NewAPIDcimRearPortTemplatesReadParams creates a new APIDcimRearPortTemplatesReadParams object
// with the default values initialized.
func NewAPIDcimRearPortTemplatesReadParams() *APIDcimRearPortTemplatesReadParams {
	var ()
	return &APIDcimRearPortTemplatesReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIDcimRearPortTemplatesReadParamsWithTimeout creates a new APIDcimRearPortTemplatesReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIDcimRearPortTemplatesReadParamsWithTimeout(timeout time.Duration) *APIDcimRearPortTemplatesReadParams {
	var ()
	return &APIDcimRearPortTemplatesReadParams{

		timeout: timeout,
	}
}

// NewAPIDcimRearPortTemplatesReadParamsWithContext creates a new APIDcimRearPortTemplatesReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIDcimRearPortTemplatesReadParamsWithContext(ctx context.Context) *APIDcimRearPortTemplatesReadParams {
	var ()
	return &APIDcimRearPortTemplatesReadParams{

		Context: ctx,
	}
}

// NewAPIDcimRearPortTemplatesReadParamsWithHTTPClient creates a new APIDcimRearPortTemplatesReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIDcimRearPortTemplatesReadParamsWithHTTPClient(client *http.Client) *APIDcimRearPortTemplatesReadParams {
	var ()
	return &APIDcimRearPortTemplatesReadParams{
		HTTPClient: client,
	}
}

/*APIDcimRearPortTemplatesReadParams contains all the parameters to send to the API endpoint
for the api dcim rear port templates read operation typically these are written to a http.Request
*/
type APIDcimRearPortTemplatesReadParams struct {

	/*ID
	  A unique integer value identifying this rear port template.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the api dcim rear port templates read params
func (o *APIDcimRearPortTemplatesReadParams) WithTimeout(timeout time.Duration) *APIDcimRearPortTemplatesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the api dcim rear port templates read params
func (o *APIDcimRearPortTemplatesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the api dcim rear port templates read params
func (o *APIDcimRearPortTemplatesReadParams) WithContext(ctx context.Context) *APIDcimRearPortTemplatesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the api dcim rear port templates read params
func (o *APIDcimRearPortTemplatesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the api dcim rear port templates read params
func (o *APIDcimRearPortTemplatesReadParams) WithHTTPClient(client *http.Client) *APIDcimRearPortTemplatesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the api dcim rear port templates read params
func (o *APIDcimRearPortTemplatesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the api dcim rear port templates read params
func (o *APIDcimRearPortTemplatesReadParams) WithID(id int64) *APIDcimRearPortTemplatesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the api dcim rear port templates read params
func (o *APIDcimRearPortTemplatesReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *APIDcimRearPortTemplatesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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