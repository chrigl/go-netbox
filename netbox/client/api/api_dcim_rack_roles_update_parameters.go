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

// NewAPIDcimRackRolesUpdateParams creates a new APIDcimRackRolesUpdateParams object
// with the default values initialized.
func NewAPIDcimRackRolesUpdateParams() *APIDcimRackRolesUpdateParams {
	var ()
	return &APIDcimRackRolesUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIDcimRackRolesUpdateParamsWithTimeout creates a new APIDcimRackRolesUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIDcimRackRolesUpdateParamsWithTimeout(timeout time.Duration) *APIDcimRackRolesUpdateParams {
	var ()
	return &APIDcimRackRolesUpdateParams{

		timeout: timeout,
	}
}

// NewAPIDcimRackRolesUpdateParamsWithContext creates a new APIDcimRackRolesUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIDcimRackRolesUpdateParamsWithContext(ctx context.Context) *APIDcimRackRolesUpdateParams {
	var ()
	return &APIDcimRackRolesUpdateParams{

		Context: ctx,
	}
}

// NewAPIDcimRackRolesUpdateParamsWithHTTPClient creates a new APIDcimRackRolesUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIDcimRackRolesUpdateParamsWithHTTPClient(client *http.Client) *APIDcimRackRolesUpdateParams {
	var ()
	return &APIDcimRackRolesUpdateParams{
		HTTPClient: client,
	}
}

/*APIDcimRackRolesUpdateParams contains all the parameters to send to the API endpoint
for the api dcim rack roles update operation typically these are written to a http.Request
*/
type APIDcimRackRolesUpdateParams struct {

	/*Data*/
	Data *models.RackRole
	/*ID
	  A unique integer value identifying this rack role.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the api dcim rack roles update params
func (o *APIDcimRackRolesUpdateParams) WithTimeout(timeout time.Duration) *APIDcimRackRolesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the api dcim rack roles update params
func (o *APIDcimRackRolesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the api dcim rack roles update params
func (o *APIDcimRackRolesUpdateParams) WithContext(ctx context.Context) *APIDcimRackRolesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the api dcim rack roles update params
func (o *APIDcimRackRolesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the api dcim rack roles update params
func (o *APIDcimRackRolesUpdateParams) WithHTTPClient(client *http.Client) *APIDcimRackRolesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the api dcim rack roles update params
func (o *APIDcimRackRolesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the api dcim rack roles update params
func (o *APIDcimRackRolesUpdateParams) WithData(data *models.RackRole) *APIDcimRackRolesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the api dcim rack roles update params
func (o *APIDcimRackRolesUpdateParams) SetData(data *models.RackRole) {
	o.Data = data
}

// WithID adds the id to the api dcim rack roles update params
func (o *APIDcimRackRolesUpdateParams) WithID(id int64) *APIDcimRackRolesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the api dcim rack roles update params
func (o *APIDcimRackRolesUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *APIDcimRackRolesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
