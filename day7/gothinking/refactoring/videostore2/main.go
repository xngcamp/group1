package main

import (
	"fmt"
	"gothinking/refactoring/videostore2/customer"
	"gothinking/refactoring/videostore2/movie"
	"gothinking/refactoring/videostore2/rental"
)

/*
重构：
1. 分解并重组Statement()方法
* 把一个长的方法分解成短的方法， 提取出amountFor方法
* 把相应的代码移到合适的struct中

2. 把amountFor方法中租金计算从Customer中提取到Rental中
* 因为amountFor中用到的信息是Rental中的，并没有用到Customer中的信息
* 提取出一个新方法GetCharge， 放到Rental中

3. 测试
* 做每一步重构后一定要及时测试
* 不要等积累很多修改后再做测试，不利于查找问题

4. 再看amountFor方法
* 发现amountFor方法只有一行，就是调用Rental的GetCharge方法
* 因此，Statement可以直接调用GetCharge方法

5. 回过头再看Statement方法
* 发现thisAmount变量是多余的
* 尽可能删除临时变量





Goland，对Statement中计算租金的代码右键:
Refactor->Extract->Method

提取新的方法要注意:
* 看方法中的局部变量，还有临时变量
* 提取新方法后看上下文的影响
* 对新方法中的变量命名等，做合理化调整



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

	result := aCustomer.Statement()
	fmt.Println(result)
}
