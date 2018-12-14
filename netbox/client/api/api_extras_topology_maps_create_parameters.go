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

// NewAPIExtrasTopologyMapsCreateParams creates a new APIExtrasTopologyMapsCreateParams object
// with the default values initialized.
func NewAPIExtrasTopologyMapsCreateParams() *APIExtrasTopologyMapsCreateParams {
	var ()
	return &APIExtrasTopologyMapsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIExtrasTopologyMapsCreateParamsWithTimeout creates a new APIExtrasTopologyMapsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIExtrasTopologyMapsCreateParamsWithTimeout(timeout time.Duration) *APIExtrasTopologyMapsCreateParams {
	var ()
	return &APIExtrasTopologyMapsCreateParams{

		timeout: timeout,
	}
}

// NewAPIExtrasTopologyMapsCreateParamsWithContext creates a new APIExtrasTopologyMapsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIExtrasTopologyMapsCreateParamsWithContext(ctx context.Context) *APIExtrasTopologyMapsCreateParams {
	var ()
	return &APIExtrasTopologyMapsCreateParams{

		Context: ctx,
	}
}

// NewAPIExtrasTopologyMapsCreateParamsWithHTTPClient creates a new APIExtrasTopologyMapsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIExtrasTopologyMapsCreateParamsWithHTTPClient(client *http.Client) *APIExtrasTopologyMapsCreateParams {
	var ()
	return &APIExtrasTopologyMapsCreateParams{
		HTTPClient: client,
	}
}

/*APIExtrasTopologyMapsCreateParams contains all the parameters to send to the API endpoint
for the api extras topology maps create operation typically these are written to a http.Request
*/
type APIExtrasTopologyMapsCreateParams struct {

	/*Data*/
	Data *models.WritableTopologyMap

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the api extras topology maps create params
func (o *APIExtrasTopologyMapsCreateParams) WithTimeout(timeout time.Duration) *APIExtrasTopologyMapsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the api extras topology maps create params
func (o *APIExtrasTopologyMapsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the api extras topology maps create params
func (o *APIExtrasTopologyMapsCreateParams) WithContext(ctx context.Context) *APIExtrasTopologyMapsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the api extras topology maps create params
func (o *APIExtrasTopologyMapsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the api extras topology maps create params
func (o *APIExtrasTopologyMapsCreateParams) WithHTTPClient(client *http.Client) *APIExtrasTopologyMapsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the api extras topology maps create params
func (o *APIExtrasTopologyMapsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the api extras topology maps create params
func (o *APIExtrasTopologyMapsCreateParams) WithData(data *models.WritableTopologyMap) *APIExtrasTopologyMapsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the api extras topology maps create params
func (o *APIExtrasTopologyMapsCreateParams) SetData(data *models.WritableTopologyMap) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *APIExtrasTopologyMapsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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