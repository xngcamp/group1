package circle

type Circle struct {
	R float64

}

func (c Circle)Area() float64 {
	const PI=3.115926

	return PI*c.R*c.R

	}
