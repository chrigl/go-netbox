package netbox

type Interface struct {
	ID         int     `json:"id"`
	Device     *Device `json:"device"`
	Name       string  `json:"name"`
	FormFactor struct {
		Value int    `json:"value"`
		Label string `json:"label"`
	} `json:"form_factor"`
	Enabled             bool                 `json:"enabled"`
	Lag                 interface{}          `json:"lag"`
	Mtu                 interface{}          `json:"mtu"`
	MacAddress          interface{}          `json:"mac_address"`
	MgmtOnly            bool                 `json:"mgmt_only"`
	Description         string               `json:"description"`
	IsConnected         bool                 `json:"is_connected"`
	InterfaceConnection *InterfaceConnection `json:"interface_connection"`
	CircuitTermination  *CircuitTermination  `json:"circuit_termination"`
}

type InterfaceConnection struct {
	Interface *Interface `json:"interface"`
	Status    bool       `json:"status"`
}
