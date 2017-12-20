package netbox

// A VLANGroup is a representation of netk vlan-group.
type VLANGroup struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	Site *Site  `json:"site"`
}
