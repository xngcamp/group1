package main

import (
	"fmt"
	"gothinking/refactoring/videostore4/customer"
	"gothinking/refactoring/videostore4/movie"
	"gothinking/refactoring/videostore4/rental"
)

/*

重构：
1. 用多态代替价格条件逻辑代码
* 基于其它对象的属性做switch-case不是的好的做法
* 如果你必须用switch-case, 应该用自己类/Struct中的数据，而不是其它的数据
* GetCharge计算租金，用到的也不是Rental中数据，而是Movie的数据，因此它应该被移到Movie中

2. 为什么倾向把租赁的时长daysRented传递给Movie，而不是把Movie类型传递给Rental？
* 因为可能的变化是增加新的电影类型，类型信息通常是易变的
* 如果更改了电影类型，希望最小的扩散效应，因此在Movie内计算租金

3. 对租金逻辑进行了移动到Movie的处理，对积分也做同样处理

4. 看类图4-1和4-2的变化


*/

func main() {
	//aCustomer := new(customer.Customer)
	var aCustomer customer.Customer
	aCustomer.Init("Tom")

	var rental1 rental.Rental
	movie1 := movie.Movie{Title: "Star Wars", PriceCode: movie.REGULAR}
	rental1.Init(movie1, 3)
	aCustomer.AddRental(rental1)

	var rental2 rental.Rental
	movie2 := movie.Movie{Title: "The Godfather Part II", PriceCode: movie.NEW_RELEASE}
	rental2.Init(movie2, 1)
	aCustomer.AddRental(rental2)

	var rental3 rental.Rental
	movie3 := movie.Movie{Title: "Casablanca", PriceCode: movie.CHILDRES}
	rental3.Init(movie3, 7)
	aCustomer.AddRental(rental3)

	result := aCustomer.Statement()
	fmt.Println(result)
}
