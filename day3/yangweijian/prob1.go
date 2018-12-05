package main

import "fmt"

const PI = 3.14

type Graph interface {
	Area() float64
}

type Triangle struct {
	bottom, height float64
}

func (t Triangle) Area() float64 {
	return t.bottom * t.height / 2
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * PI
}

type Rectangle struct {
	a, b float64
}

func (r Rectangle) Area() float64 {
	return r.a * r.b
}

func main() {
	var g Graph
	g = Triangle{2.0, 4.0}
	fmt.Printf("三角形：%.2f\n", g.Area())

	g = Circle{2.0}
	fmt.Printf("圆形：%.2f\n", g.Area())

	g = Rectangle{2.0, 4.0}
	fmt.Printf("矩形：%.2f\n", g.Area())
}
