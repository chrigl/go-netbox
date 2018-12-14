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

// NewAPIDcimRearPortTemplatesListParams creates a new APIDcimRearPortTemplatesListParams object
// with the default values initialized.
func NewAPIDcimRearPortTemplatesListParams() *APIDcimRearPortTemplatesListParams {
	var ()
	return &APIDcimRearPortTemplatesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIDcimRearPortTemplatesListParamsWithTimeout creates a new APIDcimRearPortTemplatesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIDcimRearPortTemplatesListParamsWithTimeout(timeout time.Duration) *APIDcimRearPortTemplatesListParams {
	var ()
	return &APIDcimRearPortTemplatesListParams{

		timeout: timeout,
	}
}

// NewAPIDcimRearPortTemplatesListParamsWithContext creates a new APIDcimRearPortTemplatesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIDcimRearPortTemplatesListParamsWithContext(ctx context.Context) *APIDcimRearPortTemplatesListParams {
	var ()
	return &APIDcimRearPortTemplatesListParams{

		Context: ctx,
	}
}

// NewAPIDcimRearPortTemplatesListParamsWithHTTPClient creates a new APIDcimRearPortTemplatesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIDcimRearPortTemplatesListParamsWithHTTPClient(client *http.Client) *APIDcimRearPortTemplatesListParams {
	var ()
	return &APIDcimRearPortTemplatesListParams{
		HTTPClient: client,
	}
}

/*APIDcimRearPortTemplatesListParams contains all the parameters to send to the API endpoint
for the api dcim rear port templates list operation typically these are written to a http.Request
*/
type APIDcimRearPortTemplatesListParams struct {

	/*DevicetypeID*/
	DevicetypeID *string
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
	/*Type*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the api dcim rear port templates list params
func (o *APIDcimRearPortTemplatesListParams) WithTimeout(timeout time.Duration) *APIDcimRearPortTemplatesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the api dcim rear port templates list params
func (o *APIDcimRearPortTemplatesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the api dcim rear port templates list params
func (o *APIDcimRearPortTemplatesListParams) WithContext(ctx context.Context) *APIDcimRearPortTemplatesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the api dcim rear port templates list params
func (o *APIDcimRearPortTemplatesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the api dcim rear port templates list params
func (o *APIDcimRearPortTemplatesListParams) WithHTTPClient(client *http.Client) *APIDcimRearPortTemplatesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the api dcim rear port templates list params
func (o *APIDcimRearPortTemplatesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDevicetypeID adds the devicetypeID to the api dcim rear port templates list params
func (o *APIDcimRearPortTemplatesListParams) WithDevicetypeID(devicetypeID *string) *APIDcimRearPortTemplatesListParams {
	o.SetDevicetypeID(devicetypeID)
	return o
}

// SetDevicetypeID adds the devicetypeId to the api dcim rear port templates list params
func (o *APIDcimRearPortTemplatesListParams) SetDevicetypeID(devicetypeID *string) {
	o.DevicetypeID = devicetypeID
}

// WithLimit adds the limit to the api dcim rear port templates list params
func (o *APIDcimRearPortTemplatesListParams) WithLimit(limit *int64) *APIDcimRearPortTemplatesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the api dcim rear port templates list params
func (o *APIDcimRearPortTemplatesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the api dcim rear port templates list params
func (o *APIDcimRearPortTemplatesListParams) WithName(name *string) *APIDcimRearPortTemplatesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the api dcim rear port templates list params
func (o *APIDcimRearPortTemplatesListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the api dcim rear port templates list params
func (o *APIDcimRearPortTemplatesListParams) WithOffset(offset *int64) *APIDcimRearPortTemplatesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the api dcim rear port templates list params
func (o *APIDcimRearPortTemplatesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithType adds the typeVar to the api dcim rear port templates list params
func (o *APIDcimRearPortTemplatesListParams) WithType(typeVar *string) *APIDcimRearPortTemplatesListParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the api dcim rear port templates list params
func (o *APIDcimRearPortTemplatesListParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *APIDcimRearPortTemplatesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DevicetypeID != nil {

		// query param devicetype_id
		var qrDevicetypeID string
		if o.DevicetypeID != nil {
			qrDevicetypeID = *o.DevicetypeID
		}
		qDevicetypeID := qrDevicetypeID
		if qDevicetypeID != "" {
			if err := r.SetQueryParam("devicetype_id", qDevicetypeID); err != nil {
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

	if o.Type != nil {

		// query param type
		var qrType string
		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {
			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}