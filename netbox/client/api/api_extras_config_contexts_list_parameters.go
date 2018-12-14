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

// NewAPIExtrasConfigContextsListParams creates a new APIExtrasConfigContextsListParams object
// with the default values initialized.
func NewAPIExtrasConfigContextsListParams() *APIExtrasConfigContextsListParams {
	var ()
	return &APIExtrasConfigContextsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIExtrasConfigContextsListParamsWithTimeout creates a new APIExtrasConfigContextsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIExtrasConfigContextsListParamsWithTimeout(timeout time.Duration) *APIExtrasConfigContextsListParams {
	var ()
	return &APIExtrasConfigContextsListParams{

		timeout: timeout,
	}
}

// NewAPIExtrasConfigContextsListParamsWithContext creates a new APIExtrasConfigContextsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIExtrasConfigContextsListParamsWithContext(ctx context.Context) *APIExtrasConfigContextsListParams {
	var ()
	return &APIExtrasConfigContextsListParams{

		Context: ctx,
	}
}

// NewAPIExtrasConfigContextsListParamsWithHTTPClient creates a new APIExtrasConfigContextsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIExtrasConfigContextsListParamsWithHTTPClient(client *http.Client) *APIExtrasConfigContextsListParams {
	var ()
	return &APIExtrasConfigContextsListParams{
		HTTPClient: client,
	}
}

