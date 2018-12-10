package exam

const PI float64 = 3.14159265358

type Triangle struct {
	Edge float64
	Height float64
}

type Circle struct {
	Radis float64
}

type Rectangle struct {
	Width float64
	Height float64
} 

func (t Triangle) Area() float64 {
	if t.Height < 0 || t.Edge < 0 {
		return -1
	}
	return 0.5*t.Edge*t.Height
}

func (c Circle) Area() float64 {
	if c.Radis < 0 {
		return -1
	}

	return PI*c.Radis*c.Radis
}

func (r Rectangle) Area() float64 {
	if r.Height < 0 || r.Width < 0 {
		return -1
	}
	return r.Width*r.Height
}
