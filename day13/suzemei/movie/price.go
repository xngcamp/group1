package movie

// 接口
type Pricer interface {
	GetPriceCode() int
	GetCharge(daysRented int) float64
}
