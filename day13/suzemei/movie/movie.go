package movie

import (
	"log"
)

const (
	REGULAR          = iota //0
	NEW_RELEASE             //1
	CHILDREN                // 2
)

type Movie struct {
	title string //片名
	//priceCode int    //价格代号
	price Pricer
}

func (m *Movie) Init(title string, priceCode int) {
	m.title = title
	m.SetPriceCode(priceCode)
}

func (m Movie) GetTitle() string {
	return m.title
}

func (m Movie) GetPriceCode() int {
	return m.price.GetPriceCode()
}

func (m *Movie) SetPriceCode(priceCode int) {
	switch priceCode {
	case REGULAR:
		m.price = new(RegularPrice)
	case CHILDREN:
		m.price = new(ChildrenPrice)
	case NEW_RELEASE:
		m.price = new(ChildrenPrice)
	default:
		log.Fatal("Incorrect price code")
	}

}

func (m Movie) GetCharge(daysRented int) float64 {
	return m.price.GetCharge(daysRented)
}

//func (m Movie) GetCharge(daysRented int) float64 {
//	result := 0.0
//	switch m.GetPriceCode() {
//	case REGULAR:
//		result += 2
//		if daysRented > 2 {
//			result += float64(daysRented-2.0) * float64(1.5)
//		}
//	case NEW_RELEASE:
//		result += float64(daysRented) * float64(3)
//	case CHILDREN:
//		result += 1.5
//		if daysRented > 3 {
//			result += float64(daysRented-3) * float64(1.5)
//		}
//	}
//	return result
//}

func (m Movie) GetFrequentRenterPoints(daysRented int) int {
	if m.GetPriceCode() == NEW_RELEASE && daysRented > 1 {
		return 2
	} else {
		return 1
	}
}
