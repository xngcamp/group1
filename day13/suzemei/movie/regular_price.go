package movie

type RegularPrice struct {
}

func (r RegularPrice) GetPriceCode() int {
	return REGULAR
}

func (r RegularPrice) GetCharge(daysRented int) float64 {
	result := 2.0
	if daysRented > 2 {
		result += float64(daysRented-2) * 1.5
	}
	return result
}

func (r RegularPrice) GetFrequentRenterPoints(daysRented int) int {
	return 1
}
