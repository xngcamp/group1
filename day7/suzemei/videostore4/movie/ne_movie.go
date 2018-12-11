package movie


type Ne_Movie struct {
	//Movie
	Title     string //片名
	PriceCode int    //价格代号
}

func (m Ne_Movie) GetCharge(daysRented int) float64 {
	result := 0.0
	result += float64(daysRented) * float64(3)
	return  result

}

func (m Ne_Movie) GetPoint(daysRented int) int {
		if daysRented > 1 {
			return 2
		} else {
			return 1
		}
	}