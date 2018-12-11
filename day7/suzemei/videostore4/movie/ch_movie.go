package movie

type ChMovie struct {
	//Movie
	Title     string //片名
	PriceCode int    //价格代号
}

func (m ChMovie) GetCharge(daysRented int) float64 {
	result := 0.0
	result += 1.5
			if daysRented > 3 {
				result += float64(daysRented-3) * float64(1.5)
			}
	return  result

}



func (m ChMovie) GetPoint(daysRented int) int {
	return 1
}
