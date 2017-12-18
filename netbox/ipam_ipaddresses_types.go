package netbox

import "net"

type nestedIPAddress struct {
	ID      int
	Family  int
	Address *net.IPNet
}
