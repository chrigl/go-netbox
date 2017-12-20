package netbox

// A PrefixRole represents a netbox prefix/vrf role.
type PrefixRole struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Slug   string `json:"slug"`
	Weight int    `json:"weight"`
}
