package netbox

type VLANStatus struct {
	Value int    `json:"value"`
	Label string `json:"label"`
}

type VLAN struct {
	ID          int         `json:"id"`
	Site        *Site       `json:"site"`
	Group       *VLANGroup  `json:"group"`
	Vid         int         `json:"vid"`
	Name        string      `json:"name"`
	Tenant      *Tenant     `json:"tenant"`
	Status      *VLANStatus `json:"status"`
	Role        *PrefixRole `json:"role"`
	Description string      `json:"description"`
	DisplayName string      `json:"display_name"`
}
