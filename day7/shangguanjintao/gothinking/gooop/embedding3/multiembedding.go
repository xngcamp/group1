package main

import "fmt"

type NamedObj struct {
	Name string
}

type Shape struct {
	NamedObj  //inheritance
	color     int32
	isRegular bool
}

type Point struct {
	x, y float64
}

type Rectangle struct {
	NamedObj            //multiple inheritance
	Shape               //^^
	center        Point //standard composition
	Width, Height float64
}

func main() {
	var aRect = Rectangle{
		NamedObj{"name1"},
		Shape{NamedObj{"name2"}, 0, true},
		Point{0, 0},
		20, 2.5,
	}

	fmt.Println(aRect.Name)
	fmt.Println(aRect.Shape)
	fmt.Println(aRect.Shape.Name)
}
