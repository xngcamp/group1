package customer

import (
	"fmt"

	"gothinking/refactoring/videostore2/movie"
	"gothinking/refactoring/videostore2/rental"
)

/*
重构：
1. 把计算每一个租金的方法充Statement中抽取出来，形成新的amountFor()方法
2. amountFor() 参数更合理的命名
3. amountFor()中计算租金，只用了了rental，并没有用到customer， 放在customer中不合理，继续抽取
3. struct中作用域, 不暴露内部信息
*/

type Customer struct {
	name    string
	rentals []rental.Rental
}

// 注意： Init()的receiver要用指针 *Customer，考虑一下为什么？
func (c *Customer) Init(name string) {
	c.name = name
}

func (c *Customer) AddRental(arg rental.Rental) {
	c.rentals = append(c.rentals, arg)
}

func (c Customer) GetName() string {
	return c.name
}

func (c Customer) Statement() string {
	var totalAmount float64      //消费总额
	var frequentRenterPoints int //常客积分
	var result = "Rental Record for  " + c.GetName() + ":\n"
	for index, r := range c.rentals {

		//thisAmount := c.amountFor(r)
		//不需要再调用amountFor， 直接调用Rental的GetCharge()
		thisAmount := r.GetCharge()

		// add frequent renter points (累计常客积分)
		frequentRenterPoints++

		// add bonus for a two day new release rental
		if r.GetMovie().GetPriceCode() == movie.NEW_RELEASE && r.GetDaysRented() > 1 {
			frequentRenterPoints++
		}

		//show figures for this rental (显示此次租赁的数据)
		thisResult := fmt.Sprintf("\t%d. %s\t%.2f\n", index+1, r.GetMovie().GetTitle(), thisAmount)
		result += thisResult
		totalAmount += thisAmount
	}
	// add footer lines (添加页脚信息)
	result += fmt.Sprintf("Amount owned is %.2f\n", totalAmount)
	result += fmt.Sprintf("%s earned %d frequent renter points.\n", c.GetName(), frequentRenterPoints)
	return result
}

/*
// 第一次重构后的amountFor()
func (c Customer)amountFor(aRental rental.Rental) float64 {
	//
	thisAmount := 0.0
	// determine amounts for each record ,  各种片子价格不同
	aMovie := aRental.aMovie
	switch aMovie.PriceCode {
	case movie.REGULAR:
		thisAmount += 2
		if aRental.DaysRented > 2 {
			thisAmount += float64(aRental.DaysRented-2.0) * float64(1.5)
		}
	case movie.NEW_RELEASE:
		thisAmount += float64(aRental.DaysRented) * float64(3)
	case movie.CHILDRES:
		thisAmount += 1.5
		if aRental.DaysRented > 3 {
			thisAmount += float64(aRental.DaysRented-3) * float64(1.5)
		}
	}
	return thisAmount
}
*/

func (c Customer) amountFor(aRental rental.Rental) float64 {
	return aRental.GetCharge()
}
