package netbox

type CircuitTermination struct {
	ID            int            `json:"id"`
	Circuit       *nestedCircuit `json:"circuit"`
	TermSide      string         `json:"term_side"`
	Site          *Site          `json:"site"`
	Interface     *Interface     `json:"interface"`
	PortSpeed     float64        `json:"port_speed"`
	UpstreamSpeed float64        `json:"upstream_speed"`
	XconnectID    string         `json:"xconnect_id"`
	PpInfo        string         `json:"pp_info"`
}
