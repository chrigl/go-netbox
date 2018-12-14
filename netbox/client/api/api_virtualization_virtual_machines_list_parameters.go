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

// NewAPIVirtualizationVirtualMachinesListParams creates a new APIVirtualizationVirtualMachinesListParams object
// with the default values initialized.
func NewAPIVirtualizationVirtualMachinesListParams() *APIVirtualizationVirtualMachinesListParams {
	var ()
	return &APIVirtualizationVirtualMachinesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIVirtualizationVirtualMachinesListParamsWithTimeout creates a new APIVirtualizationVirtualMachinesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIVirtualizationVirtualMachinesListParamsWithTimeout(timeout time.Duration) *APIVirtualizationVirtualMachinesListParams {
	var ()
	return &APIVirtualizationVirtualMachinesListParams{

		timeout: timeout,
	}
}

// NewAPIVirtualizationVirtualMachinesListParamsWithContext creates a new APIVirtualizationVirtualMachinesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIVirtualizationVirtualMachinesListParamsWithContext(ctx context.Context) *APIVirtualizationVirtualMachinesListParams {
	var ()
	return &APIVirtualizationVirtualMachinesListParams{

		Context: ctx,
	}
}

// NewAPIVirtualizationVirtualMachinesListParamsWithHTTPClient creates a new APIVirtualizationVirtualMachinesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIVirtualizationVirtualMachinesListParamsWithHTTPClient(client *http.Client) *APIVirtualizationVirtualMachinesListParams {
	var ()
	return &APIVirtualizationVirtualMachinesListParams{
		HTTPClient: client,
	}
}

/*APIVirtualizationVirtualMachinesListParams contains all the parameters to send to the API endpoint
for the api virtualization virtual machines list operation typically these are written to a http.Request
*/
type APIVirtualizationVirtualMachinesListParams struct {

	/*Cluster*/
	Cluster *string
	/*ClusterGroup*/
	ClusterGroup *string
	/*ClusterGroupID*/
	ClusterGroupID *string
	/*ClusterID*/
	ClusterID *string
	/*ClusterType*/
	ClusterType *string
	/*ClusterTypeID*/
	ClusterTypeID *string
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
	/*Platform*/
	Platform *string
	/*PlatformID*/
	PlatformID *string
	/*Q*/
	Q *string
	/*Region*/
	Region *string
	/*RegionID*/
	RegionID *float64
	/*Role*/
	Role *string
	/*RoleID*/
	RoleID *string
	/*Site*/
	Site *string
	/*SiteID*/
	SiteID *string
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

// WithTimeout adds the timeout to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) WithTimeout(timeout time.Duration) *APIVirtualizationVirtualMachinesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) WithContext(ctx context.Context) *APIVirtualizationVirtualMachinesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) WithHTTPClient(client *http.Client) *APIVirtualizationVirtualMachinesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCluster adds the cluster to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) WithCluster(cluster *string) *APIVirtualizationVirtualMachinesListParams {
	o.SetCluster(cluster)
	return o
}

// SetCluster adds the cluster to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) SetCluster(cluster *string) {
	o.Cluster = cluster
}

// WithClusterGroup adds the clusterGroup to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) WithClusterGroup(clusterGroup *string) *APIVirtualizationVirtualMachinesListParams {
	o.SetClusterGroup(clusterGroup)
	return o
}

// SetClusterGroup adds the clusterGroup to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) SetClusterGroup(clusterGroup *string) {
	o.ClusterGroup = clusterGroup
}

// WithClusterGroupID adds the clusterGroupID to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) WithClusterGroupID(clusterGroupID *string) *APIVirtualizationVirtualMachinesListParams {
	o.SetClusterGroupID(clusterGroupID)
	return o
}

// SetClusterGroupID adds the clusterGroupId to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) SetClusterGroupID(clusterGroupID *string) {
	o.ClusterGroupID = clusterGroupID
}

// WithClusterID adds the clusterID to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) WithClusterID(clusterID *string) *APIVirtualizationVirtualMachinesListParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) SetClusterID(clusterID *string) {
	o.ClusterID = clusterID
}

// WithClusterType adds the clusterType to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) WithClusterType(clusterType *string) *APIVirtualizationVirtualMachinesListParams {
	o.SetClusterType(clusterType)
	return o
}

// SetClusterType adds the clusterType to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) SetClusterType(clusterType *string) {
	o.ClusterType = clusterType
}

// WithClusterTypeID adds the clusterTypeID to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) WithClusterTypeID(clusterTypeID *string) *APIVirtualizationVirtualMachinesListParams {
	o.SetClusterTypeID(clusterTypeID)
	return o
}

// SetClusterTypeID adds the clusterTypeId to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) SetClusterTypeID(clusterTypeID *string) {
	o.ClusterTypeID = clusterTypeID
}

// WithIDIn adds the iDIn to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) WithIDIn(iDIn *string) *APIVirtualizationVirtualMachinesListParams {
	o.SetIDIn(iDIn)
	return o
}

// SetIDIn adds the idIn to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) SetIDIn(iDIn *string) {
	o.IDIn = iDIn
}

