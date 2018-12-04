package map_1

import (
	"testing"
)

func TestNoDuplicates(t *testing.T) {
	arr := []struct{
		str, expected string
	} {
		{"ababab", "ab"},
		{"abcdefbgac", "cdefbga"},
	}

	for _, s := range arr {
		actual := NoDuplicates(s.str)
		if actual != s.expected {
			t.Errorf("FAIL: \nExpected: %s\nActual: %s",
				s.expected, actual)
		}
	}
}
