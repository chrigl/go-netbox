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

// NewAPIDcimDeviceTypesListParams creates a new APIDcimDeviceTypesListParams object
// with the default values initialized.
func NewAPIDcimDeviceTypesListParams() *APIDcimDeviceTypesListParams {
	var ()
	return &APIDcimDeviceTypesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIDcimDeviceTypesListParamsWithTimeout creates a new APIDcimDeviceTypesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIDcimDeviceTypesListParamsWithTimeout(timeout time.Duration) *APIDcimDeviceTypesListParams {
	var ()
	return &APIDcimDeviceTypesListParams{

		timeout: timeout,
	}
}

// NewAPIDcimDeviceTypesListParamsWithContext creates a new APIDcimDeviceTypesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIDcimDeviceTypesListParamsWithContext(ctx context.Context) *APIDcimDeviceTypesListParams {
	var ()
	return &APIDcimDeviceTypesListParams{

		Context: ctx,
	}
}

// NewAPIDcimDeviceTypesListParamsWithHTTPClient creates a new APIDcimDeviceTypesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIDcimDeviceTypesListParamsWithHTTPClient(client *http.Client) *APIDcimDeviceTypesListParams {
	var ()
	return &APIDcimDeviceTypesListParams{
		HTTPClient: client,
	}
}

/*APIDcimDeviceTypesListParams contains all the parameters to send to the API endpoint
for the api dcim device types list operation typically these are written to a http.Request
*/
type APIDcimDeviceTypesListParams struct {

	/*ConsolePorts*/
	ConsolePorts *string
	/*ConsoleServerPorts*/
	ConsoleServerPorts *string
	/*IDIn
	  Multiple values may be separated by commas.

	*/
	IDIn *string
	/*Interfaces*/
	Interfaces *string
	/*IsFullDepth*/
	IsFullDepth *string
	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*Manufacturer*/
	Manufacturer *string
	/*ManufacturerID*/
	ManufacturerID *string
	/*Model*/
	Model *string
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*PartNumber*/
	PartNumber *string
	/*PassThroughPorts*/
	PassThroughPorts *string
	/*PowerOutlets*/
	PowerOutlets *string
	/*PowerPorts*/
	PowerPorts *string
	/*Q*/
	Q *string
	/*Slug*/
	Slug *string
	/*SubdeviceRole*/
	SubdeviceRole *string
	/*Tag*/
	Tag *string
	/*UHeight*/
	UHeight *float64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) WithTimeout(timeout time.Duration) *APIDcimDeviceTypesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) WithContext(ctx context.Context) *APIDcimDeviceTypesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) WithHTTPClient(client *http.Client) *APIDcimDeviceTypesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConsolePorts adds the consolePorts to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) WithConsolePorts(consolePorts *string) *APIDcimDeviceTypesListParams {
	o.SetConsolePorts(consolePorts)
	return o
}

// SetConsolePorts adds the consolePorts to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) SetConsolePorts(consolePorts *string) {
	o.ConsolePorts = consolePorts
}

// WithConsoleServerPorts adds the consoleServerPorts to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) WithConsoleServerPorts(consoleServerPorts *string) *APIDcimDeviceTypesListParams {
	o.SetConsoleServerPorts(consoleServerPorts)
	return o
}

// SetConsoleServerPorts adds the consoleServerPorts to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) SetConsoleServerPorts(consoleServerPorts *string) {
	o.ConsoleServerPorts = consoleServerPorts
}

// WithIDIn adds the iDIn to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) WithIDIn(iDIn *string) *APIDcimDeviceTypesListParams {
	o.SetIDIn(iDIn)
	return o
}

// SetIDIn adds the idIn to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) SetIDIn(iDIn *string) {
	o.IDIn = iDIn
}

// WithInterfaces adds the interfaces to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) WithInterfaces(interfaces *string) *APIDcimDeviceTypesListParams {
	o.SetInterfaces(interfaces)
	return o
}

// SetInterfaces adds the interfaces to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) SetInterfaces(interfaces *string) {
	o.Interfaces = interfaces
}

// WithIsFullDepth adds the isFullDepth to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) WithIsFullDepth(isFullDepth *string) *APIDcimDeviceTypesListParams {
	o.SetIsFullDepth(isFullDepth)
	return o
}

// SetIsFullDepth adds the isFullDepth to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) SetIsFullDepth(isFullDepth *string) {
	o.IsFullDepth = isFullDepth
}

// WithLimit adds the limit to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) WithLimit(limit *int64) *APIDcimDeviceTypesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithManufacturer adds the manufacturer to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) WithManufacturer(manufacturer *string) *APIDcimDeviceTypesListParams {
	o.SetManufacturer(manufacturer)
	return o
}

