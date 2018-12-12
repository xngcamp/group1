package main

import "fmt"

type base struct {
	a string
	b int
}

type derived struct {
	base // embedding1
	d    int
	a    float32 //-SHADOWS
}

/*
When working with structs and embedding1, everything is STATICALLY LINKED. All references are resolved at compile time.

*/

func main() {
	var x derived
	fmt.Printf("%T\n", x.a)      //=> x.a, float32 (derived.a shadows base.a)
	fmt.Printf("%T\n", x.base.a) //=> x.base.a, string (accessing shadowed member)

	//
	fmt.Printf("%T\n", x.b)
}
