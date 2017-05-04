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
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
	"time"
)

func TestClientBadJSON(t *testing.T) {
	c, done := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("foo"))
	})
	defer done()

	req, err := c.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal("expected an error, but no error returned")
	}

	// Pass empty struct to trigger JSON unmarshaling path
	var v struct{}

	err = c.Do(req, &v)
	if _, ok := err.(*json.SyntaxError); !ok {
		t.Fatalf("unexpected error type: %T", err)
	}
}

func TestClientQueryParameters(t *testing.T) {
	c := &Client{
		u:      &url.URL{},
		client: &http.Client{},
	}

	const (
		wantFoo = "foo"
		wantBar = 1
	)

	req, err := c.NewRequest(http.MethodGet, "/", testValuer{
		Foo: wantFoo,
		Bar: wantBar,
	})
	if err != nil {
		t.Fatal("expected an error, but no error returned")
	}

	q := req.URL.Query()
	if want, got := 2, len(q); want != got {
		t.Fatalf("unexpected number of query parameters:\n- want: %v\n-  got: %v",
			want, got)
	}

	if want, got := wantFoo, q.Get("foo"); want != got {
		t.Fatalf("unexpected foo:\n- want: %v\n-  got: %v", want, got)
	}

	if want, got := strconv.Itoa(wantBar), q.Get("bar"); want != got {
		t.Fatalf("unexpected bar:\n- want: %v\n-  got: %v", want, got)
	}
}

func TestClientPrependBaseURLPath(t *testing.T) {
	u, err := url.Parse("http://example.com/netbox/")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	c := &Client{
		u:      u,
		client: &http.Client{},
	}

	req, err := c.NewRequest(http.MethodGet, "/api/ipam/vlans", nil)
	if err != nil {
		t.Fatal("expected an error, but no error returned")
	}

	if want, got := "/netbox/api/ipam/vlans", req.URL.Path; want != got {
		t.Fatalf("unexpected URL path:\n- want: %q\n-  got: %q",
			want, got)
	}
}

type testValuer struct {
	Foo string
	Bar int
}

func (q testValuer) values() (url.Values, error) {
	v := url.Values{}

	if q.Foo != "" {
		v.Set("foo", q.Foo)
	}

	if q.Bar != 0 {
		v.Set("bar", strconv.Itoa(q.Bar))
	}

	return v, nil
}

func testClient(t *testing.T, fn func(w http.ResponseWriter, r *http.Request)) (*Client, func()) {
	s := httptest.NewServer(http.HandlerFunc(fn))

	c, err := NewClient(s.URL, nil)
	if err != nil {
		t.Fatalf("error creating Client: %v", err)
	}

	return c, func() { s.Close() }
}

// Just to be able to override MarshalJSON.
// For testing we need this impl e.g. in testHandler.
// If we Marshal a real IPAddress, it needs to be much smaller
// to be PUT, PATCHED against netbox-api v2
type mockIPAddress IPAddress

func getMockIPAddress(ip *IPAddress) *mockIPAddress {
	var newIP mockIPAddress = mockIPAddress(*ip)
	return &newIP
}

func getMockIPAddresses(ips []*IPAddress) []*mockIPAddress {
	res := make([]*mockIPAddress, len(ips))
	for i, ip := range ips {
		res[i] = getMockIPAddress(ip)
	}
	return res
}

func (ip *mockIPAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(ipAddress{
		ID:          ip.ID,
		Family:      ip.Family,
		Address:     ip.Address.String(),
		VRF:         ip.VRF,
		Tenant:      ip.Tenant,
		Interface:   ip.Interface,
		Description: ip.Description,
		NATInside:   ip.NATInside,
		NATOutside:  ip.NATOutside,
		Status:      ip.Status,
	})
}

func testHandler(t *testing.T, method string, path string, v interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if want, got := method, r.Method; want != got {
			t.Fatalf("unexpected HTTP method:\n- want: %v\n-  got: %v", want, got)
		}

		if want, got := path, r.URL.Path; want != got {
			t.Fatalf("unexpected URL path:\n- want: %v\n-  got: %v", want, got)
		}

		if err := json.NewEncoder(w).Encode(v); err != nil {
			t.Fatalf("error while encoding JSON: %v", err)
		}
	}
}

