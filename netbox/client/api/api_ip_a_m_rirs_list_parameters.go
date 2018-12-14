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

// NewAPIIPAMRirsListParams creates a new APIIPAMRirsListParams object
// with the default values initialized.
func NewAPIIPAMRirsListParams() *APIIPAMRirsListParams {
	var ()
	return &APIIPAMRirsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIIPAMRirsListParamsWithTimeout creates a new APIIPAMRirsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIIPAMRirsListParamsWithTimeout(timeout time.Duration) *APIIPAMRirsListParams {
	var ()
	return &APIIPAMRirsListParams{

		timeout: timeout,
	}
}

// NewAPIIPAMRirsListParamsWithContext creates a new APIIPAMRirsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIIPAMRirsListParamsWithContext(ctx context.Context) *APIIPAMRirsListParams {
	var ()
	return &APIIPAMRirsListParams{

		Context: ctx,
	}
}

// NewAPIIPAMRirsListParamsWithHTTPClient creates a new APIIPAMRirsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIIPAMRirsListParamsWithHTTPClient(client *http.Client) *APIIPAMRirsListParams {
	var ()
	return &APIIPAMRirsListParams{
		HTTPClient: client,
	}
}

/*APIIPAMRirsListParams contains all the parameters to send to the API endpoint
for the api ipam rirs list operation typically these are written to a http.Request
*/
type APIIPAMRirsListParams struct {

	/*IDIn
	  Multiple values may be separated by commas.

	*/
	IDIn *string
	/*IsPrivate*/
	IsPrivate *string
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
	/*Slug*/
	Slug *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the api ipam rirs list params
func (o *APIIPAMRirsListParams) WithTimeout(timeout time.Duration) *APIIPAMRirsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the api ipam rirs list params
func (o *APIIPAMRirsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the api ipam rirs list params
func (o *APIIPAMRirsListParams) WithContext(ctx context.Context) *APIIPAMRirsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the api ipam rirs list params
func (o *APIIPAMRirsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the api ipam rirs list params
func (o *APIIPAMRirsListParams) WithHTTPClient(client *http.Client) *APIIPAMRirsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the api ipam rirs list params
func (o *APIIPAMRirsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIDIn adds the iDIn to the api ipam rirs list params
func (o *APIIPAMRirsListParams) WithIDIn(iDIn *string) *APIIPAMRirsListParams {
	o.SetIDIn(iDIn)
	return o
}

// SetIDIn adds the idIn to the api ipam rirs list params
func (o *APIIPAMRirsListParams) SetIDIn(iDIn *string) {
	o.IDIn = iDIn
}

// WithIsPrivate adds the isPrivate to the api ipam rirs list params
func (o *APIIPAMRirsListParams) WithIsPrivate(isPrivate *string) *APIIPAMRirsListParams {
	o.SetIsPrivate(isPrivate)
	return o
}

// SetIsPrivate adds the isPrivate to the api ipam rirs list params
func (o *APIIPAMRirsListParams) SetIsPrivate(isPrivate *string) {
	o.IsPrivate = isPrivate
}

// WithLimit adds the limit to the api ipam rirs list params
func (o *APIIPAMRirsListParams) WithLimit(limit *int64) *APIIPAMRirsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the api ipam rirs list params
func (o *APIIPAMRirsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the api ipam rirs list params
func (o *APIIPAMRirsListParams) WithName(name *string) *APIIPAMRirsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the api ipam rirs list params
func (o *APIIPAMRirsListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the api ipam rirs list params
func (o *APIIPAMRirsListParams) WithOffset(offset *int64) *APIIPAMRirsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the api ipam rirs list params
func (o *APIIPAMRirsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithSlug adds the slug to the api ipam rirs list params
func (o *APIIPAMRirsListParams) WithSlug(slug *string) *APIIPAMRirsListParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the api ipam rirs list params
func (o *APIIPAMRirsListParams) SetSlug(slug *string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *APIIPAMRirsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.IsPrivate != nil {

		// query param is_private
		var qrIsPrivate string
		if o.IsPrivate != nil {
			qrIsPrivate = *o.IsPrivate
		}
		qIsPrivate := qrIsPrivate
		if qIsPrivate != "" {
			if err := r.SetQueryParam("is_private", qIsPrivate); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}