package main

import (
	"coolgo/oop/student1"
	"fmt"
	"gothinking/oop/student2"
)

func TestStudent1() {
	s1 := student1.Student{}
	s1.Init("John", 25, "cs")
	s1.SayHi()

	// Error: implicit assignment of unexported field 'name' in student1.Student
	//s2 := student1.Student{"Aimi", 24, "math"}
	//s2.SayHi()

	//Error: s1.name undefined (cannot refer to unexported field or method name)
	//fmt.Println(s1.name)
}

func TestStudent2() {
	s1 := student2.Student{}
	s1.Init("John", 25, "cs")
	s1.SayHi()

	s1.Major = "finance"
	s1.SayHi()

	//试图修改私有属性
	//Error: s1.age undefined
	//s1.age = 28

	// OK: s1.GetAge()
	fmt.Printf("age is %d\n", s1.GetAge())

	//试图调用私有方法
	// Error：s1.getAge undefined
	//fmt.Printf("age is %d\n" , s1.getAge())

}

func main() {
	TestStudent1()
	//TestStudent2()

}