// Test helpers for generating mock data

func testAggregate(family Family, n int) *Aggregate {
	prefix := &net.IPNet{
		IP:   net.IPv4(8, 0, 0, 0),
		Mask: net.CIDRMask(8, 32),
	}
	if family == FamilyIPv6 {
		prefix = &net.IPNet{
			IP:   net.ParseIP("2001::"),
			Mask: net.CIDRMask(16, 128),
		}
	}

	return &Aggregate{
		ID:          n,
		Family:      family,
		Prefix:      prefix,
		RIR:         testRIRIdentifier(n),
		DateAdded:   time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC),
		Description: fmt.Sprintf("description %d", n),
	}
}

func testConsolePort(n int) *ConsolePort {
	return &ConsolePort{
		ID:               n,
		Device:           testDeviceIdentifier(n),
		Name:             fmt.Sprintf("deviceport %d", n),
		CSPort:           testConsoleServerPort(n),
		ConnectionStatus: true,
	}
}

func testConsolePortIdentifier(n int) *ConsolePortIdentifier {
	return &ConsolePortIdentifier{
		Device: fmt.Sprintf("device %d", n),
		Name:   fmt.Sprintf("rc console port %d", n),
		Port:   fmt.Sprintf("port %d", n),
	}
}

func testConsoleServerPort(n int) *ConsoleServerPort {
	return &ConsoleServerPort{
		ID:     n,
		Device: testDeviceIdentifier(n),
		Name:   fmt.Sprintf("consoleserverport %d", n),
	}
}

func testDevice(n int) *Device {
	return &Device{
		ID:           n,
		Name:         fmt.Sprintf("device %d", n),
		DisplayName:  fmt.Sprintf("Device %d", n),
		DeviceType:   testDeviceTypeIdentifier(n),
		DeviceRole:   testSimpleIdentifier(n),
		Platform:     testSimpleIdentifier(n),
		Serial:       fmt.Sprintf("relatedconnection%d", n),
		Rack:         testRackIdentifier(n),
		Position:     n,
		Face:         n,
		ParentDevice: testDeviceIdentifier(n),
		Status:       true,
		PrimaryIP:    testIPAddressIdentifier(FamilyIPv4, n),
		PrimaryIP4:   testIPAddressIdentifier(FamilyIPv4, n),
		PrimaryIP6:   testIPAddressIdentifier(FamilyIPv6, n),
		Comments:     "",
	}
}

func testDeviceIdentifier(n int) *DeviceIdentifier {
	return &DeviceIdentifier{
		ID:   n,
		Name: fmt.Sprintf("DeviceIdentifier %d", n),
	}
}

func testDeviceTypeIdentifier(n int) *DeviceTypeIdentifier {
	return &DeviceTypeIdentifier{
		ID:           n,
		Manufacturer: testSimpleIdentifier(n),
		Model:        fmt.Sprintf("device model %d", n),
		Slug:         fmt.Sprintf("devicetype%d", n),
	}
}

func testFormFactor(n int) *FormFactor {
	return &FormFactor{
		Value: n,
		Label: fmt.Sprintf("FormFactor %d", n),
	}
}

func testInterfaceConnectionIdentifier(n int) *InterfaceConnectionIdentifier {
	return &InterfaceConnectionIdentifier{
		ID:     n,
		Status: true,
	}
}

func testInterface(n int) *Interface {
	return &Interface{
		ID:                 n,
		Name:               fmt.Sprintf("interface %d", n),
		Device:             testDeviceIdentifier(n),
		FormFactor:         testFormFactor(n),
		MacAddress:         fmt.Sprintf("f4:9d:82:9e:34:c%d", n),
		MgmtOnly:           true,
		Description:        fmt.Sprintf("Description %d", n),
		ConnectedInterface: testInterfaceDetail(n),
		Connection:         testInterfaceConnectionIdentifier(n),
	}
}

func testInterfaceDetail(n int) *InterfaceDetail {
	return &InterfaceDetail{
		ID:          n,
		Device:      testDeviceIdentifier(n),
		FormFactor:  n,
		Name:        fmt.Sprintf("InterfaceDetail %d", n),
		MacAddress:  fmt.Sprintf("f4:9d:82:9e:34:c%d", n),
		MgmtOnly:    true,
		Description: fmt.Sprintf("Description %d", n),
	}
}

