package movie

type NewReleasePrice struct {
}

func (n NewReleasePrice) GetPriceCode() int {
	return NEW_RELEASE
}

func (n NewReleasePrice) GetCharge(daysRented int) float64 {
	return float64(daysRented * 3)
}

func (n NewReleasePrice) GetFrequentRenterPoints(daysRented int) int {
	if n.GetPriceCode() == NEW_RELEASE && daysRented > 1 {
		return 2
	} else {
		return 1
	}
}
