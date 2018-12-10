package main

import (
	"camp/src/day_3/jiekou/circle"
	"camp/src/day_3/jiekou/rectangle"
	"camp/src/day_3/jiekou/triangle"
	"fmt"
)

type Graph interface {
	Area() float64
}
func CaculArea(g Graph) float64{
	return  g.Area();
}

func main()  {
	var t Graph
	t =triangle.Triangle{4,5,}
	s := CaculArea(t)
	fmt.Printf("%f",s)

	t =rectangle.Rectangle{5,6}
	p := CaculArea(t)
	fmt.Printf("%f",p)

	t=circle.Circle{5}
	x := CaculArea(t)
	fmt.Printf("%f",x)
}

