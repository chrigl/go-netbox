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

// NewAPICircuitsCircuitTerminationsUpdateParams creates a new APICircuitsCircuitTerminationsUpdateParams object
// with the default values initialized.
func NewAPICircuitsCircuitTerminationsUpdateParams() *APICircuitsCircuitTerminationsUpdateParams {
	var ()
	return &APICircuitsCircuitTerminationsUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPICircuitsCircuitTerminationsUpdateParamsWithTimeout creates a new APICircuitsCircuitTerminationsUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPICircuitsCircuitTerminationsUpdateParamsWithTimeout(timeout time.Duration) *APICircuitsCircuitTerminationsUpdateParams {
	var ()
	return &APICircuitsCircuitTerminationsUpdateParams{

		timeout: timeout,
	}
}

// NewAPICircuitsCircuitTerminationsUpdateParamsWithContext creates a new APICircuitsCircuitTerminationsUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPICircuitsCircuitTerminationsUpdateParamsWithContext(ctx context.Context) *APICircuitsCircuitTerminationsUpdateParams {
	var ()
	return &APICircuitsCircuitTerminationsUpdateParams{

		Context: ctx,
	}
}

// NewAPICircuitsCircuitTerminationsUpdateParamsWithHTTPClient creates a new APICircuitsCircuitTerminationsUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPICircuitsCircuitTerminationsUpdateParamsWithHTTPClient(client *http.Client) *APICircuitsCircuitTerminationsUpdateParams {
	var ()
	return &APICircuitsCircuitTerminationsUpdateParams{
		HTTPClient: client,
	}
}

/*APICircuitsCircuitTerminationsUpdateParams contains all the parameters to send to the API endpoint
for the api circuits circuit terminations update operation typically these are written to a http.Request
*/
type APICircuitsCircuitTerminationsUpdateParams struct {

	/*Data*/
	Data *models.WritableCircuitTermination
	/*ID
	  A unique integer value identifying this circuit termination.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the api circuits circuit terminations update params
func (o *APICircuitsCircuitTerminationsUpdateParams) WithTimeout(timeout time.Duration) *APICircuitsCircuitTerminationsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the api circuits circuit terminations update params
func (o *APICircuitsCircuitTerminationsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the api circuits circuit terminations update params
func (o *APICircuitsCircuitTerminationsUpdateParams) WithContext(ctx context.Context) *APICircuitsCircuitTerminationsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the api circuits circuit terminations update params
func (o *APICircuitsCircuitTerminationsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the api circuits circuit terminations update params
func (o *APICircuitsCircuitTerminationsUpdateParams) WithHTTPClient(client *http.Client) *APICircuitsCircuitTerminationsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the api circuits circuit terminations update params
func (o *APICircuitsCircuitTerminationsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the api circuits circuit terminations update params
func (o *APICircuitsCircuitTerminationsUpdateParams) WithData(data *models.WritableCircuitTermination) *APICircuitsCircuitTerminationsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the api circuits circuit terminations update params
func (o *APICircuitsCircuitTerminationsUpdateParams) SetData(data *models.WritableCircuitTermination) {
	o.Data = data
}

// WithID adds the id to the api circuits circuit terminations update params
func (o *APICircuitsCircuitTerminationsUpdateParams) WithID(id int64) *APICircuitsCircuitTerminationsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the api circuits circuit terminations update params
func (o *APICircuitsCircuitTerminationsUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *APICircuitsCircuitTerminationsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
