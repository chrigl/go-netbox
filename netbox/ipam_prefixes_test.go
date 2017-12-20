package netbox

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"reflect"
	"testing"
)

type testIPNetwork struct {
	Network *Network `json:"network"`
}

func TestNetworkUnmashalJSONWithinStruct(t *testing.T) {
	data := `{"network": "8.8.8.8/24"}`
	expect := net.IPNet{
		IP:   net.ParseIP("8.8.8.0"),
		Mask: net.CIDRMask(24, 32),
	}

	var ip testIPNetwork
	err := json.Unmarshal([]byte(data), &ip)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if want, got := expect.IP, ip.Network.IP; !want.Equal(got) {
		t.Fatalf("unexpected ip:\n- want: %v\n-  got: %v", want, got)
	}
	if want, got := expect.Mask, ip.Network.Mask; !bytes.Equal(want, got) {
		t.Fatalf("unexpected mask:\n- want: %v\n-  got: %v", want, got)
	}

}

func TestNetworkUnmarshalJSON(t *testing.T) {
	var tests = []struct {
		data []byte
		want net.IPNet
		err  error
	}{
		{
			data: []byte(`"` + "8.8.8.8/24" + `"`),
			want: net.IPNet{
				IP:   net.ParseIP("8.8.8.0"),
				Mask: net.CIDRMask(24, 32),
			},
			err: nil,
		},
		{
			data: []byte(`"` + "2001:4860:4860::8888/48" + `"`),
			want: net.IPNet{
				IP:   net.ParseIP("2001:4860:4860::"),
				Mask: net.CIDRMask(48, 128),
			},
			err: nil,
		},
		{
			data: []byte(`""`),
			want: net.IPNet{},
			err:  nil,
		},
		{
			data: []byte("null"),
			want: net.IPNet{},
			err:  nil,
		},
		{
			data: []byte(`"` + "Hello World" + `"`),
			want: net.IPNet{},
			err:  &net.ParseError{"CIDR address", "Hello World"},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("[%d] UnmarshalJSON of %s", i, string(tt.data)), func(t *testing.T) {
			var ipNet Network
			err := ipNet.UnmarshalJSON(tt.data)
			if want, got := err, tt.err; !reflect.DeepEqual(want, got) {
				t.Fatalf("unexpected error:\n- want: %v\n-  got: %v", want, got)
			}
			if want, got := tt.want.IP, ipNet.IP; !want.Equal(got) {
				t.Fatalf("unexpected ipNet:\n- want: %v\n-  got: %v", want, got)
			}
			if want, got := tt.want.Mask, ipNet.Mask; !bytes.Equal(want, got) {
				t.Fatalf("unexpected mask:\n- want: %v\n-  got: %v", want, got)
			}
		})
	}
}

func TestNetworkMarshalJSON(t *testing.T) {
	var tests = []struct {
		ipNet Network
		err   error
	}{
		{
			ipNet: Network{
				IPNet: net.IPNet{
					IP:   net.ParseIP("8.8.8.0"),
					Mask: net.CIDRMask(24, 32),
				},
			},
			err: nil,
		},
		{
			ipNet: Network{
				IPNet: net.IPNet{
					IP:   net.ParseIP("2001:4860:4860::"),
					Mask: net.CIDRMask(48, 128),
				},
			},
			err: nil,
		},
		{
			ipNet: Network{},
			err:   errors.New("network.MarshalJSON: inner IPNet not defined"),
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("[%d] MarshalJSON of %s", i, tt.ipNet.IPNet.String()), func(t *testing.T) {
			data, err := tt.ipNet.MarshalJSON()
			if want, got := tt.err, err; !reflect.DeepEqual(want, got) {
				t.Fatalf("unexpected error:\n- want: %v\n-  got: %v", want, got)
			}
			if err != nil {
				// No need to do further unmarshaling.
				return
			}

			var ipNet Network
			err = ipNet.UnmarshalJSON(data)
			if err != nil {
				t.Fatalf("unexpected error from unmarshaling just created data: %v", err)
			}
			if want, got := tt.ipNet.IP, ipNet.IP; !want.Equal(got) {
				t.Fatalf("unexpected ipNet:\n- want: %v\n-  got: %v", want, got)
			}
			if want, got := tt.ipNet.Mask, ipNet.Mask; !bytes.Equal(want, got) {
				t.Fatalf("unexpected mask:\n- want: %v\n-  got: %v", want, got)
			}
		})
	}
}

func TestNetworkMarshalJSONWithinStruct(t *testing.T) {
	expect := testIPNetwork{
		Network: &Network{
			IPNet: net.IPNet{
				IP:   net.ParseIP("8.8.8.0"),
				Mask: net.CIDRMask(24, 32),
			},
		},
	}

	data, err := json.Marshal(expect)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var ip testIPNetwork
	err = json.Unmarshal([]byte(data), &ip)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if want, got := expect.Network.IP, ip.Network.IP; !want.Equal(got) {
		t.Fatalf("unexpected ip:\n- want: %v\n-  got: %v", want, got)
	}
	if want, got := expect.Network.Mask, ip.Network.Mask; !bytes.Equal(want, got) {
		t.Fatalf("unexpected mask:\n- want: %v\n-  got: %v", want, got)
	}

}