func testInterfaceIdentifier(n int) *InterfaceIdentifier {
	return &InterfaceIdentifier{
		ID:     n,
		Device: testDeviceIdentifier(n),
		Name:   fmt.Sprintf("InterfaceIdentifier %d", n),
	}
}

type testIPAddressPageData struct {
	Count       int    `json:"count"`
	NextURL     string `json:"next"`
	PreviousURL string `json:"previous"`
	Results     []*mockIPAddress
}

func testIPAddressPage(ips []*mockIPAddress) *testIPAddressPageData {
	return &testIPAddressPageData{
		Count:       len(ips),
		NextURL:     "",
		PreviousURL: "",
		Results:     ips,
	}
}

func testTenantIdentifier(n int) *TenantIdentifier {
	return &TenantIdentifier{
		ID:   n,
		Name: fmt.Sprintf("Tenant %d", n),
		Slug: fmt.Sprintf("tenant-%d", n),
	}
}

func testIPAMStatus(n int) *IPAMStatus {
	return &IPAMStatus{
		Value: n,
		Label: fmt.Sprintf("Label %d", n),
	}
}

func testIPAddress(family Family, n int) *IPAddress {
	address := &net.IPNet{
		IP:   net.IPv4(8, 8, 8, 0),
		Mask: net.CIDRMask(24, 32),
	}
	if family == FamilyIPv6 {
		address = &net.IPNet{
			IP:   net.ParseIP("2001:4860:4860::"),
			Mask: net.CIDRMask(48, 128),
		}
	}

	return &IPAddress{
		ID:          n,
		Family:      family,
		Address:     address,
		VRF:         testVRFIdentifier(n),
		Tenant:      testTenantIdentifier(n),
		Interface:   testInterface(n),
		Description: fmt.Sprintf("description %d", n),
		NATInside:   testIPAddressIdentifier(family, n),
		NATOutside:  testIPAddressIdentifier(family, n),
		Status:      testIPAMStatus(n),
	}
}

func testIPAddressIdentifier(family Family, n int) *IPAddressIdentifier {
	address := &net.IPNet{
		IP:   net.IPv4(8, 8, 8, 0),
		Mask: net.CIDRMask(24, 32),
	}
	if family == FamilyIPv6 {
		address = &net.IPNet{
			IP:   net.ParseIP("2001:4860:4860::"),
			Mask: net.CIDRMask(48, 128),
		}
	}

	return &IPAddressIdentifier{
		ID:      n,
		Family:  family,
		Address: address,
	}
}

func testPrefix(family Family, n int) *Prefix {
	prefix := &net.IPNet{
		IP:   net.IPv4(8, 8, 0, 0),
		Mask: net.CIDRMask(16, 32),
	}
	if family == FamilyIPv6 {
		prefix = &net.IPNet{
			IP:   net.ParseIP("2001:4860::"),
			Mask: net.CIDRMask(32, 128),
		}
	}

	return &Prefix{
		ID:          n,
		Family:      family,
		Prefix:      prefix,
		Site:        testSiteIdentifier(n),
		VRF:         testVRFIdentifier(n),
		VLAN:        testVLANIdentifier(n),
		Status:      StatusActive,
		Role:        testRoleIdentifier(n),
		Description: fmt.Sprintf("description %d", n),
	}
}

func testRackIdentifier(n int) *RackIdentifier {
	return &RackIdentifier{
		ID:          n,
		Name:        fmt.Sprintf("rack %d", n),
		FacilityID:  fmt.Sprintf("facility%d", n),
		DisplayName: fmt.Sprintf("Rack %d", n),
	}
}

func testPowerOutlet(n int) *PowerOutletIdentifier {
	return &PowerOutletIdentifier{
		ID:     n,
		Device: testDeviceIdentifier(n),
		Name:   fmt.Sprintf("poweroutlet %d", n),
	}
}

