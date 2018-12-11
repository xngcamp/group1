package customer

import (
	"fmt"

	"gothinking/refactoring/videostore1/movie"
	"gothinking/refactoring/videostore1/rental"
)

type Customer struct {
	name    string
	rentals []rental.Rental
}

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
	var result = "Rental Record for " + c.GetName() + ":\n"
	for index, r := range c.rentals {

		//thisAmount := c.funcName(r)
		thisAmount := r.GetCharge()
		

		// add frequent renter points (累计常客积分)
		frequentRenterPoints++

		// add bonus for a two day new release rental
		if r.GetMovie().GetPriceCode() == movie.NEW_RELEASE && r.GetDaysRented() > 1 {
			frequentRenterPoints++
		}

		//show figures for this rental (显示此次租赁的数据)
		//thisResult := fmt.Sprintf("\t%d. %s\t%.2f\n", index+1, r.Title, thisAmount)
		thisResult := fmt.Sprintf("\t%d. %s\t%.2f\n", index+1, r.GetMovie().GetTitle(), thisAmount)
		result += thisResult
		totalAmount += thisAmount
	}
	// add footer lines (添加页脚信息)
	result += fmt.Sprintf("Amount owned is %.2f\n", totalAmount)
	result += fmt.Sprintf("%s earned %d frequent renter points.\n", c.GetName(), frequentRenterPoints)
	return result
}

/*func (c Customer) funcName(aRental rental.Rental) float64 {
	thisAmount := 0.0
	// determine amounts for each record ,  各种片子价格不同
	aMovie := aRental.AMovie
	switch aMovie.PriceCode {
	case movie.REGULAR:
		thisAmount += 2
		if aRental.DaysRented > 2 {
			thisAmount += float64(aRental.DaysRented-2) * float64(1.5)
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
}*/
