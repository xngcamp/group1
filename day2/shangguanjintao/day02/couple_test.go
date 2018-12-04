package day02

import (
	"fmt"
	"testing"
)

func TestCouple(t *testing.T) {
	s := "((1)23(45))(aB)"
	m := 6
	fmt.Print(couple(s, m))
}
