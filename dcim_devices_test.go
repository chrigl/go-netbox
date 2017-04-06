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
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func deviceEqual(a, b *Device) bool {
	var obsAB = []struct {
		a interface{}
		b interface{}
	}{
		{a: a.DeviceType, b: b.DeviceType},
		{a: a.DeviceRole, b: b.DeviceRole},
		{a: a.Platform, b: b.Platform},
		{a: a.Rack, b: b.Rack},
		{a: a.ParentDevice, b: b.ParentDevice},
		{a: a.ID, b: b.ID},
		{a: a.Name, b: b.Name},
		{a: a.DisplayName, b: b.DisplayName},
		{a: a.Serial, b: b.Serial},
		{a: a.Position, b: b.Position},
		{a: a.Face, b: b.Face},
		{a: a.Status, b: b.Status},
		{a: a.Comments, b: b.Comments},
		{a: a.PrimaryIP, b: b.PrimaryIP},
		{a: a.PrimaryIP4, b: b.PrimaryIP4},
		{a: a.PrimaryIP6, b: b.PrimaryIP6},
	}
	for _, o := range obsAB {

		switch o.a.(type) {
		case *IPAddressIdentifier:
			i, j := o.a.(*IPAddressIdentifier), o.b.(*IPAddressIdentifier)
			if ok := ipAddressIdentifiersEqual(*i, *j); !ok {
				return false
			}
		default:
			if !reflect.DeepEqual(o.a, o.b) {
				return false
			}
		}
	}

	return true
}

func devicesSlicesEqual(a, b []*Device) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if !deviceEqual(a[i], b[i]) {
			return false
		}
	}
	return true
}

