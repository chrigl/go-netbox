package netbox

import "net/url"

type PrefixStatus struct {
	Value int    `json:"value"`
	Label string `json:"label"`
}

type Prefix struct {
	ID     int `json:"id"`
	Family int `json:"family"`
	// Prefix      *net.IPNet    `json:"prefix"`
	Prefix      string        `json:"prefix"`
	Site        *Site         `json:"site"`
	Vrf         *VRF          `json:"vrf"`
	Tenant      *Tenant       `json:"tenant"`
	Vlan        *VLAN         `json:"vlan"`
	Status      *PrefixStatus `json:"status"`
	Role        *PrefixRole   `json:"role"`
	IsPool      bool          `json:"is_pool"`
	Description string        `json:"description"`
}

type writablePrefix struct {
	ID int `json:"id"`
}

type ListPrefixOptions struct {
}

func (o *ListPrefixOptions) Values() (url.Values, error) {
	if o == nil {
		return nil, nil
	}

	v := url.Values{}

	return v, nil
}

//go:generate go run generate_functions.go -type-name Prefix -update-type-name writablePrefix -service-name PrefixesService -endpoint ipam -service prefixes
//go:generate go run generate_basic_tests.go -type-name Prefix -service-name PrefixesService -endpoint ipam -service prefixes -client-endpoint IPAM -client-service Prefixes
