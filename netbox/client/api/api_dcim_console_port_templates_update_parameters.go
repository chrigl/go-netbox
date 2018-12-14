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

// NewAPIDcimConsolePortTemplatesUpdateParams creates a new APIDcimConsolePortTemplatesUpdateParams object
// with the default values initialized.
func NewAPIDcimConsolePortTemplatesUpdateParams() *APIDcimConsolePortTemplatesUpdateParams {
	var ()
	return &APIDcimConsolePortTemplatesUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIDcimConsolePortTemplatesUpdateParamsWithTimeout creates a new APIDcimConsolePortTemplatesUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIDcimConsolePortTemplatesUpdateParamsWithTimeout(timeout time.Duration) *APIDcimConsolePortTemplatesUpdateParams {
	var ()
	return &APIDcimConsolePortTemplatesUpdateParams{

		timeout: timeout,
	}
}

// NewAPIDcimConsolePortTemplatesUpdateParamsWithContext creates a new APIDcimConsolePortTemplatesUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIDcimConsolePortTemplatesUpdateParamsWithContext(ctx context.Context) *APIDcimConsolePortTemplatesUpdateParams {
	var ()
	return &APIDcimConsolePortTemplatesUpdateParams{

		Context: ctx,
	}
}

// NewAPIDcimConsolePortTemplatesUpdateParamsWithHTTPClient creates a new APIDcimConsolePortTemplatesUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIDcimConsolePortTemplatesUpdateParamsWithHTTPClient(client *http.Client) *APIDcimConsolePortTemplatesUpdateParams {
	var ()
	return &APIDcimConsolePortTemplatesUpdateParams{
		HTTPClient: client,
	}
}

/*APIDcimConsolePortTemplatesUpdateParams contains all the parameters to send to the API endpoint
for the api dcim console port templates update operation typically these are written to a http.Request
*/
type APIDcimConsolePortTemplatesUpdateParams struct {

	/*Data*/
	Data *models.WritableConsolePortTemplate
	/*ID
	  A unique integer value identifying this console port template.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the api dcim console port templates update params
func (o *APIDcimConsolePortTemplatesUpdateParams) WithTimeout(timeout time.Duration) *APIDcimConsolePortTemplatesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the api dcim console port templates update params
func (o *APIDcimConsolePortTemplatesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the api dcim console port templates update params
func (o *APIDcimConsolePortTemplatesUpdateParams) WithContext(ctx context.Context) *APIDcimConsolePortTemplatesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the api dcim console port templates update params
func (o *APIDcimConsolePortTemplatesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the api dcim console port templates update params
func (o *APIDcimConsolePortTemplatesUpdateParams) WithHTTPClient(client *http.Client) *APIDcimConsolePortTemplatesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the api dcim console port templates update params
func (o *APIDcimConsolePortTemplatesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the api dcim console port templates update params
func (o *APIDcimConsolePortTemplatesUpdateParams) WithData(data *models.WritableConsolePortTemplate) *APIDcimConsolePortTemplatesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the api dcim console port templates update params
func (o *APIDcimConsolePortTemplatesUpdateParams) SetData(data *models.WritableConsolePortTemplate) {
	o.Data = data
}

// WithID adds the id to the api dcim console port templates update params
func (o *APIDcimConsolePortTemplatesUpdateParams) WithID(id int64) *APIDcimConsolePortTemplatesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the api dcim console port templates update params
func (o *APIDcimConsolePortTemplatesUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *APIDcimConsolePortTemplatesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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