// WithLimit adds the limit to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) WithLimit(limit *int64) *APIVirtualizationVirtualMachinesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) WithName(name *string) *APIVirtualizationVirtualMachinesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) WithOffset(offset *int64) *APIVirtualizationVirtualMachinesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPlatform adds the platform to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) WithPlatform(platform *string) *APIVirtualizationVirtualMachinesListParams {
	o.SetPlatform(platform)
	return o
}

// SetPlatform adds the platform to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) SetPlatform(platform *string) {
	o.Platform = platform
}

// WithPlatformID adds the platformID to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) WithPlatformID(platformID *string) *APIVirtualizationVirtualMachinesListParams {
	o.SetPlatformID(platformID)
	return o
}

// SetPlatformID adds the platformId to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) SetPlatformID(platformID *string) {
	o.PlatformID = platformID
}

// WithQ adds the q to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) WithQ(q *string) *APIVirtualizationVirtualMachinesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) SetQ(q *string) {
	o.Q = q
}

// WithRegion adds the region to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) WithRegion(region *string) *APIVirtualizationVirtualMachinesListParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) SetRegion(region *string) {
	o.Region = region
}

// WithRegionID adds the regionID to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) WithRegionID(regionID *float64) *APIVirtualizationVirtualMachinesListParams {
	o.SetRegionID(regionID)
	return o
}

// SetRegionID adds the regionId to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) SetRegionID(regionID *float64) {
	o.RegionID = regionID
}

// WithRole adds the role to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) WithRole(role *string) *APIVirtualizationVirtualMachinesListParams {
	o.SetRole(role)
	return o
}

// SetRole adds the role to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) SetRole(role *string) {
	o.Role = role
}

// WithRoleID adds the roleID to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) WithRoleID(roleID *string) *APIVirtualizationVirtualMachinesListParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) SetRoleID(roleID *string) {
	o.RoleID = roleID
}

// WithSite adds the site to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) WithSite(site *string) *APIVirtualizationVirtualMachinesListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) SetSite(site *string) {
	o.Site = site
}

// WithSiteID adds the siteID to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) WithSiteID(siteID *string) *APIVirtualizationVirtualMachinesListParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithStatus adds the status to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) WithStatus(status *string) *APIVirtualizationVirtualMachinesListParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) SetStatus(status *string) {
	o.Status = status
}

// WithTag adds the tag to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) WithTag(tag *string) *APIVirtualizationVirtualMachinesListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTenant adds the tenant to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) WithTenant(tenant *string) *APIVirtualizationVirtualMachinesListParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) SetTenant(tenant *string) {
	o.Tenant = tenant
}

// WithTenantID adds the tenantID to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) WithTenantID(tenantID *string) *APIVirtualizationVirtualMachinesListParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the api virtualization virtual machines list params
func (o *APIVirtualizationVirtualMachinesListParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *APIVirtualizationVirtualMachinesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cluster != nil {

		// query param cluster
		var qrCluster string
		if o.Cluster != nil {
			qrCluster = *o.Cluster
		}
		qCluster := qrCluster
		if qCluster != "" {
			if err := r.SetQueryParam("cluster", qCluster); err != nil {
				return err
			}
		}

	}

	if o.ClusterGroup != nil {

		// query param cluster_group
		var qrClusterGroup string
		if o.ClusterGroup != nil {
			qrClusterGroup = *o.ClusterGroup
		}
		qClusterGroup := qrClusterGroup
		if qClusterGroup != "" {
			if err := r.SetQueryParam("cluster_group", qClusterGroup); err != nil {
				return err
			}
		}

	}

	if o.ClusterGroupID != nil {

		// query param cluster_group_id
		var qrClusterGroupID string
		if o.ClusterGroupID != nil {
			qrClusterGroupID = *o.ClusterGroupID
		}
		qClusterGroupID := qrClusterGroupID
		if qClusterGroupID != "" {
			if err := r.SetQueryParam("cluster_group_id", qClusterGroupID); err != nil {
				return err
			}
		}

	}

	if o.ClusterID != nil {

		// query param cluster_id
		var qrClusterID string
		if o.ClusterID != nil {
			qrClusterID = *o.ClusterID
		}
		qClusterID := qrClusterID
		if qClusterID != "" {
			if err := r.SetQueryParam("cluster_id", qClusterID); err != nil {
				return err
			}
		}

	}

	if o.ClusterType != nil {

		// query param cluster_type
		var qrClusterType string
		if o.ClusterType != nil {
			qrClusterType = *o.ClusterType
		}
		qClusterType := qrClusterType
		if qClusterType != "" {
			if err := r.SetQueryParam("cluster_type", qClusterType); err != nil {
				return err
			}
		}

	}

	if o.ClusterTypeID != nil {

		// query param cluster_type_id
		var qrClusterTypeID string
		if o.ClusterTypeID != nil {
			qrClusterTypeID = *o.ClusterTypeID
		}
		qClusterTypeID := qrClusterTypeID
		if qClusterTypeID != "" {
			if err := r.SetQueryParam("cluster_type_id", qClusterTypeID); err != nil {
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
		var qrRegionID float64
		if o.RegionID != nil {
			qrRegionID = *o.RegionID
		}
		qRegionID := swag.FormatFloat64(qrRegionID)
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
