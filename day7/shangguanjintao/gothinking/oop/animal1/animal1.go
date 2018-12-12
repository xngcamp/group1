package main

import "fmt"

/*
多态：
基类指针可以指向任意派生类的对象，并在运行时绑定最终调用的方法的过程被称为多态。
多态是运行时特性，而继承则是编译时特征，
也就是说，继承关系在编译时就已经确定了，而多态则可以实现运行时的动态绑定。

基类指针可以指向任意派生类的对象，并在运行时动态绑定最终使用的方法。
这里指针是广义上的概念，在C++中是真实的指针，在Java和Golang里面，则可以是一个接口类型的对象

小狗和小鸟都是动物，它们都会移动，也都会叫唤。我们把它们共同的方法提炼出来定义一个抽象的接口

*/
type Animal interface {
	Move()
	Shout()
}

//虽然小狗和小鸟都会移动，但小狗是用四条腿爬行，小鸟是用翅膀飞行，虽然它们都会叫唤，但是叫唤的方式也不一样
type Dog struct {
}

func (dog Dog) Move() {
	fmt.Println("A dog moves with its legs.")
}

func (dog Dog) Shout() {
	fmt.Println("wang wang wang")
}

type Bird struct {
}

func (bird Bird) Move() {
	fmt.Println("A bird fly with its wings.")
}

func (bird Bird) Shout() {
	fmt.Println("ji ji, ji ji")
}

//通过定义抽象的接口，以及实现接口方法的具体类型的方式，Golang实现了运行时的动态绑定，这就是所谓的抽象与多态。
func main() {
	var animal Animal
	animal = Dog{}
	animal.Move()
	animal.Shout()

	animal = Bird{}
	animal.Move()
	animal.Shout()

}
