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

// NewAPIDcimPowerOutletsPartialUpdateParams creates a new APIDcimPowerOutletsPartialUpdateParams object
// with the default values initialized.
func NewAPIDcimPowerOutletsPartialUpdateParams() *APIDcimPowerOutletsPartialUpdateParams {
	var ()
	return &APIDcimPowerOutletsPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIDcimPowerOutletsPartialUpdateParamsWithTimeout creates a new APIDcimPowerOutletsPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIDcimPowerOutletsPartialUpdateParamsWithTimeout(timeout time.Duration) *APIDcimPowerOutletsPartialUpdateParams {
	var ()
	return &APIDcimPowerOutletsPartialUpdateParams{

		timeout: timeout,
	}
}

// NewAPIDcimPowerOutletsPartialUpdateParamsWithContext creates a new APIDcimPowerOutletsPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIDcimPowerOutletsPartialUpdateParamsWithContext(ctx context.Context) *APIDcimPowerOutletsPartialUpdateParams {
	var ()
	return &APIDcimPowerOutletsPartialUpdateParams{

		Context: ctx,
	}
}

// NewAPIDcimPowerOutletsPartialUpdateParamsWithHTTPClient creates a new APIDcimPowerOutletsPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIDcimPowerOutletsPartialUpdateParamsWithHTTPClient(client *http.Client) *APIDcimPowerOutletsPartialUpdateParams {
	var ()
	return &APIDcimPowerOutletsPartialUpdateParams{
		HTTPClient: client,
	}
}

/*APIDcimPowerOutletsPartialUpdateParams contains all the parameters to send to the API endpoint
for the api dcim power outlets partial update operation typically these are written to a http.Request
*/
type APIDcimPowerOutletsPartialUpdateParams struct {

	/*Data*/
	Data *models.WritablePowerOutlet
	/*ID
	  A unique integer value identifying this power outlet.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the api dcim power outlets partial update params
func (o *APIDcimPowerOutletsPartialUpdateParams) WithTimeout(timeout time.Duration) *APIDcimPowerOutletsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the api dcim power outlets partial update params
func (o *APIDcimPowerOutletsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the api dcim power outlets partial update params
func (o *APIDcimPowerOutletsPartialUpdateParams) WithContext(ctx context.Context) *APIDcimPowerOutletsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the api dcim power outlets partial update params
func (o *APIDcimPowerOutletsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the api dcim power outlets partial update params
func (o *APIDcimPowerOutletsPartialUpdateParams) WithHTTPClient(client *http.Client) *APIDcimPowerOutletsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the api dcim power outlets partial update params
func (o *APIDcimPowerOutletsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the api dcim power outlets partial update params
func (o *APIDcimPowerOutletsPartialUpdateParams) WithData(data *models.WritablePowerOutlet) *APIDcimPowerOutletsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the api dcim power outlets partial update params
func (o *APIDcimPowerOutletsPartialUpdateParams) SetData(data *models.WritablePowerOutlet) {
	o.Data = data
}

// WithID adds the id to the api dcim power outlets partial update params
func (o *APIDcimPowerOutletsPartialUpdateParams) WithID(id int64) *APIDcimPowerOutletsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the api dcim power outlets partial update params
func (o *APIDcimPowerOutletsPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *APIDcimPowerOutletsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
