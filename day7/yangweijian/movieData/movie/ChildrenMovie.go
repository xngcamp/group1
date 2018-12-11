package movie

type ChildrenMovie struct {
	Param
}

func (c ChildrenMovie) GetCost(days int) float64 {
	if days <= 3 {
		return 1.5
	}
	return 1.5 + float64(days) * 1.5
}

func (c ChildrenMovie) GetPoint(days int) int {
	return 1
}

func (c ChildrenMovie) GetTitle() string {
	return c.Title
}