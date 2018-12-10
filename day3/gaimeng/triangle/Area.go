package triangle
type Triangle struct {
	Len float64
	Wid float64

}

//func (a Triangle) Area() float64 {
//	panic("implement me")
//}

func (t Triangle)Area() float64{

	return  (t.Len*t.Wid)/2

}