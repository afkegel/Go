package tutorial

import "testing"

func TestString(t *testing.T) {
	// Tests the IP stringer method by comparing an expected result against the
	// actual result.
	result := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	expected := map[string]string{
		"loopback":  "127.0.0.1",
		"googleDNS": "8.8.8.8",
	}
	for name := range result {
		if result[name].String() != expected[name] {
			t.Errorf("IPAddr String() failed!")
		}
	}
}
