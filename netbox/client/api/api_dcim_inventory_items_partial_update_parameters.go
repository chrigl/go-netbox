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

// NewAPIDcimInventoryItemsPartialUpdateParams creates a new APIDcimInventoryItemsPartialUpdateParams object
// with the default values initialized.
func NewAPIDcimInventoryItemsPartialUpdateParams() *APIDcimInventoryItemsPartialUpdateParams {
	var ()
	return &APIDcimInventoryItemsPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIDcimInventoryItemsPartialUpdateParamsWithTimeout creates a new APIDcimInventoryItemsPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIDcimInventoryItemsPartialUpdateParamsWithTimeout(timeout time.Duration) *APIDcimInventoryItemsPartialUpdateParams {
	var ()
	return &APIDcimInventoryItemsPartialUpdateParams{

		timeout: timeout,
	}
}

// NewAPIDcimInventoryItemsPartialUpdateParamsWithContext creates a new APIDcimInventoryItemsPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIDcimInventoryItemsPartialUpdateParamsWithContext(ctx context.Context) *APIDcimInventoryItemsPartialUpdateParams {
	var ()
	return &APIDcimInventoryItemsPartialUpdateParams{

		Context: ctx,
	}
}

// NewAPIDcimInventoryItemsPartialUpdateParamsWithHTTPClient creates a new APIDcimInventoryItemsPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIDcimInventoryItemsPartialUpdateParamsWithHTTPClient(client *http.Client) *APIDcimInventoryItemsPartialUpdateParams {
	var ()
	return &APIDcimInventoryItemsPartialUpdateParams{
		HTTPClient: client,
	}
}

/*APIDcimInventoryItemsPartialUpdateParams contains all the parameters to send to the API endpoint
for the api dcim inventory items partial update operation typically these are written to a http.Request
*/
type APIDcimInventoryItemsPartialUpdateParams struct {

	/*Data*/
	Data *models.WritableInventoryItem
	/*ID
	  A unique integer value identifying this inventory item.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the api dcim inventory items partial update params
func (o *APIDcimInventoryItemsPartialUpdateParams) WithTimeout(timeout time.Duration) *APIDcimInventoryItemsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the api dcim inventory items partial update params
func (o *APIDcimInventoryItemsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the api dcim inventory items partial update params
func (o *APIDcimInventoryItemsPartialUpdateParams) WithContext(ctx context.Context) *APIDcimInventoryItemsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the api dcim inventory items partial update params
func (o *APIDcimInventoryItemsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the api dcim inventory items partial update params
func (o *APIDcimInventoryItemsPartialUpdateParams) WithHTTPClient(client *http.Client) *APIDcimInventoryItemsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the api dcim inventory items partial update params
func (o *APIDcimInventoryItemsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the api dcim inventory items partial update params
func (o *APIDcimInventoryItemsPartialUpdateParams) WithData(data *models.WritableInventoryItem) *APIDcimInventoryItemsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the api dcim inventory items partial update params
func (o *APIDcimInventoryItemsPartialUpdateParams) SetData(data *models.WritableInventoryItem) {
	o.Data = data
}

// WithID adds the id to the api dcim inventory items partial update params
func (o *APIDcimInventoryItemsPartialUpdateParams) WithID(id int64) *APIDcimInventoryItemsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the api dcim inventory items partial update params
func (o *APIDcimInventoryItemsPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *APIDcimInventoryItemsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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