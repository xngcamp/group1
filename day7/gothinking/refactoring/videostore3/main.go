package main

import (
	"fmt"
	"gothinking/refactoring/videostore3/customer"
	"gothinking/refactoring/videostore3/movie"
	"gothinking/refactoring/videostore3/rental"
)

/*

重构:

1. 提取出常客积分计算代码
*  常客的积分代码计算用到的信息和Customer也没有关系
*  把这部分计算抽取到Rental中



对比两个序列图的变化

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
