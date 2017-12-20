package netbox

// A Circuit is a representation of netbox circuit.
type Circuit struct {
	ID  int    `json:"id"`
	CID string `json:"cid"`
}
