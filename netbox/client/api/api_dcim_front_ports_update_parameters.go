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

	models "github.com/digitalocean/go-netbox/netbox/models"
)

// NewAPIDcimFrontPortsUpdateParams creates a new APIDcimFrontPortsUpdateParams object
// with the default values initialized.
func NewAPIDcimFrontPortsUpdateParams() *APIDcimFrontPortsUpdateParams {
	var ()
	return &APIDcimFrontPortsUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIDcimFrontPortsUpdateParamsWithTimeout creates a new APIDcimFrontPortsUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIDcimFrontPortsUpdateParamsWithTimeout(timeout time.Duration) *APIDcimFrontPortsUpdateParams {
	var ()
	return &APIDcimFrontPortsUpdateParams{

		timeout: timeout,
	}
}

// NewAPIDcimFrontPortsUpdateParamsWithContext creates a new APIDcimFrontPortsUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIDcimFrontPortsUpdateParamsWithContext(ctx context.Context) *APIDcimFrontPortsUpdateParams {
	var ()
	return &APIDcimFrontPortsUpdateParams{

		Context: ctx,
	}
}

// NewAPIDcimFrontPortsUpdateParamsWithHTTPClient creates a new APIDcimFrontPortsUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIDcimFrontPortsUpdateParamsWithHTTPClient(client *http.Client) *APIDcimFrontPortsUpdateParams {
	var ()
	return &APIDcimFrontPortsUpdateParams{
		HTTPClient: client,
	}
}

/*APIDcimFrontPortsUpdateParams contains all the parameters to send to the API endpoint
for the api dcim front ports update operation typically these are written to a http.Request
*/
type APIDcimFrontPortsUpdateParams struct {

	/*Data*/
	Data *models.WritableFrontPort
	/*ID
	  A unique integer value identifying this front port.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the api dcim front ports update params
func (o *APIDcimFrontPortsUpdateParams) WithTimeout(timeout time.Duration) *APIDcimFrontPortsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the api dcim front ports update params
func (o *APIDcimFrontPortsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the api dcim front ports update params
func (o *APIDcimFrontPortsUpdateParams) WithContext(ctx context.Context) *APIDcimFrontPortsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the api dcim front ports update params
func (o *APIDcimFrontPortsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the api dcim front ports update params
func (o *APIDcimFrontPortsUpdateParams) WithHTTPClient(client *http.Client) *APIDcimFrontPortsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the api dcim front ports update params
func (o *APIDcimFrontPortsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the api dcim front ports update params
func (o *APIDcimFrontPortsUpdateParams) WithData(data *models.WritableFrontPort) *APIDcimFrontPortsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the api dcim front ports update params
func (o *APIDcimFrontPortsUpdateParams) SetData(data *models.WritableFrontPort) {
	o.Data = data
}

// WithID adds the id to the api dcim front ports update params
func (o *APIDcimFrontPortsUpdateParams) WithID(id int64) *APIDcimFrontPortsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the api dcim front ports update params
func (o *APIDcimFrontPortsUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *APIDcimFrontPortsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
