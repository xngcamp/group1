package movie

type RegularMovie struct {
	Param
}

func (r RegularMovie) GetCost(days int) float64 {
	if days <= 2 {
		return 2.0
	}
	return 2 + (float64(days) - 2) * 1.5
}

func (r RegularMovie) GetPoint(days int) int {
	return 1
}

func (r RegularMovie) GetTitle() string {
	return r.Title
}