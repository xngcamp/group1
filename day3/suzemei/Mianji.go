package main

import "fmt"

type Graph interface {
     Area() float64
}
type Circle struct {
	R float64
}
func (c Circle)Area() float64{
	return c.R*c.R*0.5
}
type Rectangle struct {
	l float64
	w float64
}
func (r Rectangle)Area() float64{
	return r.w*r.l
}
type Triangle struct {
	h float64
	d float64
}
func (t Triangle)Area() float64{
	return t.h*t.d*0.5
}
func main(){
	var  a Triangle=Triangle{3,4}
	fmt.Println("三角形面积为：",a.Area())
	var  b Circle=Circle{3}
	fmt.Println("圆面积为：",b.Area())
	var  c Rectangle=Rectangle{5,5}
	fmt.Println("矩形面积为：",c.Area())
}
