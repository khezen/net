package ip

import "testing"

func TestSubnetIPV4(t *testing.T) {
	cases := []struct {
		ipv4   string
		mask   int
		subnet string
		err    error
	}{
		{"192.168.220.254", 17, "192.168.92.0/17", nil},
		{"192.168.220.254", 5, "192.0.0.0/5", nil},
		{"192.168.220.254", 1, "64.0.0.0/1", nil},
		{"192.168.220.254", 31, "192.168.220.252/31", nil},
		{"192.168.220.254", 9, "192.40.0.0/9", nil},
		{"192.168.220.254", 0, "0.0.0.0/0", nil},
		{"192.168.220.254", 32, "192.168.220.254/32", nil},
		{"192.168.220.254", -1, "", ErrMaskOutOfBound},
		{"192.168.220.254", 33, "", ErrMaskOutOfBound},
	}
	for _, c := range cases {
		subnet, err := SubnetIPV4(c.ipv4, c.mask)
		if subnet != c.subnet {
			t.Errorf("expected %vv, got %v", c.subnet, subnet)
		}
		if err != c.err {
			t.Errorf("expected %vv, got %v", c.err, err)
		}
	}
}
