package movie

type Movie interface {
	GetCost(days int) float64
	GetPoint(days int) int
	GetTitle() string
}