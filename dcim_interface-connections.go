package netbox

// An InterfaceConnectionIdentifier is a reduced version of an
// InterfaceConnection, returned as a nested object in some top-level objects.
// It contains information wich can be used in subsequent API calls to
// identify and retrieve a full InterfaceConnection
type InterfaceConnectionIdentifier struct {
	ID     int  `json:"id"`
	Status bool `json:"connection_status"`
}