func testPowerPort(n int) *PowerPort {
	return &PowerPort{
		ID:               n,
		Name:             fmt.Sprintf("powerport %d", n),
		PowerOutlet:      testPowerOutlet(n),
		ConnectionStatus: true,
	}
}

func testPowerPortIdentifier(n int) *PowerPortIdentifier {
	return &PowerPortIdentifier{
		Device: fmt.Sprintf("device %d", n),
		Name:   fmt.Sprintf("rc power port %d", n),
		Outlet: fmt.Sprintf("outlet %d", n),
	}
}

func testRelatedConnection(n int) *RelatedConnection {
	return &RelatedConnection{
		Device:       testDevice(n),
		ConsolePorts: []*ConsolePort{testConsolePort(n)},
		Interfaces:   []*Interface{testInterface(n)},
		PowerPorts:   []*PowerPort{testPowerPort(n)},
	}
}

func testRIR(n int) *RIR {
	return &RIR{
		ID:   n,
		Name: fmt.Sprintf("RIR %d", n),
		Slug: fmt.Sprintf("rir%d", n),
	}
}

func testRIRIdentifier(n int) *RIRIdentifier {
	return &RIRIdentifier{
		ID:   n,
		Name: fmt.Sprintf("RIRIdentifier %d", n),
		Slug: fmt.Sprintf("riridentifier%d", n),
	}
}

func testRole(n int) *Role {
	return &Role{
		ID:     n,
		Name:   fmt.Sprintf("Role %d", n),
		Slug:   fmt.Sprintf("role%d", n),
		Weight: n,
	}
}

func testRoleIdentifier(n int) *RoleIdentifier {
	return &RoleIdentifier{
		ID:   n,
		Name: fmt.Sprintf("RoleIdentifier %d", n),
		Slug: fmt.Sprintf("roleidentifier%d", n),
	}
}

func testSimpleIdentifier(n int) *SimpleIdentifier {
	return &SimpleIdentifier{
		ID:   n,
		Name: fmt.Sprintf("simple %d", n),
		Slug: fmt.Sprintf("simple%d", n),
	}
}

func testSite(n int) *Site {
	return &Site{
		ID:              n,
		Name:            fmt.Sprintf("Site %d", n),
		Slug:            fmt.Sprintf("site%d", n),
		Facility:        fmt.Sprintf("Facility %d", n),
		ASN:             ASN(n),
		PhysicalAddress: fmt.Sprintf("%d Facility Street, City, State 12345", n),
		ShippingAddress: fmt.Sprintf("%d Facility Street, ATTN: Shipping & Receiving, City, State 12345", n),
		Comments:        fmt.Sprintf("comment %d", n),
		CountPrefixes:   n,
		CountVLANs:      n,
		CountRacks:      n,
		CountDevices:    n,
		CountCircuits:   n,
	}
}

func testSiteIdentifier(n int) *SiteIdentifier {
	return &SiteIdentifier{
		ID:   n,
		Name: fmt.Sprintf("SiteIdentifier %d", n),
		Slug: fmt.Sprintf("siteidentifier%d", n),
	}
}

func testVLAN(n int) *VLAN {
	return &VLAN{
		ID:          n,
		Site:        testSiteIdentifier(n),
		VID:         VLANID(n),
		Name:        fmt.Sprintf("vlan %d", n),
		Status:      StatusActive,
		Role:        testRoleIdentifier(n),
		DisplayName: fmt.Sprintf("VLAN %d", n),
	}
}

func testVLANIdentifier(n int) *VLANIdentifier {
	return &VLANIdentifier{
		ID:          n,
		VID:         VLANID(n),
		Name:        fmt.Sprintf("vlanidentifier %d", n),
		DisplayName: fmt.Sprintf("VLANIdentifier %d", n),
	}
}

func testVRF(n int) *VRF {
	return &VRF{
		ID:          n,
		Name:        fmt.Sprintf("VRF %d", n),
		RD:          fmt.Sprintf("rd %d", n),
		Description: fmt.Sprintf("description %d", n),
	}
}

func testVRFIdentifier(n int) *VRFIdentifier {
	return &VRFIdentifier{
		ID:   n,
		Name: fmt.Sprintf("VRFIdentifier %d", n),
		RD:   fmt.Sprintf("rd %d", n),
	}
}