// SetManufacturer adds the manufacturer to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) SetManufacturer(manufacturer *string) {
	o.Manufacturer = manufacturer
}

// WithManufacturerID adds the manufacturerID to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) WithManufacturerID(manufacturerID *string) *APIDcimDeviceTypesListParams {
	o.SetManufacturerID(manufacturerID)
	return o
}

// SetManufacturerID adds the manufacturerId to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) SetManufacturerID(manufacturerID *string) {
	o.ManufacturerID = manufacturerID
}

// WithModel adds the model to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) WithModel(model *string) *APIDcimDeviceTypesListParams {
	o.SetModel(model)
	return o
}

// SetModel adds the model to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) SetModel(model *string) {
	o.Model = model
}

// WithOffset adds the offset to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) WithOffset(offset *int64) *APIDcimDeviceTypesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPartNumber adds the partNumber to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) WithPartNumber(partNumber *string) *APIDcimDeviceTypesListParams {
	o.SetPartNumber(partNumber)
	return o
}

// SetPartNumber adds the partNumber to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) SetPartNumber(partNumber *string) {
	o.PartNumber = partNumber
}

// WithPassThroughPorts adds the passThroughPorts to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) WithPassThroughPorts(passThroughPorts *string) *APIDcimDeviceTypesListParams {
	o.SetPassThroughPorts(passThroughPorts)
	return o
}

// SetPassThroughPorts adds the passThroughPorts to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) SetPassThroughPorts(passThroughPorts *string) {
	o.PassThroughPorts = passThroughPorts
}

// WithPowerOutlets adds the powerOutlets to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) WithPowerOutlets(powerOutlets *string) *APIDcimDeviceTypesListParams {
	o.SetPowerOutlets(powerOutlets)
	return o
}

// SetPowerOutlets adds the powerOutlets to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) SetPowerOutlets(powerOutlets *string) {
	o.PowerOutlets = powerOutlets
}

// WithPowerPorts adds the powerPorts to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) WithPowerPorts(powerPorts *string) *APIDcimDeviceTypesListParams {
	o.SetPowerPorts(powerPorts)
	return o
}

// SetPowerPorts adds the powerPorts to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) SetPowerPorts(powerPorts *string) {
	o.PowerPorts = powerPorts
}

// WithQ adds the q to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) WithQ(q *string) *APIDcimDeviceTypesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) SetQ(q *string) {
	o.Q = q
}

// WithSlug adds the slug to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) WithSlug(slug *string) *APIDcimDeviceTypesListParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) SetSlug(slug *string) {
	o.Slug = slug
}

// WithSubdeviceRole adds the subdeviceRole to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) WithSubdeviceRole(subdeviceRole *string) *APIDcimDeviceTypesListParams {
	o.SetSubdeviceRole(subdeviceRole)
	return o
}

// SetSubdeviceRole adds the subdeviceRole to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) SetSubdeviceRole(subdeviceRole *string) {
	o.SubdeviceRole = subdeviceRole
}

// WithTag adds the tag to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) WithTag(tag *string) *APIDcimDeviceTypesListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithUHeight adds the uHeight to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) WithUHeight(uHeight *float64) *APIDcimDeviceTypesListParams {
	o.SetUHeight(uHeight)
	return o
}

// SetUHeight adds the uHeight to the api dcim device types list params
func (o *APIDcimDeviceTypesListParams) SetUHeight(uHeight *float64) {
	o.UHeight = uHeight
}

