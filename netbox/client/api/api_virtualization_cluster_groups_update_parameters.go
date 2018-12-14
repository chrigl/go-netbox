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

// NewAPIVirtualizationClusterGroupsUpdateParams creates a new APIVirtualizationClusterGroupsUpdateParams object
// with the default values initialized.
func NewAPIVirtualizationClusterGroupsUpdateParams() *APIVirtualizationClusterGroupsUpdateParams {
	var ()
	return &APIVirtualizationClusterGroupsUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIVirtualizationClusterGroupsUpdateParamsWithTimeout creates a new APIVirtualizationClusterGroupsUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIVirtualizationClusterGroupsUpdateParamsWithTimeout(timeout time.Duration) *APIVirtualizationClusterGroupsUpdateParams {
	var ()
	return &APIVirtualizationClusterGroupsUpdateParams{

		timeout: timeout,
	}
}

// NewAPIVirtualizationClusterGroupsUpdateParamsWithContext creates a new APIVirtualizationClusterGroupsUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIVirtualizationClusterGroupsUpdateParamsWithContext(ctx context.Context) *APIVirtualizationClusterGroupsUpdateParams {
	var ()
	return &APIVirtualizationClusterGroupsUpdateParams{

		Context: ctx,
	}
}

// NewAPIVirtualizationClusterGroupsUpdateParamsWithHTTPClient creates a new APIVirtualizationClusterGroupsUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIVirtualizationClusterGroupsUpdateParamsWithHTTPClient(client *http.Client) *APIVirtualizationClusterGroupsUpdateParams {
	var ()
	return &APIVirtualizationClusterGroupsUpdateParams{
		HTTPClient: client,
	}
}

/*APIVirtualizationClusterGroupsUpdateParams contains all the parameters to send to the API endpoint
for the api virtualization cluster groups update operation typically these are written to a http.Request
*/
type APIVirtualizationClusterGroupsUpdateParams struct {

	/*Data*/
	Data *models.ClusterGroup
	/*ID
	  A unique integer value identifying this cluster group.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the api virtualization cluster groups update params
func (o *APIVirtualizationClusterGroupsUpdateParams) WithTimeout(timeout time.Duration) *APIVirtualizationClusterGroupsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the api virtualization cluster groups update params
func (o *APIVirtualizationClusterGroupsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the api virtualization cluster groups update params
func (o *APIVirtualizationClusterGroupsUpdateParams) WithContext(ctx context.Context) *APIVirtualizationClusterGroupsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the api virtualization cluster groups update params
func (o *APIVirtualizationClusterGroupsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the api virtualization cluster groups update params
func (o *APIVirtualizationClusterGroupsUpdateParams) WithHTTPClient(client *http.Client) *APIVirtualizationClusterGroupsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the api virtualization cluster groups update params
func (o *APIVirtualizationClusterGroupsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the api virtualization cluster groups update params
func (o *APIVirtualizationClusterGroupsUpdateParams) WithData(data *models.ClusterGroup) *APIVirtualizationClusterGroupsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the api virtualization cluster groups update params
func (o *APIVirtualizationClusterGroupsUpdateParams) SetData(data *models.ClusterGroup) {
	o.Data = data
}

// WithID adds the id to the api virtualization cluster groups update params
func (o *APIVirtualizationClusterGroupsUpdateParams) WithID(id int64) *APIVirtualizationClusterGroupsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the api virtualization cluster groups update params
func (o *APIVirtualizationClusterGroupsUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *APIVirtualizationClusterGroupsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
