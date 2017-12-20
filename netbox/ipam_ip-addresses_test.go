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

type testIPAddr struct {
	Address *Address `json:"address"`
}

func TestAddressWithinStruct(t *testing.T) {
	data := `{"address": "8.8.8.8/24"}`
	expect := net.IPNet{
		IP:   net.ParseIP("8.8.8.8"),
		Mask: net.CIDRMask(24, 32),
	}

	var ip testIPAddr
	err := json.Unmarshal([]byte(data), &ip)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if want, got := expect.IP, ip.Address.IP; !want.Equal(got) {
		t.Fatalf("unexpected ip:\n- want: %v\n-  got: %v", want, got)
	}
	if want, got := expect.Mask, ip.Address.Mask; !bytes.Equal(want, got) {
		t.Fatalf("unexpected mask:\n- want: %v\n-  got: %v", want, got)
	}

}

func TestAddressUnmarshalJSON(t *testing.T) {
	var tests = []struct {
		data []byte
		want net.IPNet
		err  error
	}{
		{
			data: []byte(`"` + "8.8.8.8/24" + `"`),
			want: net.IPNet{
				IP:   net.ParseIP("8.8.8.8"),
				Mask: net.CIDRMask(24, 32),
			},
			err: nil,
		},
		{
			data: []byte(`"` + "8.8.8.0/24" + `"`),
			want: net.IPNet{
				IP:   net.ParseIP("8.8.8.0"),
				Mask: net.CIDRMask(24, 32),
			},
			err: nil,
		},
		{
			data: []byte(`"` + "2001:4860:4860::8888/48" + `"`),
			want: net.IPNet{
				IP:   net.ParseIP("2001:4860:4860::8888"),
				Mask: net.CIDRMask(48, 128),
			},
			err: nil,
		},
		{
			data: []byte(`"` + "2001:4860:4860::/48" + `"`),
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
			var ip Address
			err := ip.UnmarshalJSON(tt.data)
			if want, got := err, tt.err; !reflect.DeepEqual(want, got) {
				t.Fatalf("unexpected error:\n- want: %v\n-  got: %v", want, got)
			}
			if want, got := tt.want.IP, ip.IP; !want.Equal(got) {
				t.Fatalf("unexpected ip:\n- want: %v\n-  got: %v", want, got)
			}
			if want, got := tt.want.Mask, ip.Mask; !bytes.Equal(want, got) {
				t.Fatalf("unexpected mask:\n- want: %v\n-  got: %v", want, got)
			}
		})
	}
}

func TestAddressMarshalJSON(t *testing.T) {
	var tests = []struct {
		ipAddress Address
		err       error
	}{
		{
			ipAddress: Address{
				IPNet: net.IPNet{
					IP:   net.ParseIP("8.8.8.0"),
					Mask: net.CIDRMask(24, 32),
				},
			},
			err: nil,
		},
		{
			ipAddress: Address{
				IPNet: net.IPNet{
					IP:   net.ParseIP("2001:4860:4860::"),
					Mask: net.CIDRMask(48, 128),
				},
			},
			err: nil,
		},
		{
			ipAddress: Address{},
			err:       errors.New("address.MarshalJSON: inner IPNet not defined"),
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("[%d] MarshalJSON of %s", i, tt.ipAddress.IPNet.String()), func(t *testing.T) {
			data, err := tt.ipAddress.MarshalJSON()
			if want, got := tt.err, err; !reflect.DeepEqual(want, got) {
				t.Fatalf("unexpected error:\n- want: %v\n-  got: %v", want, got)
			}
			if err != nil {
				// No need to do further unmarshaling.
				return
			}

			var ipAddress Address
			err = ipAddress.UnmarshalJSON(data)
			if err != nil {
				t.Fatalf("unexpected error from unmarshaling just created data: %v", err)
			}
			if want, got := tt.ipAddress.IP, ipAddress.IP; !want.Equal(got) {
				t.Fatalf("unexpected ipAddress:\n- want: %v\n-  got: %v", want, got)
			}
			if want, got := tt.ipAddress.Mask, ipAddress.Mask; !bytes.Equal(want, got) {
				t.Fatalf("unexpected mask:\n- want: %v\n-  got: %v", want, got)
			}
		})
	}
}

func TestAddressMarshalJSONWithinStruct(t *testing.T) {
	expect := testIPAddr{
		Address: &Address{
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

	var ip testIPAddr
	err = json.Unmarshal([]byte(data), &ip)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if want, got := expect.Address.IP, ip.Address.IP; !want.Equal(got) {
		t.Fatalf("unexpected ip:\n- want: %v\n-  got: %v", want, got)
	}
	if want, got := expect.Address.Mask, ip.Address.Mask; !bytes.Equal(want, got) {
		t.Fatalf("unexpected mask:\n- want: %v\n-  got: %v", want, got)
	}

}
