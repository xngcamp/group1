package exam

import (
	"reflect"
	"testing"
)

func TestGraph_Area(t *testing.T)  {
	var tests = []struct{
		input_g Graph
		want_area float64
	}{
		{Circle{5},PI*5*5},
		{Circle{2.5},PI*2.5*2.5},
		{Circle{-2.5},-1},

		{Rectangle{10,10},100},
		{Rectangle{10,2.5},25},
		{Rectangle{2.5,10},25},
		{Rectangle{2.5,2.5},6.25},
		{Rectangle{0,0},0},
		{Rectangle{-10,-2.5},-1},

		{Triangle{10,10},50},
		{Triangle{10,2.5},12.5},
		{Triangle{2.5,10},12.5},
		{Rectangle{0,0},0},
		{Triangle{2.5,2.5},3.125},
		{Triangle{-10,-2.5},-1},
	}

	for _,row := range tests{
		var g Graph = row.input_g
		if row.want_area != g.Area() {
			t.Errorf("input Graph is %v%v,want %f,but got %f",reflect.TypeOf(g),row.input_g,row.want_area,g.Area())
		}
	}
}