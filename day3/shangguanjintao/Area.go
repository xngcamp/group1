package main

import "fmt"
type Graph interface {
	Area() float64
}

type Rectangle struct {
	H float64
	L float64
}

func (r Rectangle) Area() float64{
	area:= (r.H * r.L)
	return  area
}
type Triangle struct {
	H float64
	L float64
}

func (t Triangle) Area() float64{
	area:= (t.H * t.L)/2
	return  area
}
type Cricle struct {
	R float64
}

func (c Cricle) Area() float64{
	area:= c.R *c.R*3.14
	return  area
}
func main(){
	//r := Rectangle{H:10.0,L:5.0}
	var g Graph
	g = Rectangle{H:10.0,L:5.0}
	fmt.Println(g.Area())
	g = Triangle{H:6.0,L:6.0}
	fmt.Println(g.Area())
	g = Cricle{R:8.0}
	fmt.Println(g.Area())
}