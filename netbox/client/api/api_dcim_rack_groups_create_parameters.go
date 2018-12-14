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

// NewAPIDcimRackGroupsCreateParams creates a new APIDcimRackGroupsCreateParams object
// with the default values initialized.
func NewAPIDcimRackGroupsCreateParams() *APIDcimRackGroupsCreateParams {
	var ()
	return &APIDcimRackGroupsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIDcimRackGroupsCreateParamsWithTimeout creates a new APIDcimRackGroupsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIDcimRackGroupsCreateParamsWithTimeout(timeout time.Duration) *APIDcimRackGroupsCreateParams {
	var ()
	return &APIDcimRackGroupsCreateParams{

		timeout: timeout,
	}
}

// NewAPIDcimRackGroupsCreateParamsWithContext creates a new APIDcimRackGroupsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIDcimRackGroupsCreateParamsWithContext(ctx context.Context) *APIDcimRackGroupsCreateParams {
	var ()
	return &APIDcimRackGroupsCreateParams{

		Context: ctx,
	}
}

// NewAPIDcimRackGroupsCreateParamsWithHTTPClient creates a new APIDcimRackGroupsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIDcimRackGroupsCreateParamsWithHTTPClient(client *http.Client) *APIDcimRackGroupsCreateParams {
	var ()
	return &APIDcimRackGroupsCreateParams{
		HTTPClient: client,
	}
}

/*APIDcimRackGroupsCreateParams contains all the parameters to send to the API endpoint
for the api dcim rack groups create operation typically these are written to a http.Request
*/
type APIDcimRackGroupsCreateParams struct {

	/*Data*/
	Data *models.WritableRackGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the api dcim rack groups create params
func (o *APIDcimRackGroupsCreateParams) WithTimeout(timeout time.Duration) *APIDcimRackGroupsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the api dcim rack groups create params
func (o *APIDcimRackGroupsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the api dcim rack groups create params
func (o *APIDcimRackGroupsCreateParams) WithContext(ctx context.Context) *APIDcimRackGroupsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the api dcim rack groups create params
func (o *APIDcimRackGroupsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the api dcim rack groups create params
func (o *APIDcimRackGroupsCreateParams) WithHTTPClient(client *http.Client) *APIDcimRackGroupsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the api dcim rack groups create params
func (o *APIDcimRackGroupsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the api dcim rack groups create params
func (o *APIDcimRackGroupsCreateParams) WithData(data *models.WritableRackGroup) *APIDcimRackGroupsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the api dcim rack groups create params
func (o *APIDcimRackGroupsCreateParams) SetData(data *models.WritableRackGroup) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *APIDcimRackGroupsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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