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

// NewAPIDcimSitesListParams creates a new APIDcimSitesListParams object
// with the default values initialized.
func NewAPIDcimSitesListParams() *APIDcimSitesListParams {
	var ()
	return &APIDcimSitesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIDcimSitesListParamsWithTimeout creates a new APIDcimSitesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIDcimSitesListParamsWithTimeout(timeout time.Duration) *APIDcimSitesListParams {
	var ()
	return &APIDcimSitesListParams{

		timeout: timeout,
	}
}

// NewAPIDcimSitesListParamsWithContext creates a new APIDcimSitesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIDcimSitesListParamsWithContext(ctx context.Context) *APIDcimSitesListParams {
	var ()
	return &APIDcimSitesListParams{

		Context: ctx,
	}
}

// NewAPIDcimSitesListParamsWithHTTPClient creates a new APIDcimSitesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIDcimSitesListParamsWithHTTPClient(client *http.Client) *APIDcimSitesListParams {
	var ()
	return &APIDcimSitesListParams{
		HTTPClient: client,
	}
}

/*APIDcimSitesListParams contains all the parameters to send to the API endpoint
for the api dcim sites list operation typically these are written to a http.Request
*/
type APIDcimSitesListParams struct {

	/*Asn*/
	Asn *float64
	/*ContactEmail*/
	ContactEmail *string
	/*ContactName*/
	ContactName *string
	/*ContactPhone*/
	ContactPhone *string
	/*Facility*/
	Facility *string
	/*IDIn
	  Multiple values may be separated by commas.

	*/
	IDIn *string
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
	/*Q*/
	Q *string
	/*Region*/
	Region *string
	/*RegionID*/
	RegionID *string
	/*Slug*/
	Slug *string
	/*Status*/
	Status *string
	/*Tag*/
	Tag *string
	/*Tenant*/
	Tenant *string
	/*TenantID*/
	TenantID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the api dcim sites list params
func (o *APIDcimSitesListParams) WithTimeout(timeout time.Duration) *APIDcimSitesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the api dcim sites list params
func (o *APIDcimSitesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the api dcim sites list params
func (o *APIDcimSitesListParams) WithContext(ctx context.Context) *APIDcimSitesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the api dcim sites list params
func (o *APIDcimSitesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the api dcim sites list params
func (o *APIDcimSitesListParams) WithHTTPClient(client *http.Client) *APIDcimSitesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the api dcim sites list params
func (o *APIDcimSitesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAsn adds the asn to the api dcim sites list params
func (o *APIDcimSitesListParams) WithAsn(asn *float64) *APIDcimSitesListParams {
	o.SetAsn(asn)
	return o
}

// SetAsn adds the asn to the api dcim sites list params
func (o *APIDcimSitesListParams) SetAsn(asn *float64) {
	o.Asn = asn
}

// WithContactEmail adds the contactEmail to the api dcim sites list params
func (o *APIDcimSitesListParams) WithContactEmail(contactEmail *string) *APIDcimSitesListParams {
	o.SetContactEmail(contactEmail)
	return o
}

// SetContactEmail adds the contactEmail to the api dcim sites list params
func (o *APIDcimSitesListParams) SetContactEmail(contactEmail *string) {
	o.ContactEmail = contactEmail
}

// WithContactName adds the contactName to the api dcim sites list params
func (o *APIDcimSitesListParams) WithContactName(contactName *string) *APIDcimSitesListParams {
	o.SetContactName(contactName)
	return o
}

// SetContactName adds the contactName to the api dcim sites list params
func (o *APIDcimSitesListParams) SetContactName(contactName *string) {
	o.ContactName = contactName
}

// WithContactPhone adds the contactPhone to the api dcim sites list params
func (o *APIDcimSitesListParams) WithContactPhone(contactPhone *string) *APIDcimSitesListParams {
	o.SetContactPhone(contactPhone)
	return o
}

// SetContactPhone adds the contactPhone to the api dcim sites list params
func (o *APIDcimSitesListParams) SetContactPhone(contactPhone *string) {
	o.ContactPhone = contactPhone
}

// WithFacility adds the facility to the api dcim sites list params
func (o *APIDcimSitesListParams) WithFacility(facility *string) *APIDcimSitesListParams {
	o.SetFacility(facility)
	return o
}

// SetFacility adds the facility to the api dcim sites list params
func (o *APIDcimSitesListParams) SetFacility(facility *string) {
	o.Facility = facility
}

// WithIDIn adds the iDIn to the api dcim sites list params
func (o *APIDcimSitesListParams) WithIDIn(iDIn *string) *APIDcimSitesListParams {
	o.SetIDIn(iDIn)
	return o
}

// SetIDIn adds the idIn to the api dcim sites list params
func (o *APIDcimSitesListParams) SetIDIn(iDIn *string) {
	o.IDIn = iDIn
}

// WithLimit adds the limit to the api dcim sites list params
func (o *APIDcimSitesListParams) WithLimit(limit *int64) *APIDcimSitesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the api dcim sites list params
func (o *APIDcimSitesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the api dcim sites list params
func (o *APIDcimSitesListParams) WithName(name *string) *APIDcimSitesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the api dcim sites list params
func (o *APIDcimSitesListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the api dcim sites list params
func (o *APIDcimSitesListParams) WithOffset(offset *int64) *APIDcimSitesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the api dcim sites list params
func (o *APIDcimSitesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the api dcim sites list params
func (o *APIDcimSitesListParams) WithQ(q *string) *APIDcimSitesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the api dcim sites list params
func (o *APIDcimSitesListParams) SetQ(q *string) {
	o.Q = q
}

// WithRegion adds the region to the api dcim sites list params
func (o *APIDcimSitesListParams) WithRegion(region *string) *APIDcimSitesListParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the api dcim sites list params
func (o *APIDcimSitesListParams) SetRegion(region *string) {
	o.Region = region
}

// WithRegionID adds the regionID to the api dcim sites list params
func (o *APIDcimSitesListParams) WithRegionID(regionID *string) *APIDcimSitesListParams {
	o.SetRegionID(regionID)
	return o
}

// SetRegionID adds the regionId to the api dcim sites list params
func (o *APIDcimSitesListParams) SetRegionID(regionID *string) {
	o.RegionID = regionID
}

// WithSlug adds the slug to the api dcim sites list params
func (o *APIDcimSitesListParams) WithSlug(slug *string) *APIDcimSitesListParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the api dcim sites list params
func (o *APIDcimSitesListParams) SetSlug(slug *string) {
	o.Slug = slug
}

// WithStatus adds the status to the api dcim sites list params
func (o *APIDcimSitesListParams) WithStatus(status *string) *APIDcimSitesListParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the api dcim sites list params
func (o *APIDcimSitesListParams) SetStatus(status *string) {
	o.Status = status
}

// WithTag adds the tag to the api dcim sites list params
func (o *APIDcimSitesListParams) WithTag(tag *string) *APIDcimSitesListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the api dcim sites list params
func (o *APIDcimSitesListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTenant adds the tenant to the api dcim sites list params
func (o *APIDcimSitesListParams) WithTenant(tenant *string) *APIDcimSitesListParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the api dcim sites list params
func (o *APIDcimSitesListParams) SetTenant(tenant *string) {
	o.Tenant = tenant
}

// WithTenantID adds the tenantID to the api dcim sites list params
func (o *APIDcimSitesListParams) WithTenantID(tenantID *string) *APIDcimSitesListParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the api dcim sites list params
func (o *APIDcimSitesListParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *APIDcimSitesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Asn != nil {

		// query param asn
		var qrAsn float64
		if o.Asn != nil {
			qrAsn = *o.Asn
		}
		qAsn := swag.FormatFloat64(qrAsn)
		if qAsn != "" {
			if err := r.SetQueryParam("asn", qAsn); err != nil {
				return err
			}
		}

	}

	if o.ContactEmail != nil {

		// query param contact_email
		var qrContactEmail string
		if o.ContactEmail != nil {
			qrContactEmail = *o.ContactEmail
		}
		qContactEmail := qrContactEmail
		if qContactEmail != "" {
			if err := r.SetQueryParam("contact_email", qContactEmail); err != nil {
				return err
			}
		}

	}

	if o.ContactName != nil {

		// query param contact_name
		var qrContactName string
		if o.ContactName != nil {
			qrContactName = *o.ContactName
		}
		qContactName := qrContactName
		if qContactName != "" {
			if err := r.SetQueryParam("contact_name", qContactName); err != nil {
				return err
			}
		}

	}

	if o.ContactPhone != nil {

		// query param contact_phone
		var qrContactPhone string
		if o.ContactPhone != nil {
			qrContactPhone = *o.ContactPhone
		}
		qContactPhone := qrContactPhone
		if qContactPhone != "" {
			if err := r.SetQueryParam("contact_phone", qContactPhone); err != nil {
				return err
			}
		}

	}

	if o.Facility != nil {

		// query param facility
		var qrFacility string
		if o.Facility != nil {
			qrFacility = *o.Facility
		}
		qFacility := qrFacility
		if qFacility != "" {
			if err := r.SetQueryParam("facility", qFacility); err != nil {
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

	if o.Region != nil {

		// query param region
		var qrRegion string
		if o.Region != nil {
			qrRegion = *o.Region
		}
		qRegion := qrRegion
		if qRegion != "" {
			if err := r.SetQueryParam("region", qRegion); err != nil {
				return err
			}
		}

	}

	if o.RegionID != nil {

		// query param region_id
		var qrRegionID string
		if o.RegionID != nil {
			qrRegionID = *o.RegionID
		}
		qRegionID := qrRegionID
		if qRegionID != "" {
			if err := r.SetQueryParam("region_id", qRegionID); err != nil {
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

	if o.Status != nil {

		// query param status
		var qrStatus string
		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {
			if err := r.SetQueryParam("status", qStatus); err != nil {
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

	if o.Tenant != nil {

		// query param tenant
		var qrTenant string
		if o.Tenant != nil {
			qrTenant = *o.Tenant
		}
		qTenant := qrTenant
		if qTenant != "" {
			if err := r.SetQueryParam("tenant", qTenant); err != nil {
				return err
			}
		}

	}

	if o.TenantID != nil {

		// query param tenant_id
		var qrTenantID string
		if o.TenantID != nil {
			qrTenantID = *o.TenantID
		}
		qTenantID := qrTenantID
		if qTenantID != "" {
			if err := r.SetQueryParam("tenant_id", qTenantID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}