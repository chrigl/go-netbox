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

// NewAPIDcimRearPortsPartialUpdateParams creates a new APIDcimRearPortsPartialUpdateParams object
// with the default values initialized.
func NewAPIDcimRearPortsPartialUpdateParams() *APIDcimRearPortsPartialUpdateParams {
	var ()
	return &APIDcimRearPortsPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIDcimRearPortsPartialUpdateParamsWithTimeout creates a new APIDcimRearPortsPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIDcimRearPortsPartialUpdateParamsWithTimeout(timeout time.Duration) *APIDcimRearPortsPartialUpdateParams {
	var ()
	return &APIDcimRearPortsPartialUpdateParams{

		timeout: timeout,
	}
}

// NewAPIDcimRearPortsPartialUpdateParamsWithContext creates a new APIDcimRearPortsPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIDcimRearPortsPartialUpdateParamsWithContext(ctx context.Context) *APIDcimRearPortsPartialUpdateParams {
	var ()
	return &APIDcimRearPortsPartialUpdateParams{

		Context: ctx,
	}
}

// NewAPIDcimRearPortsPartialUpdateParamsWithHTTPClient creates a new APIDcimRearPortsPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIDcimRearPortsPartialUpdateParamsWithHTTPClient(client *http.Client) *APIDcimRearPortsPartialUpdateParams {
	var ()
	return &APIDcimRearPortsPartialUpdateParams{
		HTTPClient: client,
	}
}

/*APIDcimRearPortsPartialUpdateParams contains all the parameters to send to the API endpoint
for the api dcim rear ports partial update operation typically these are written to a http.Request
*/
type APIDcimRearPortsPartialUpdateParams struct {

	/*Data*/
	Data *models.WritableRearPort
	/*ID
	  A unique integer value identifying this rear port.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the api dcim rear ports partial update params
func (o *APIDcimRearPortsPartialUpdateParams) WithTimeout(timeout time.Duration) *APIDcimRearPortsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the api dcim rear ports partial update params
func (o *APIDcimRearPortsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the api dcim rear ports partial update params
func (o *APIDcimRearPortsPartialUpdateParams) WithContext(ctx context.Context) *APIDcimRearPortsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the api dcim rear ports partial update params
func (o *APIDcimRearPortsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the api dcim rear ports partial update params
func (o *APIDcimRearPortsPartialUpdateParams) WithHTTPClient(client *http.Client) *APIDcimRearPortsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the api dcim rear ports partial update params
func (o *APIDcimRearPortsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the api dcim rear ports partial update params
func (o *APIDcimRearPortsPartialUpdateParams) WithData(data *models.WritableRearPort) *APIDcimRearPortsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the api dcim rear ports partial update params
func (o *APIDcimRearPortsPartialUpdateParams) SetData(data *models.WritableRearPort) {
	o.Data = data
}

// WithID adds the id to the api dcim rear ports partial update params
func (o *APIDcimRearPortsPartialUpdateParams) WithID(id int64) *APIDcimRearPortsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the api dcim rear ports partial update params
func (o *APIDcimRearPortsPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *APIDcimRearPortsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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