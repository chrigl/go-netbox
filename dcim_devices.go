// Copyright 2016 The go-netbox Authors.
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

package netbox

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// Device is a network device.
type Device struct {
	ID           int                   `json:"id"`
	Name         string                `json:"name"`
	DisplayName  string                `json:"display_name"`
	DeviceType   *DeviceTypeIdentifier `json:"device_type"`
	DeviceRole   *SimpleIdentifier     `json:"device_role"`
	Platform     *SimpleIdentifier     `json:"platform"`
	Serial       string                `json:"serial"`
	Rack         *RackIdentifier       `json:"rack"`
	Position     int                   `json:"position"`
	Face         int                   `json:"face"`
	ParentDevice *DeviceIdentifier     `json:"parent_device"`
	Status       bool                  `json:"status"`
	PrimaryIP    *IPAddressIdentifier  `json:"primary_ip"`
	PrimaryIP4   *IPAddressIdentifier  `json:"primary_ip4"`
	PrimaryIP6   *IPAddressIdentifier  `json:"primary_ip6"`
	Comments     string                `json:"comments"`
}

// A DeviceIdentifier is a reduced version of a Device, returned as a nested
// object in some top-level objects.  It contains information which can
// be used in subsequent API calls to identify and retrieve a full Device.
type DeviceIdentifier struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// A DeviceTypeIdentifier indicates the device type of a network device.
type DeviceTypeIdentifier struct {
	ID           int               `json:"id"`
	Manufacturer *SimpleIdentifier `json:"manufacturer"`
	Model        string            `json:"model"`
	Slug         string            `json:"slug"`
}

// RackIdentifier represents a server rack.
type RackIdentifier struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	FacilityID  string `json:"facility_id"`
	DisplayName string `json:"display_name"`
}

func (s *DCIMService) GetDevice(id int) (*Device, error) {
	req, err := s.c.NewRequest(
		http.MethodGet,
		fmt.Sprintf("/api/dcim/devices/%d", id),
		nil,
	)
	if err != nil {
		return nil, err
	}

	device := new(Device)
	err = s.c.Do(req, device)
	return device, err
}

func (s *DCIMService) ListDevices(options *ListDevicesOptions) ([]*Device, error) {
	req, err := s.c.NewRequest(http.MethodGet, "/api/dcim/devices/", options)
	if err != nil {
		return nil, err
	}

	var devices []*Device
	err = s.c.Do(req, &devices)
	return devices, err
}

// ListDevicesOptions is used as an argument for Client.DCIM.ListDevices.
// Integer fields with an *ID suffix are preferred over their string
// counterparts, and if both are set, only the *ID field will be used.
type ListDevicesOptions struct {
	Name   string
	Serial string

	SiteID         int
	Site           string
	RackGroupID    []int
	RackID         []int
	RoleID         []int
	Role           []string
	TenantID       []int
	Tenant         []string
	DeviceTypeID   []int
	ManufacturerID []int
	Manufacturer   []string

	// Query is a special option which enables free-form search.
	// For example, Query could be an IP address such as "8.8.8.8".
	Query string
}

func (o *ListDevicesOptions) values() (url.Values, error) {
	if o == nil {
		return nil, nil
	}

	v := url.Values{}

	if o.Name != "" {
		v.Set("name", o.Name)
	}
	if o.Serial != "" {
		v.Set("serial", o.Serial)
	}

	switch {
	case o.SiteID > 0:
		v.Add("site_id", strconv.Itoa(o.SiteID))
	case o.Site != "":
		v.Add("site", o.Site)
	}

	for _, r := range o.RackGroupID {
		v.Add("rack_group_id", strconv.Itoa(r))
	}

	for _, r := range o.RackID {
		v.Add("rack_id", strconv.Itoa(r))
	}

	switch {
	case len(o.RoleID) > 0:
		for _, r := range o.RoleID {
			v.Add("role_id", strconv.Itoa(r))
		}
	case len(o.Role) > 0:
		for _, r := range o.Role {
			v.Add("role", r)
		}
	}

	switch {
	case len(o.TenantID) > 0:
		for _, t := range o.TenantID {
			v.Add("tenant_id", strconv.Itoa(t))
		}
	case len(o.Tenant) > 0:
		for _, t := range o.Tenant {
			v.Add("tenant", t)
		}
	}

	for _, d := range o.DeviceTypeID {
		v.Add("device_type_id", strconv.Itoa(d))
	}

	switch {
	case len(o.ManufacturerID) > 0:
		for _, m := range o.ManufacturerID {
			v.Add("manufacturer_id", strconv.Itoa(m))
		}
	case len(o.Manufacturer) > 0:
		for _, m := range o.Manufacturer {
			v.Add("manufacturer", m)
		}
	}

	if o.Query != "" {
		v.Set("q", o.Query)
	}

	return v, nil
}
