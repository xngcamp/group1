package rectangle
type Rectangle struct {
	Len float64
	Wid float64

}


func (r Rectangle)Area() float64 {

	return  r.Len*r.Wid

}

