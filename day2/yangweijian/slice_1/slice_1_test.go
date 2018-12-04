package slice_1

import "testing"

func TestSolve(t *testing.T) {
	arr := []struct{
		str string
		left int
		expected int
	} {
		{"((1)23(45))(aB)", 0, 10},
		{"((1)23(45))(aB)", 1, 3},
		{"((1)23(45))(aB)", 2, -1},
		{"((1)23(45))(aB)", 6, 9},
		{"((1)23(45))(aB)", 11, 14},
		{"((>)|?(*'))(yZ)", 11, 14},
	}

	for _, s := range arr {
		actual := Solve(s.str, s.left)
		if actual != s.expected {
			t.Errorf("slice_1_test has error")
		}
	}
}
