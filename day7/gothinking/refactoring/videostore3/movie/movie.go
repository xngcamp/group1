package movie

const (
	REGULAR     = iota //0
	NEW_RELEASE        //1
	CHILDRES           // 2
)

type Movie struct {
	Title     string //片名
	PriceCode int    //价格代号
}

func (m Movie) GetTitle() string {
	return m.Title
}

func (m Movie) GetPriceCode() int {
	return m.PriceCode
}
