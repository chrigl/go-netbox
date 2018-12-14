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

// NewAPIDcimSitesGraphsParams creates a new APIDcimSitesGraphsParams object
// with the default values initialized.
func NewAPIDcimSitesGraphsParams() *APIDcimSitesGraphsParams {
	var ()
	return &APIDcimSitesGraphsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIDcimSitesGraphsParamsWithTimeout creates a new APIDcimSitesGraphsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIDcimSitesGraphsParamsWithTimeout(timeout time.Duration) *APIDcimSitesGraphsParams {
	var ()
	return &APIDcimSitesGraphsParams{

		timeout: timeout,
	}
}

// NewAPIDcimSitesGraphsParamsWithContext creates a new APIDcimSitesGraphsParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIDcimSitesGraphsParamsWithContext(ctx context.Context) *APIDcimSitesGraphsParams {
	var ()
	return &APIDcimSitesGraphsParams{

		Context: ctx,
	}
}

// NewAPIDcimSitesGraphsParamsWithHTTPClient creates a new APIDcimSitesGraphsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIDcimSitesGraphsParamsWithHTTPClient(client *http.Client) *APIDcimSitesGraphsParams {
	var ()
	return &APIDcimSitesGraphsParams{
		HTTPClient: client,
	}
}

/*APIDcimSitesGraphsParams contains all the parameters to send to the API endpoint
for the api dcim sites graphs operation typically these are written to a http.Request
*/
type APIDcimSitesGraphsParams struct {

	/*ID
	  A unique integer value identifying this site.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the api dcim sites graphs params
func (o *APIDcimSitesGraphsParams) WithTimeout(timeout time.Duration) *APIDcimSitesGraphsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the api dcim sites graphs params
func (o *APIDcimSitesGraphsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the api dcim sites graphs params
func (o *APIDcimSitesGraphsParams) WithContext(ctx context.Context) *APIDcimSitesGraphsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the api dcim sites graphs params
func (o *APIDcimSitesGraphsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the api dcim sites graphs params
func (o *APIDcimSitesGraphsParams) WithHTTPClient(client *http.Client) *APIDcimSitesGraphsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the api dcim sites graphs params
func (o *APIDcimSitesGraphsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the api dcim sites graphs params
func (o *APIDcimSitesGraphsParams) WithID(id int64) *APIDcimSitesGraphsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the api dcim sites graphs params
func (o *APIDcimSitesGraphsParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *APIDcimSitesGraphsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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