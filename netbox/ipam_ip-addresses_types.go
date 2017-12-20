package netbox

import (
	"encoding/json"
	"errors"
	"net"
	"net/url"
)

// A IPAddress is a representation of netbox ip-address.
type IPAddress struct {
	ID          int            `json:"id"`
	Family      int            `json:"family"`
	Description string         `json:"description"`
	Address     *Address       `json:"address"`
	Vrf         *VRF           `json:"vrf"`
	Tenant      *Tenant        `json:"tenant"`
	Status      *AddressStatus `json:"status"`
	Role        *PrefixRole    `json:"role"`
	Interface   *Interface     `json:"interface"`
	NatInside   *IPAddress     `json:"nat_inside"`
	NatOutside  *IPAddress     `json:"nat_outside"`
}

// A writableIPAddress corresponds to the Netbox API's
// writable serializer for an IPAddress. It is used transparently
// when IPAddress are serialized into JSON.
type writableIPAddress struct {
	ID int `json:"id"`
}

// ListIPAddressOptions is used as an argument for Client.IPAM.IPAddresses.List.
type ListIPAddressOptions struct {
}

// Values generates a url.Values map from the tata in ListIPAddressOptions.
func (o *ListIPAddressOptions) Values() (url.Values, error) {
	if o == nil {
		return nil, nil
	}

	v := url.Values{}

	return v, nil
}

// AddressStatus represents the status of an IPAddress.
type AddressStatus struct {
	Value int    `json:"value"`
	Label string `json:"label"`
}

// Address represents the effective ip address in form of a net.IPNet.
type Address struct {
	net.IPNet
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// The address is expected to be quoted string in CIDR form.
func (a *Address) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	if s == "" {
		return nil
	}

	ip, ipNet, err := net.ParseCIDR(s)
	if err != nil {
		return err
	}

	a.IPNet.Mask = ipNet.Mask
	a.IPNet.IP = ipNet.IP

	if !ip.Equal(ipNet.IP) {
		a.IPNet.IP = ip
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
// The address is a quoted string in CIDR form.
func (a Address) MarshalJSON() ([]byte, error) {
	if a.IPNet.IP == nil || a.IPNet.Mask == nil {
		return nil, errors.New("address.MarshalJSON: inner IPNet not defined")
	}
	return []byte(`"` + a.IPNet.String() + `"`), nil
}

//go:generate go run generate_functions.go -type-name IPAddress -update-type-name writableIPAddress -service-name IPAddressesService -endpoint ipam -service ip-addresses
//go:generate go run generate_basic_tests.go -type-name IPAddress -service-name IPAddressesService -endpoint ipam -service ip-addresses -client-endpoint IPAM -client-service IPAddresses
