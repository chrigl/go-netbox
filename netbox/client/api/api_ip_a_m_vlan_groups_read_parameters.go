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
)

// NewAPIIPAMVlanGroupsReadParams creates a new APIIPAMVlanGroupsReadParams object
// with the default values initialized.
func NewAPIIPAMVlanGroupsReadParams() *APIIPAMVlanGroupsReadParams {
	var ()
	return &APIIPAMVlanGroupsReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIIPAMVlanGroupsReadParamsWithTimeout creates a new APIIPAMVlanGroupsReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIIPAMVlanGroupsReadParamsWithTimeout(timeout time.Duration) *APIIPAMVlanGroupsReadParams {
	var ()
	return &APIIPAMVlanGroupsReadParams{

		timeout: timeout,
	}
}

// NewAPIIPAMVlanGroupsReadParamsWithContext creates a new APIIPAMVlanGroupsReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIIPAMVlanGroupsReadParamsWithContext(ctx context.Context) *APIIPAMVlanGroupsReadParams {
	var ()
	return &APIIPAMVlanGroupsReadParams{

		Context: ctx,
	}
}

// NewAPIIPAMVlanGroupsReadParamsWithHTTPClient creates a new APIIPAMVlanGroupsReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIIPAMVlanGroupsReadParamsWithHTTPClient(client *http.Client) *APIIPAMVlanGroupsReadParams {
	var ()
	return &APIIPAMVlanGroupsReadParams{
		HTTPClient: client,
	}
}

/*APIIPAMVlanGroupsReadParams contains all the parameters to send to the API endpoint
for the api ipam vlan groups read operation typically these are written to a http.Request
*/
type APIIPAMVlanGroupsReadParams struct {

	/*ID
	  A unique integer value identifying this VLAN group.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the api ipam vlan groups read params
func (o *APIIPAMVlanGroupsReadParams) WithTimeout(timeout time.Duration) *APIIPAMVlanGroupsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the api ipam vlan groups read params
func (o *APIIPAMVlanGroupsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the api ipam vlan groups read params
func (o *APIIPAMVlanGroupsReadParams) WithContext(ctx context.Context) *APIIPAMVlanGroupsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the api ipam vlan groups read params
func (o *APIIPAMVlanGroupsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the api ipam vlan groups read params
func (o *APIIPAMVlanGroupsReadParams) WithHTTPClient(client *http.Client) *APIIPAMVlanGroupsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the api ipam vlan groups read params
func (o *APIIPAMVlanGroupsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the api ipam vlan groups read params
func (o *APIIPAMVlanGroupsReadParams) WithID(id int64) *APIIPAMVlanGroupsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the api ipam vlan groups read params
func (o *APIIPAMVlanGroupsReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *APIIPAMVlanGroupsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}