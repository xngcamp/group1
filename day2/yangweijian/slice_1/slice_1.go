package slice_1

func Solve(str string, left int) int {
	arr := []rune(str)
	if arr[left] != '(' {
		return -1
	}

	right := 1
	var res int

	for res = left + 1; res < len(arr); res ++ {
		if arr[res] == ')' {
			right --
		} else if arr[res] == '(' {
			right ++
		}
		if right == 0 {
			break
		}
	}
	if right == 0 {
		return res
	} else {
		return -1
	}
}