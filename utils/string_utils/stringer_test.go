package string_utils

import "testing"

func TestIPAddr_String(t *testing.T) {
	var testsTable = []struct {
		name     string
		IPAddr   IPAddr
		expected string
	}{
		{"empty IPAddr", IPAddr{}, "0.0.0.0"},
		{"IPAddr with missing values", IPAddr{1, 4}, "1.4.0.0"},
		{"valid IPAddr #1", IPAddr{127, 0, 0, 1}, "127.0.0.1"},
		{"valid IPAddr #2", IPAddr{192, 168, 44, 128}, "192.168.44.128"},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			got := test.IPAddr.String()
			if got != test.expected {
				t.Errorf("\nToUpperCase(%v) \nexpected: %s \ngot: %s", test.IPAddr, test.expected, got)
			}
		})
	}
}
