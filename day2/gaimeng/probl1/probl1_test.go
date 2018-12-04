package probl1

import "testing"


func TestMaxlength(t *testing.T) {
	var tests = [] struct {
		a    []int
		want int
	}{{[]int{1, 3, 4, 5, 6, 9}, 2},
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{998, 999, 1000}, 2},
	}

	for _, c := range tests{
		got := Maxlength(c.a)
		if got != c.want {
		t.Errorf( "Got %d, want %d",got,c.want)
	}
	}
}

