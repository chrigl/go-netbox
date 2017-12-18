package netbox

type Site struct {
	ID              int     `json:"id"`
	Name            string  `json:"name"`
	Slug            string  `json:"slug"`
	Region          *Region `json:"region"`
	Tenant          *Tenant `json:"tenant"`
	Facility        string  `json:"facility"`
	Asn             int     `json:"asn"`
	PhysicalAddress string  `json:"physical_address"`
	ShippingAddress string  `json:"shipping_address"`
	ContactName     string  `json:"contact_name"`
	ContactPhone    string  `json:"contact_phone"`
	ContactEmail    string  `json:"contact_email"`
	Comments        string  `json:"comments"`
	CountPrefixes   int     `json:"count_prefixes"`
	CountVlans      int     `json:"count_vlans"`
	CountRacks      int     `json:"count_racks"`
	CountDevices    int     `json:"count_devices"`
	CountCircuits   int     `json:"count_circuits"`
}