// WriteToRequest writes these params to a swagger request
func (o *APIDcimDeviceTypesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ConsolePorts != nil {

		// query param console_ports
		var qrConsolePorts string
		if o.ConsolePorts != nil {
			qrConsolePorts = *o.ConsolePorts
		}
		qConsolePorts := qrConsolePorts
		if qConsolePorts != "" {
			if err := r.SetQueryParam("console_ports", qConsolePorts); err != nil {
				return err
			}
		}

	}

	if o.ConsoleServerPorts != nil {

		// query param console_server_ports
		var qrConsoleServerPorts string
		if o.ConsoleServerPorts != nil {
			qrConsoleServerPorts = *o.ConsoleServerPorts
		}
		qConsoleServerPorts := qrConsoleServerPorts
		if qConsoleServerPorts != "" {
			if err := r.SetQueryParam("console_server_ports", qConsoleServerPorts); err != nil {
				return err
			}
		}

	}

	if o.IDIn != nil {

		// query param id__in
		var qrIDIn string
		if o.IDIn != nil {
			qrIDIn = *o.IDIn
		}
		qIDIn := qrIDIn
		if qIDIn != "" {
			if err := r.SetQueryParam("id__in", qIDIn); err != nil {
				return err
			}
		}

	}

	if o.Interfaces != nil {

		// query param interfaces
		var qrInterfaces string
		if o.Interfaces != nil {
			qrInterfaces = *o.Interfaces
		}
		qInterfaces := qrInterfaces
		if qInterfaces != "" {
			if err := r.SetQueryParam("interfaces", qInterfaces); err != nil {
				return err
			}
		}

	}

	if o.IsFullDepth != nil {

		// query param is_full_depth
		var qrIsFullDepth string
		if o.IsFullDepth != nil {
			qrIsFullDepth = *o.IsFullDepth
		}
		qIsFullDepth := qrIsFullDepth
		if qIsFullDepth != "" {
			if err := r.SetQueryParam("is_full_depth", qIsFullDepth); err != nil {
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

	if o.Manufacturer != nil {

		// query param manufacturer
		var qrManufacturer string
		if o.Manufacturer != nil {
			qrManufacturer = *o.Manufacturer
		}
		qManufacturer := qrManufacturer
		if qManufacturer != "" {
			if err := r.SetQueryParam("manufacturer", qManufacturer); err != nil {
				return err
			}
		}

	}

	if o.ManufacturerID != nil {

		// query param manufacturer_id
		var qrManufacturerID string
		if o.ManufacturerID != nil {
			qrManufacturerID = *o.ManufacturerID
		}
		qManufacturerID := qrManufacturerID
		if qManufacturerID != "" {
			if err := r.SetQueryParam("manufacturer_id", qManufacturerID); err != nil {
				return err
			}
		}

	}

	if o.Model != nil {

		// query param model
		var qrModel string
		if o.Model != nil {
			qrModel = *o.Model
		}
		qModel := qrModel
		if qModel != "" {
			if err := r.SetQueryParam("model", qModel); err != nil {
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

	if o.PartNumber != nil {

		// query param part_number
		var qrPartNumber string
		if o.PartNumber != nil {
			qrPartNumber = *o.PartNumber
		}
		qPartNumber := qrPartNumber
		if qPartNumber != "" {
			if err := r.SetQueryParam("part_number", qPartNumber); err != nil {
				return err
			}
		}

	}

	if o.PassThroughPorts != nil {

		// query param pass_through_ports
		var qrPassThroughPorts string
		if o.PassThroughPorts != nil {
			qrPassThroughPorts = *o.PassThroughPorts
		}
		qPassThroughPorts := qrPassThroughPorts
		if qPassThroughPorts != "" {
			if err := r.SetQueryParam("pass_through_ports", qPassThroughPorts); err != nil {
				return err
			}
		}

	}

	if o.PowerOutlets != nil {

		// query param power_outlets
		var qrPowerOutlets string
		if o.PowerOutlets != nil {
			qrPowerOutlets = *o.PowerOutlets
		}
		qPowerOutlets := qrPowerOutlets
		if qPowerOutlets != "" {
			if err := r.SetQueryParam("power_outlets", qPowerOutlets); err != nil {
				return err
			}
		}

	}

	if o.PowerPorts != nil {

		// query param power_ports
		var qrPowerPorts string
		if o.PowerPorts != nil {
			qrPowerPorts = *o.PowerPorts
		}
		qPowerPorts := qrPowerPorts
		if qPowerPorts != "" {
			if err := r.SetQueryParam("power_ports", qPowerPorts); err != nil {
				return err
			}
		}

	}

	if o.Q != nil {

		// query param q
		var qrQ string
		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {
			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}

	}

	if o.Slug != nil {

		// query param slug
		var qrSlug string
		if o.Slug != nil {
			qrSlug = *o.Slug
		}
		qSlug := qrSlug
		if qSlug != "" {
			if err := r.SetQueryParam("slug", qSlug); err != nil {
				return err
			}
		}

	}

	if o.SubdeviceRole != nil {

		// query param subdevice_role
		var qrSubdeviceRole string
		if o.SubdeviceRole != nil {
			qrSubdeviceRole = *o.SubdeviceRole
		}
		qSubdeviceRole := qrSubdeviceRole
		if qSubdeviceRole != "" {
			if err := r.SetQueryParam("subdevice_role", qSubdeviceRole); err != nil {
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

	if o.UHeight != nil {

		// query param u_height
		var qrUHeight float64
		if o.UHeight != nil {
			qrUHeight = *o.UHeight
		}
		qUHeight := swag.FormatFloat64(qrUHeight)
		if qUHeight != "" {
			if err := r.SetQueryParam("u_height", qUHeight); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}