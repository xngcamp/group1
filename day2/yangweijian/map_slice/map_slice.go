package map_slice

func Crossover(ns []int, xs []int,ys []int) ([]int, []int) {
	length := len(xs)
	r1 := make([]int, length, length)
	r2 := make([]int, length, length)
	nsIndex, resIndex := 0, 0
	xsAppendR1 := true

	for i := range xs {
		if nsIndex < len(ns) && i == ns[nsIndex] {
			xsAppendR1 = !xsAppendR1
			nsIndex ++
		}
		if xsAppendR1 {
			r1[resIndex] = xs[i]
			r2[resIndex] = ys[i]
		} else {
			r1[resIndex] = ys[i]
			r2[resIndex] = xs[i]
		}
		resIndex ++
	}
	return r1, r2
}