/*APIExtrasConfigContextsListParams contains all the parameters to send to the API endpoint
for the api extras config contexts list operation typically these are written to a http.Request
*/
type APIExtrasConfigContextsListParams struct {

	/*IsActive*/
	IsActive *string
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
	/*Platform*/
	Platform *string
	/*PlatformID*/
	PlatformID *string
	/*Q*/
	Q *string
	/*Region*/
	Region *string
	/*RegionID*/
	RegionID *string
	/*Role*/
	Role *string
	/*RoleID*/
	RoleID *string
	/*Site*/
	Site *string
	/*SiteID*/
	SiteID *string
	/*Tenant*/
	Tenant *string
	/*TenantGroup*/
	TenantGroup *string
	/*TenantGroupID*/
	TenantGroupID *string
	/*TenantID*/
	TenantID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the api extras config contexts list params
func (o *APIExtrasConfigContextsListParams) WithTimeout(timeout time.Duration) *APIExtrasConfigContextsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the api extras config contexts list params
func (o *APIExtrasConfigContextsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the api extras config contexts list params
func (o *APIExtrasConfigContextsListParams) WithContext(ctx context.Context) *APIExtrasConfigContextsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the api extras config contexts list params
func (o *APIExtrasConfigContextsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the api extras config contexts list params
func (o *APIExtrasConfigContextsListParams) WithHTTPClient(client *http.Client) *APIExtrasConfigContextsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the api extras config contexts list params
func (o *APIExtrasConfigContextsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIsActive adds the isActive to the api extras config contexts list params
func (o *APIExtrasConfigContextsListParams) WithIsActive(isActive *string) *APIExtrasConfigContextsListParams {
	o.SetIsActive(isActive)
	return o
}

// SetIsActive adds the isActive to the api extras config contexts list params
func (o *APIExtrasConfigContextsListParams) SetIsActive(isActive *string) {
	o.IsActive = isActive
}

// WithLimit adds the limit to the api extras config contexts list params
func (o *APIExtrasConfigContextsListParams) WithLimit(limit *int64) *APIExtrasConfigContextsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the api extras config contexts list params
func (o *APIExtrasConfigContextsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the api extras config contexts list params
func (o *APIExtrasConfigContextsListParams) WithName(name *string) *APIExtrasConfigContextsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the api extras config contexts list params
func (o *APIExtrasConfigContextsListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the api extras config contexts list params
func (o *APIExtrasConfigContextsListParams) WithOffset(offset *int64) *APIExtrasConfigContextsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the api extras config contexts list params
func (o *APIExtrasConfigContextsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPlatform adds the platform to the api extras config contexts list params
func (o *APIExtrasConfigContextsListParams) WithPlatform(platform *string) *APIExtrasConfigContextsListParams {
	o.SetPlatform(platform)
	return o
}

// SetPlatform adds the platform to the api extras config contexts list params
func (o *APIExtrasConfigContextsListParams) SetPlatform(platform *string) {
	o.Platform = platform
}

// WithPlatformID adds the platformID to the api extras config contexts list params
func (o *APIExtrasConfigContextsListParams) WithPlatformID(platformID *string) *APIExtrasConfigContextsListParams {
	o.SetPlatformID(platformID)
	return o
}

// SetPlatformID adds the platformId to the api extras config contexts list params
func (o *APIExtrasConfigContextsListParams) SetPlatformID(platformID *string) {
	o.PlatformID = platformID
}

// WithQ adds the q to the api extras config contexts list params
func (o *APIExtrasConfigContextsListParams) WithQ(q *string) *APIExtrasConfigContextsListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the api extras config contexts list params
func (o *APIExtrasConfigContextsListParams) SetQ(q *string) {
	o.Q = q
}

// WithRegion adds the region to the api extras config contexts list params
func (o *APIExtrasConfigContextsListParams) WithRegion(region *string) *APIExtrasConfigContextsListParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the api extras config contexts list params
func (o *APIExtrasConfigContextsListParams) SetRegion(region *string) {
	o.Region = region
}

// WithRegionID adds the regionID to the api extras config contexts list params
func (o *APIExtrasConfigContextsListParams) WithRegionID(regionID *string) *APIExtrasConfigContextsListParams {
	o.SetRegionID(regionID)
	return o
}

// SetRegionID adds the regionId to the api extras config contexts list params
func (o *APIExtrasConfigContextsListParams) SetRegionID(regionID *string) {
	o.RegionID = regionID
}

// WithRole adds the role to the api extras config contexts list params
func (o *APIExtrasConfigContextsListParams) WithRole(role *string) *APIExtrasConfigContextsListParams {
	o.SetRole(role)
	return o
}

// SetRole adds the role to the api extras config contexts list params
func (o *APIExtrasConfigContextsListParams) SetRole(role *string) {
	o.Role = role
}

// WithRoleID adds the roleID to the api extras config contexts list params
func (o *APIExtrasConfigContextsListParams) WithRoleID(roleID *string) *APIExtrasConfigContextsListParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the api extras config contexts list params
func (o *APIExtrasConfigContextsListParams) SetRoleID(roleID *string) {
	o.RoleID = roleID
}

// WithSite adds the site to the api extras config contexts list params
func (o *APIExtrasConfigContextsListParams) WithSite(site *string) *APIExtrasConfigContextsListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the api extras config contexts list params
func (o *APIExtrasConfigContextsListParams) SetSite(site *string) {
	o.Site = site
}

// WithSiteID adds the siteID to the api extras config contexts list params
func (o *APIExtrasConfigContextsListParams) WithSiteID(siteID *string) *APIExtrasConfigContextsListParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the api extras config contexts list params
func (o *APIExtrasConfigContextsListParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithTenant adds the tenant to the api extras config contexts list params
func (o *APIExtrasConfigContextsListParams) WithTenant(tenant *string) *APIExtrasConfigContextsListParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the api extras config contexts list params
func (o *APIExtrasConfigContextsListParams) SetTenant(tenant *string) {
	o.Tenant = tenant
}

// WithTenantGroup adds the tenantGroup to the api extras config contexts list params
func (o *APIExtrasConfigContextsListParams) WithTenantGroup(tenantGroup *string) *APIExtrasConfigContextsListParams {
	o.SetTenantGroup(tenantGroup)
	return o
}

// SetTenantGroup adds the tenantGroup to the api extras config contexts list params
func (o *APIExtrasConfigContextsListParams) SetTenantGroup(tenantGroup *string) {
	o.TenantGroup = tenantGroup
}

// WithTenantGroupID adds the tenantGroupID to the api extras config contexts list params
func (o *APIExtrasConfigContextsListParams) WithTenantGroupID(tenantGroupID *string) *APIExtrasConfigContextsListParams {
	o.SetTenantGroupID(tenantGroupID)
	return o
}

// SetTenantGroupID adds the tenantGroupId to the api extras config contexts list params
func (o *APIExtrasConfigContextsListParams) SetTenantGroupID(tenantGroupID *string) {
	o.TenantGroupID = tenantGroupID
}

// WithTenantID adds the tenantID to the api extras config contexts list params
func (o *APIExtrasConfigContextsListParams) WithTenantID(tenantID *string) *APIExtrasConfigContextsListParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the api extras config contexts list params
func (o *APIExtrasConfigContextsListParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *APIExtrasConfigContextsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IsActive != nil {

		// query param is_active
		var qrIsActive string
		if o.IsActive != nil {
			qrIsActive = *o.IsActive
		}
		qIsActive := qrIsActive
		if qIsActive != "" {
			if err := r.SetQueryParam("is_active", qIsActive); err != nil {
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

	if o.Platform != nil {

		// query param platform
		var qrPlatform string
		if o.Platform != nil {
			qrPlatform = *o.Platform
		}
		qPlatform := qrPlatform
		if qPlatform != "" {
			if err := r.SetQueryParam("platform", qPlatform); err != nil {
				return err
			}
		}

	}

	if o.PlatformID != nil {

		// query param platform_id
		var qrPlatformID string
		if o.PlatformID != nil {
			qrPlatformID = *o.PlatformID
		}
		qPlatformID := qrPlatformID
		if qPlatformID != "" {
			if err := r.SetQueryParam("platform_id", qPlatformID); err != nil {
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

	if o.Role != nil {

		// query param role
		var qrRole string
		if o.Role != nil {
			qrRole = *o.Role
		}
		qRole := qrRole
		if qRole != "" {
			if err := r.SetQueryParam("role", qRole); err != nil {
				return err
			}
		}

	}

	if o.RoleID != nil {

		// query param role_id
		var qrRoleID string
		if o.RoleID != nil {
			qrRoleID = *o.RoleID
		}
		qRoleID := qrRoleID
		if qRoleID != "" {
			if err := r.SetQueryParam("role_id", qRoleID); err != nil {
				return err
			}
		}

	}

	if o.Site != nil {

		// query param site
		var qrSite string
		if o.Site != nil {
			qrSite = *o.Site
		}
		qSite := qrSite
		if qSite != "" {
			if err := r.SetQueryParam("site", qSite); err != nil {
				return err
			}
		}

	}

	if o.SiteID != nil {

		// query param site_id
		var qrSiteID string
		if o.SiteID != nil {
			qrSiteID = *o.SiteID
		}
		qSiteID := qrSiteID
		if qSiteID != "" {
			if err := r.SetQueryParam("site_id", qSiteID); err != nil {
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

	if o.TenantGroup != nil {

		// query param tenant_group
		var qrTenantGroup string
		if o.TenantGroup != nil {
			qrTenantGroup = *o.TenantGroup
		}
		qTenantGroup := qrTenantGroup
		if qTenantGroup != "" {
			if err := r.SetQueryParam("tenant_group", qTenantGroup); err != nil {
				return err
			}
		}

	}

	if o.TenantGroupID != nil {

		// query param tenant_group_id
		var qrTenantGroupID string
		if o.TenantGroupID != nil {
			qrTenantGroupID = *o.TenantGroupID
		}
		qTenantGroupID := qrTenantGroupID
		if qTenantGroupID != "" {
			if err := r.SetQueryParam("tenant_group_id", qTenantGroupID); err != nil {
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
