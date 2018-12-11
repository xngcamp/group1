package main

import (
	"fmt"
	"testing"
)

func TestStudentSayHi(t *testing.T) {
	s := Student{}
	s.Init("John", 25, "cs")
	s.SayHi()
	fmt.Printf("name is %s\n", s.name)
}


