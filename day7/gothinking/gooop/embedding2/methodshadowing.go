package main

import "fmt"

type base struct {
	a string
	b int
}

//method xyz
func (this base) xyz() {
	fmt.Println("xyz, a is:", this.a)
}

//method display
func (this base) display() {
	fmt.Println("base, a is:", this.a)
}

type derived struct {
	base // embedding1
	d    int
	a    float32 //-SHADOWED
}

//method display -SHADOWED
func (this derived) display() {
	fmt.Println("derived a is:", this.a)
}

func main() {
	var a = derived{base{"base-a", 10}, 20, 2.5}

	a.display() // calls Derived/display(a)
	// => "derived, a is: 2.5"

	a.base.display() // calls Base/display(a.base), the base implementation
	// => "base, a is: base-a"

	a.xyz() // "xyz" was not shadowed, calls Base/xyz(a.base)
	// => "xyz, a is: base-a"
}
