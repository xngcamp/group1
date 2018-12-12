package customer

import (
	"fmt"

	"gothinking/refactoring/videostore4/rental"
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

func (c *Customer) Init(name string) {
	c.name = name
}

func (c *Customer) AddRental(arg rental.Rental) {
	c.rentals = append(c.rentals, arg)
}

func (c Customer) GetName() string {
	return c.name
}

// 图3-2对应
func (c Customer) Statement1() string {
	var totalAmount float64      //消费总额
	var frequentRenterPoints int //常客积分
	var result = "Rental Record for " + c.GetName() + ":\n"
	for index, r := range c.rentals {

		//thisAmount := c.amountFor(r)
		//不需要再调用amountFor， 直接调用Rental的GetCharge()
		thisAmount := r.GetCharge()

		// add frequent renter points (累计常客积分)
		frequentRenterPoints = r.GetFrequentRenterPoints()

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

// 图3-3对应
func (c Customer) Statement() string {
	var totalAmount float64 //消费总额
	var result = "Rental Record for  " + c.GetName() + "\n"
	for _, r := range c.rentals {

		//thisAmount := c.amountFor(r)
		//不需要再调用amountFor， 直接调用Rental的GetCharge()
		thisAmount := r.GetCharge()

		//show figures for this rental (显示此次租赁的数据)
		thisResult := fmt.Sprintf("\t%s\t%.2f\n", r.GetMovie().GetTitle(), thisAmount)
		result += thisResult
		totalAmount += thisAmount
	}
	// add footer lines (添加页脚信息)
	result += fmt.Sprintf("Amount owned is %.2f\n", totalAmount)
	result += fmt.Sprintf("You earned %d frequent renter points.\n", c.GetTotalFrequentRenterPoints())
	return result
}

func (c Customer) GetTotalFrequentRenterPoints() int {
	result := 0
	for _, r := range c.rentals {
		// add frequent renter points (累计常客积分)
		result += r.GetFrequentRenterPoints()
	}
	return result
}
