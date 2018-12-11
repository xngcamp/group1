package movie

type ReMovie struct {
	Title     string //片名
PriceCode int    //价格代号
	//Movie
}

func (m ReMovie) GetCharge(daysRented int) float64 {
	result := 0.0
	result += 2
	if daysRented > 2 {
		result += float64(daysRented-2.0) * float64(1.5)
	}
	return  result

}

func (m ReMovie) GetPoint(daysRented int) int {
		return 1
}
