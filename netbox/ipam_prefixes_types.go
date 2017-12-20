package netbox

import (
	"encoding/json"
	"errors"
	"net"
	"net/url"
)

// A Prefix is a representation of netbox prefix.
type Prefix struct {
	ID          int           `json:"id"`
	Family      int           `json:"family"`
	Prefix      *Network      `json:"prefix"`
	Site        *Site         `json:"site"`
	Vrf         *VRF          `json:"vrf"`
	Tenant      *Tenant       `json:"tenant"`
	Vlan        *VLAN         `json:"vlan"`
	Status      *PrefixStatus `json:"status"`
	Role        *PrefixRole   `json:"role"`
	IsPool      bool          `json:"is_pool"`
	Description string        `json:"description"`
}

// A writablePrefix corresponds to the Netbox API's
// writable serializer for a Prefix. It is used transparently
// when Prefix are serialized into JSON.
type writablePrefix struct {
	ID int `json:"id"`
}

// ListPrefixOptions is used as an argument for Client.IPAM.Prefixes.List.
type ListPrefixOptions struct {
}

// Values generates a url.Values map from the tata in ListPrefixOptions.
func (o *ListPrefixOptions) Values() (url.Values, error) {
	if o == nil {
		return nil, nil
	}

	v := url.Values{}

	return v, nil
}

// PrefixStatus represents the status of a Prefix.
type PrefixStatus struct {
	Value int    `json:"value"`
	Label string `json:"label"`
}

// Network represents the effective ip network in form of a net.IPNet.
type Network struct {
	net.IPNet
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// The network is expected to be quoted string in CIDR form.
func (n *Network) UnmarshalJSON(data []byte) error {
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

	_, ipNet, err := net.ParseCIDR(s)
	if err != nil {
		return err
	}

	n.IPNet.Mask = ipNet.Mask
	n.IPNet.IP = ipNet.IP

	return nil
}

// MarshalJSON implements the json.Marshaler interface.
// The network is a quoted string in CIDR form.
func (n Network) MarshalJSON() ([]byte, error) {
	if n.IPNet.IP == nil || n.IPNet.Mask == nil {
		return nil, errors.New("network.MarshalJSON: inner IPNet not defined")
	}
	return []byte(`"` + n.IPNet.String() + `"`), nil
}

//go:generate go run generate_functions.go -type-name Prefix -update-type-name writablePrefix -service-name PrefixesService -endpoint ipam -service prefixes
//go:generate go run generate_basic_tests.go -type-name Prefix -service-name PrefixesService -endpoint ipam -service prefixes -client-endpoint IPAM -client-service Prefixes
