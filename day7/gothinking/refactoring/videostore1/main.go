package main

import (
	"fmt"

	"gothinking/refactoring/videostore1/customer"
	"gothinking/refactoring/videostore1/movie"
	"gothinking/refactoring/videostore1/rental"
)

/*
目标：
给一个录影带商店编写一个计算和打印顾客租赁录影带的收费和积分详单


需求：
1. 计算顾客租赁的电影清单和每部电影的租赁费用
2. 计算总的租赁费用
3. 根据电影类型，计算用户积分


有三种类型的电影：普通电影(类型0)，儿童电影(类型2)和新发行的电影(类型1)
收费规则：
1.普通电影：前2天共2元，超过2天后，每天1.5元
2.新发行的电影：每天3元
3.儿童电影：前3天共1.5元，超过3天后，每天1.5元

积分规则：
每租一部电影得1积分
新上市电影：租期大于1天，额外得1积分

结果示例：

Rental Record for Tom:
	1. The Godfather	6.50
	2. Fast & Furious	3.00
Amount owned is 9.50
Tom earned 2 frequent renter points.

备注：
videostore1.pos  使用 https://www.processon.com 浏览

*/

func main() {
	//aCustomer := new(customer.Customer)
	var aCustomer customer.Customer
	aCustomer.Init("Tom")

	movie1 := movie.Movie{Title: "Star Wars", PriceCode: movie.REGULAR}
	var rental1 rental.Rental
	rental1.Init(movie1, 3)
	aCustomer.AddRental(rental1)

	movie2 := movie.Movie{Title: "The Godfather Part II", PriceCode: movie.NEW_RELEASE}

	var rental2 rental.Rental
	rental2.Init(movie2, 1)
	aCustomer.AddRental(rental2)

	/*rental1 := rental.Rental{AMovie: movie.Movie{Title: "The Godfather", PriceCode: movie.REGULAR}, DaysRented: 5}
	aCustomer.AddRental(rental1)

	rental2 := rental.Rental{AMovie: movie.Movie{Title: "Fast & Furious", PriceCode: movie.NEW_RELEASE}, DaysRented: 1}
	aCustomer.AddRental(rental2)*/

	result := aCustomer.Statement()
	fmt.Println(result)
}
