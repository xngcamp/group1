package main

import (
	"fmt"
)

/*
与C++、Java等完整支持面向对象的语言不同，Golang没有显式的继承，而是通过组合实现继承

定义一个基类Person，提供姓名和年龄两个属性，以及SayHi一个方法（Init类似于构造函数）

*/
type Person struct {
	name string
	age  int
}

/*
如果函数需要更新一个变量，或者如果一个实参太大而我们希望避免复制整个实参，必须使用指针来传递变量的地址
*/
func (p *Person) Init(name string, age int) {
	p.name = name
	p.age = age
}

func (p Person) SayHi() {
	fmt.Printf("Hi, I am %s, %d years old.\n", p.name, p.age)
}

// 通过组合的方式继承这个基类，实现Employee子类
// Employee组合了Person这个成员，除此之外它还拥有自己的成员company，即所属公司，
type Employee struct {
	Person
	company string
}

func (e *Employee) Init(name string, age int, company string) {
	e.Person.Init(name, age)
	e.company = company
}

// 雇员除了是一个Person之外，还需要工作，因此我们定义了Work这个方法
func (e Employee) Work() {
	fmt.Printf("I'm working at %s.\n", e.company)
}

func main() {
	p := Person{}
	p.Init("Tom", 22)
	p.SayHi()

	e := Employee{}
	e.Init("Jerry", 30, "ABC")
	e.SayHi()
	e.Work()
}