func TestListDevicesOptionsValues(t *testing.T) {
	var tests = []struct {
		desc string
		o    *ListDevicesOptions
		v    url.Values
	}{
		{
			desc: "empty options",
		},
		{
			desc: "name only",
			o:    &ListDevicesOptions{Name: "Hello"},
			v: url.Values{
				"name": []string{"Hello"},
			},
		},
		{
			desc: "serial only",
			o:    &ListDevicesOptions{Serial: "my_serial"},
			v: url.Values{
				"serial": []string{"my_serial"},
			},
		},
		{
			desc: "site_id only",
			o:    &ListDevicesOptions{SiteID: 1},
			v: url.Values{
				"site_id": []string{"1"},
			},
		},
		{
			desc: "site only",
			o:    &ListDevicesOptions{Site: "my_site"},
			v: url.Values{
				"site": []string{"my_site"},
			},
		},
		{
			desc: "site_id and site",
			o:    &ListDevicesOptions{SiteID: 1, Site: "my_site"},
			v: url.Values{
				"site_id": []string{"1"},
			},
		},
		{
			desc: "1 rack_group_id",
			o:    &ListDevicesOptions{RackGroupID: []int{1}},
			v: url.Values{
				"rack_group_id": []string{"1"},
			},
		},
		{
			desc: "3 rack_group_id",
			o:    &ListDevicesOptions{RackGroupID: []int{1, 2, 3}},
			v: url.Values{
				"rack_group_id": []string{"1", "2", "3"},
			},
		},
		{
			desc: "1 rack_id",
			o:    &ListDevicesOptions{RackID: []int{1}},
			v: url.Values{
				"rack_id": []string{"1"},
			},
		},
		{
			desc: "3 rack_id",
			o:    &ListDevicesOptions{RackID: []int{1, 2, 3}},
			v: url.Values{
				"rack_id": []string{"1", "2", "3"},
			},
		},
		{
			desc: "1 role_id",
			o:    &ListDevicesOptions{RoleID: []int{1}},
			v: url.Values{
				"role_id": []string{"1"},
			},
		},
		{
			desc: "3 role_id",
			o:    &ListDevicesOptions{RoleID: []int{1, 2, 3}},
			v: url.Values{
				"role_id": []string{"1", "2", "3"},
			},
		},
		{
			desc: "3 role",
			o:    &ListDevicesOptions{Role: []string{"role1", "role2", "role3"}},
			v: url.Values{
				"role": []string{"role1", "role2", "role3"},
			},
		},
		{
			desc: "role_id and role",
			o:    &ListDevicesOptions{RoleID: []int{1}, Role: []string{"my_role"}},
			v: url.Values{
				"role_id": []string{"1"},
			},
		},
		{
			desc: "1 tenant_id",
			o:    &ListDevicesOptions{TenantID: []int{1}},
			v: url.Values{
				"tenant_id": []string{"1"},
			},
		},
		{
			desc: "3 tenant_id",
			o:    &ListDevicesOptions{TenantID: []int{1, 2, 3}},
			v: url.Values{
				"tenant_id": []string{"1", "2", "3"},
			},
		},
		{
			desc: "3 tenant",
			o:    &ListDevicesOptions{Tenant: []string{"tenant1", "tenant2", "tenant3"}},
			v: url.Values{
				"tenant": []string{"tenant1", "tenant2", "tenant3"},
			},
		},
		{
			desc: "tenant_id and tenant",
			o:    &ListDevicesOptions{TenantID: []int{1}, Tenant: []string{"my_tenant"}},
			v: url.Values{
				"tenant_id": []string{"1"},
			},
		},
		{
			desc: "1 manufacturer_id",
			o:    &ListDevicesOptions{ManufacturerID: []int{1}},
			v: url.Values{
				"manufacturer_id": []string{"1"},
			},
		},
		{
			desc: "3 manufacturer_id",
			o:    &ListDevicesOptions{ManufacturerID: []int{1, 2, 3}},
			v: url.Values{
				"manufacturer_id": []string{"1", "2", "3"},
			},
		},
		{
			desc: "3 manufacturer",
			o:    &ListDevicesOptions{Manufacturer: []string{"manufacturer1", "manufacturer2", "manufacturer3"}},
			v: url.Values{
				"manufacturer": []string{"manufacturer1", "manufacturer2", "manufacturer3"},
			},
		},
		{
			desc: "manufacturer_id and manufacturer",
			o:    &ListDevicesOptions{ManufacturerID: []int{1}, Manufacturer: []string{"my_manufacturer"}},
			v: url.Values{
				"manufacturer_id": []string{"1"},
			},
		},
		{
			desc: "1 device_type_id",
			o:    &ListDevicesOptions{DeviceTypeID: []int{1}},
			v: url.Values{
				"device_type_id": []string{"1"},
			},
		},
		{
			desc: "3 device_type_id",
			o:    &ListDevicesOptions{DeviceTypeID: []int{1, 2, 3}},
			v: url.Values{
				"device_type_id": []string{"1", "2", "3"},
			},
		},
	}

	for i, tt := range tests {
		t.Logf("[%02d] test %q", i, tt.desc)

		v, err := tt.o.values()
		if err != nil {
			t.Fatalf("unexpected Values error: %v", err)
		}

		if want, got := tt.v, v; !reflect.DeepEqual(want, got) {
			t.Fatalf("unexpected url.Values map:\n- want: %v\n-  got: %v",
				want, got)
		}
	}
}

func TestClientDCIMGetDevice(t *testing.T) {
	wantDevice := testDevice(1)

	c, done := testClient(t, testHandler(t, http.MethodGet, "/api/dcim/devices/1", wantDevice))
	defer done()

	gotDevice, err := c.DCIM.GetDevice(wantDevice.ID)
	if err != nil {
		t.Fatalf("unexpected error from Client.DCIM.GetDevice: %v", err)
	}

	if !deviceEqual(wantDevice, gotDevice) {
		t.Fatalf("unexpected Device:\n- want: %v\n-  got: %v", wantDevice, gotDevice)
	}
}

func TestClientDCIMListDevices(t *testing.T) {
	want := []*Device{
		testDevice(1),
		testDevice(2),
	}

	c, done := testClient(t, testHandler(t, http.MethodGet, "/api/dcim/devices/", want))
	defer done()

	got, err := c.DCIM.ListDevices(nil)
	if err != nil {
		t.Fatalf("unexpected error from Client.DCIM.ListDevices: %v", err)
	}

	if !devicesSlicesEqual(want, got) {
		t.Fatalf("unexpected Devices:\n- want: %v\n-  got: %v", want, got)
	}
}
