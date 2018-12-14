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

// NewAPIDcimFrontPortsListParams creates a new APIDcimFrontPortsListParams object
// with the default values initialized.
func NewAPIDcimFrontPortsListParams() *APIDcimFrontPortsListParams {
	var ()
	return &APIDcimFrontPortsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIDcimFrontPortsListParamsWithTimeout creates a new APIDcimFrontPortsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIDcimFrontPortsListParamsWithTimeout(timeout time.Duration) *APIDcimFrontPortsListParams {
	var ()
	return &APIDcimFrontPortsListParams{

		timeout: timeout,
	}
}

// NewAPIDcimFrontPortsListParamsWithContext creates a new APIDcimFrontPortsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIDcimFrontPortsListParamsWithContext(ctx context.Context) *APIDcimFrontPortsListParams {
	var ()
	return &APIDcimFrontPortsListParams{

		Context: ctx,
	}
}

// NewAPIDcimFrontPortsListParamsWithHTTPClient creates a new APIDcimFrontPortsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIDcimFrontPortsListParamsWithHTTPClient(client *http.Client) *APIDcimFrontPortsListParams {
	var ()
	return &APIDcimFrontPortsListParams{
		HTTPClient: client,
	}
}

/*APIDcimFrontPortsListParams contains all the parameters to send to the API endpoint
for the api dcim front ports list operation typically these are written to a http.Request
*/
type APIDcimFrontPortsListParams struct {

	/*Cabled*/
	Cabled *string
	/*Device*/
	Device *string
	/*DeviceID*/
	DeviceID *string
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
	/*Tag*/
	Tag *string
	/*Type*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the api dcim front ports list params
func (o *APIDcimFrontPortsListParams) WithTimeout(timeout time.Duration) *APIDcimFrontPortsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the api dcim front ports list params
func (o *APIDcimFrontPortsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the api dcim front ports list params
func (o *APIDcimFrontPortsListParams) WithContext(ctx context.Context) *APIDcimFrontPortsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the api dcim front ports list params
func (o *APIDcimFrontPortsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the api dcim front ports list params
func (o *APIDcimFrontPortsListParams) WithHTTPClient(client *http.Client) *APIDcimFrontPortsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the api dcim front ports list params
func (o *APIDcimFrontPortsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCabled adds the cabled to the api dcim front ports list params
func (o *APIDcimFrontPortsListParams) WithCabled(cabled *string) *APIDcimFrontPortsListParams {
	o.SetCabled(cabled)
	return o
}

// SetCabled adds the cabled to the api dcim front ports list params
func (o *APIDcimFrontPortsListParams) SetCabled(cabled *string) {
	o.Cabled = cabled
}

// WithDevice adds the device to the api dcim front ports list params
func (o *APIDcimFrontPortsListParams) WithDevice(device *string) *APIDcimFrontPortsListParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the api dcim front ports list params
func (o *APIDcimFrontPortsListParams) SetDevice(device *string) {
	o.Device = device
}

// WithDeviceID adds the deviceID to the api dcim front ports list params
func (o *APIDcimFrontPortsListParams) WithDeviceID(deviceID *string) *APIDcimFrontPortsListParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the api dcim front ports list params
func (o *APIDcimFrontPortsListParams) SetDeviceID(deviceID *string) {
	o.DeviceID = deviceID
}

// WithLimit adds the limit to the api dcim front ports list params
func (o *APIDcimFrontPortsListParams) WithLimit(limit *int64) *APIDcimFrontPortsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the api dcim front ports list params
func (o *APIDcimFrontPortsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the api dcim front ports list params
func (o *APIDcimFrontPortsListParams) WithName(name *string) *APIDcimFrontPortsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the api dcim front ports list params
func (o *APIDcimFrontPortsListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the api dcim front ports list params
func (o *APIDcimFrontPortsListParams) WithOffset(offset *int64) *APIDcimFrontPortsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the api dcim front ports list params
func (o *APIDcimFrontPortsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithTag adds the tag to the api dcim front ports list params
func (o *APIDcimFrontPortsListParams) WithTag(tag *string) *APIDcimFrontPortsListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the api dcim front ports list params
func (o *APIDcimFrontPortsListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithType adds the typeVar to the api dcim front ports list params
func (o *APIDcimFrontPortsListParams) WithType(typeVar *string) *APIDcimFrontPortsListParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the api dcim front ports list params
func (o *APIDcimFrontPortsListParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *APIDcimFrontPortsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cabled != nil {

		// query param cabled
		var qrCabled string
		if o.Cabled != nil {
			qrCabled = *o.Cabled
		}
		qCabled := qrCabled
		if qCabled != "" {
			if err := r.SetQueryParam("cabled", qCabled); err != nil {
				return err
			}
		}

	}

	if o.Device != nil {

		// query param device
		var qrDevice string
		if o.Device != nil {
			qrDevice = *o.Device
		}
		qDevice := qrDevice
		if qDevice != "" {
			if err := r.SetQueryParam("device", qDevice); err != nil {
				return err
			}
		}

	}

	if o.DeviceID != nil {

		// query param device_id
		var qrDeviceID string
		if o.DeviceID != nil {
			qrDeviceID = *o.DeviceID
		}
		qDeviceID := qrDeviceID
		if qDeviceID != "" {
			if err := r.SetQueryParam("device_id", qDeviceID); err != nil {
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

	if o.Tag != nil {

		// query param tag
		var qrTag string
		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {
			if err := r.SetQueryParam("tag", qTag); err != nil {
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