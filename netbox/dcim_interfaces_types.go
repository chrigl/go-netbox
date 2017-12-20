package netbox

// An Interface is a representation of netbox interface.
type Interface struct {
	ID                  int                  `json:"id"`
	Device              *Device              `json:"device"`
	Name                string               `json:"name"`
	FormFactor          *InterfaceFormFactor `json:"form_factor"`
	Enabled             bool                 `json:"enabled"`
	Lag                 interface{}          `json:"lag"`
	Mtu                 interface{}          `json:"mtu"`
	MacAddress          interface{}          `json:"mac_address"`
	MgmtOnly            bool                 `json:"mgmt_only"`
	Description         string               `json:"description"`
	IsConnected         bool                 `json:"is_connected"`
	InterfaceConnection *ConnectedInterface  `json:"interface_connection"`
	CircuitTermination  *CircuitTermination  `json:"circuit_termination"`
}

// ConnectedInterface represents the Interface, the currently viewed
// Interface is connected to. Not to confuse with InterfaceConnection.
type ConnectedInterface struct {
	Interface *Interface `json:"interface"`
	Status    bool       `json:"status"`
}

// InterfaceFormFactor represents the form factor of an Interface.
type InterfaceFormFactor struct {
	Value int    `json:"value"`
	Label string `json:"label"`
}
