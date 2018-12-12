package main

import "fmt"

/*
封装:

学生有姓名、年龄和专业等属性

Golang通过struct定义类的属性，通过在func定义时传入类对象的方式定义类的方法，其中属性和方法的公有/私有属性是通过首字母的大小写决定的。

*/

type Student struct {
	name  string
	age   int
	major string
}

// 如果我们要给Student类定义一个“构造函数”，我们希望的是这个函数的入参可以被赋值到Student的成员内，
// 则该“构造函数”应该使用指针类型对象定义：
func (s *Student) Init(name string, age int, major string) {
	s.name = name
	s.age = age
	s.major = major
}

func (s Student) SayHi() {
	fmt.Printf("Hi, I am  %s aged %d, and my major is %s\n", s.name, s.age, s.major)
}

func main() {
	s := Student{}
	s.Init("John", 25, "cs")
	s.SayHi()
	fmt.Printf("name is %s\n", s.name)
}