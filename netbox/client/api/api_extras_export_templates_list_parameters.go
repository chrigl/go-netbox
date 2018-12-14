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

// NewAPIExtrasExportTemplatesListParams creates a new APIExtrasExportTemplatesListParams object
// with the default values initialized.
func NewAPIExtrasExportTemplatesListParams() *APIExtrasExportTemplatesListParams {
	var ()
	return &APIExtrasExportTemplatesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIExtrasExportTemplatesListParamsWithTimeout creates a new APIExtrasExportTemplatesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIExtrasExportTemplatesListParamsWithTimeout(timeout time.Duration) *APIExtrasExportTemplatesListParams {
	var ()
	return &APIExtrasExportTemplatesListParams{

		timeout: timeout,
	}
}

// NewAPIExtrasExportTemplatesListParamsWithContext creates a new APIExtrasExportTemplatesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIExtrasExportTemplatesListParamsWithContext(ctx context.Context) *APIExtrasExportTemplatesListParams {
	var ()
	return &APIExtrasExportTemplatesListParams{

		Context: ctx,
	}
}

// NewAPIExtrasExportTemplatesListParamsWithHTTPClient creates a new APIExtrasExportTemplatesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIExtrasExportTemplatesListParamsWithHTTPClient(client *http.Client) *APIExtrasExportTemplatesListParams {
	var ()
	return &APIExtrasExportTemplatesListParams{
		HTTPClient: client,
	}
}

/*APIExtrasExportTemplatesListParams contains all the parameters to send to the API endpoint
for the api extras export templates list operation typically these are written to a http.Request
*/
type APIExtrasExportTemplatesListParams struct {

	/*ContentType*/
	ContentType *string
	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*Name*/
	Name *string
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the api extras export templates list params
func (o *APIExtrasExportTemplatesListParams) WithTimeout(timeout time.Duration) *APIExtrasExportTemplatesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the api extras export templates list params
func (o *APIExtrasExportTemplatesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the api extras export templates list params
func (o *APIExtrasExportTemplatesListParams) WithContext(ctx context.Context) *APIExtrasExportTemplatesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the api extras export templates list params
func (o *APIExtrasExportTemplatesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the api extras export templates list params
func (o *APIExtrasExportTemplatesListParams) WithHTTPClient(client *http.Client) *APIExtrasExportTemplatesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the api extras export templates list params
func (o *APIExtrasExportTemplatesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentType adds the contentType to the api extras export templates list params
func (o *APIExtrasExportTemplatesListParams) WithContentType(contentType *string) *APIExtrasExportTemplatesListParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the api extras export templates list params
func (o *APIExtrasExportTemplatesListParams) SetContentType(contentType *string) {
	o.ContentType = contentType
}

// WithLimit adds the limit to the api extras export templates list params
func (o *APIExtrasExportTemplatesListParams) WithLimit(limit *int64) *APIExtrasExportTemplatesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the api extras export templates list params
func (o *APIExtrasExportTemplatesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the api extras export templates list params
func (o *APIExtrasExportTemplatesListParams) WithName(name *string) *APIExtrasExportTemplatesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the api extras export templates list params
func (o *APIExtrasExportTemplatesListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the api extras export templates list params
func (o *APIExtrasExportTemplatesListParams) WithOffset(offset *int64) *APIExtrasExportTemplatesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the api extras export templates list params
func (o *APIExtrasExportTemplatesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *APIExtrasExportTemplatesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ContentType != nil {

		// query param content_type
		var qrContentType string
		if o.ContentType != nil {
			qrContentType = *o.ContentType
		}
		qContentType := qrContentType
		if qContentType != "" {
			if err := r.SetQueryParam("content_type", qContentType); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
