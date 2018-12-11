package movie

type NewIssueMovie struct {
	Param
}

func (n NewIssueMovie) GetCost(days int) float64 {
	return float64(days) * 3
}

func (n NewIssueMovie) GetPoint(days int) int {
	if days > 1 {
		return 2
	}
	return 1
}

func (n NewIssueMovie) GetTitle() string {
	return n.Title
}