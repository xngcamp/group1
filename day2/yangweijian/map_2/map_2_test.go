package map_2

import "testing"

func TestTranslate(t *testing.T) {
	arr := []struct{
		str, expected string
	} {
		{"AUGUUUUCUUAAAUG", "MethioninePhenylalanineSerine"},
	}

	for _, s := range arr {
		actual := Translate(s.str)
		if actual != s.expected {
			t.Errorf("FAIL: \nExpected: %s\nActual: %s",
				s.expected, actual)
		}
	}
}