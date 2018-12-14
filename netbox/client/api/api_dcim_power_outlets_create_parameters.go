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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/digitalocean/go-netbox/netbox/models"
)

// NewAPIDcimPowerOutletsCreateParams creates a new APIDcimPowerOutletsCreateParams object
// with the default values initialized.
func NewAPIDcimPowerOutletsCreateParams() *APIDcimPowerOutletsCreateParams {
	var ()
	return &APIDcimPowerOutletsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIDcimPowerOutletsCreateParamsWithTimeout creates a new APIDcimPowerOutletsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIDcimPowerOutletsCreateParamsWithTimeout(timeout time.Duration) *APIDcimPowerOutletsCreateParams {
	var ()
	return &APIDcimPowerOutletsCreateParams{

		timeout: timeout,
	}
}

// NewAPIDcimPowerOutletsCreateParamsWithContext creates a new APIDcimPowerOutletsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIDcimPowerOutletsCreateParamsWithContext(ctx context.Context) *APIDcimPowerOutletsCreateParams {
	var ()
	return &APIDcimPowerOutletsCreateParams{

		Context: ctx,
	}
}

// NewAPIDcimPowerOutletsCreateParamsWithHTTPClient creates a new APIDcimPowerOutletsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIDcimPowerOutletsCreateParamsWithHTTPClient(client *http.Client) *APIDcimPowerOutletsCreateParams {
	var ()
	return &APIDcimPowerOutletsCreateParams{
		HTTPClient: client,
	}
}

/*APIDcimPowerOutletsCreateParams contains all the parameters to send to the API endpoint
for the api dcim power outlets create operation typically these are written to a http.Request
*/
type APIDcimPowerOutletsCreateParams struct {

	/*Data*/
	Data *models.WritablePowerOutlet

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the api dcim power outlets create params
func (o *APIDcimPowerOutletsCreateParams) WithTimeout(timeout time.Duration) *APIDcimPowerOutletsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the api dcim power outlets create params
func (o *APIDcimPowerOutletsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the api dcim power outlets create params
func (o *APIDcimPowerOutletsCreateParams) WithContext(ctx context.Context) *APIDcimPowerOutletsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the api dcim power outlets create params
func (o *APIDcimPowerOutletsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the api dcim power outlets create params
func (o *APIDcimPowerOutletsCreateParams) WithHTTPClient(client *http.Client) *APIDcimPowerOutletsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the api dcim power outlets create params
func (o *APIDcimPowerOutletsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the api dcim power outlets create params
func (o *APIDcimPowerOutletsCreateParams) WithData(data *models.WritablePowerOutlet) *APIDcimPowerOutletsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the api dcim power outlets create params
func (o *APIDcimPowerOutletsCreateParams) SetData(data *models.WritablePowerOutlet) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *APIDcimPowerOutletsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
