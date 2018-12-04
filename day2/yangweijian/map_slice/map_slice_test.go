package map_slice

import (
	"reflect"
	"testing"
)

func TestCrossover(t *testing.T) {
	arr := []struct{
		ns, xs, ys, r1, r2 []int
	} {
		{
			[]int{1, 3},
			[]int{1,2,3,4,5,6},
			[]int{7,8,9,10,11,12},
			[]int{1,8,9,4,5,6},
			[]int{7,2,3,10,11,12},
		},
		{
			[]int{1},
			[]int {1,2,3},
			[]int {10,11,12},
			[]int{1,11,12},
			[]int{10,2,3},
		},
	}

	for _, s := range arr {
		actualR1, actualR2 := Crossover(s.ns, s.xs, s.ys)
		if !reflect.DeepEqual(actualR1, s.r1) || !reflect.DeepEqual(actualR2, s.r2) {
			t.Errorf("map_slice has error")
		}
	}
}
