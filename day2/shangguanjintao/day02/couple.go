package day02

/*
*找出配对的右括号
*/
func couple(s string, m int) int {
	arr := []byte(s)
	flag := 1
	var c int
	for i := m + 1; flag >= 0 && i < len(arr); i++ {
		if arr[i] == '(' {
			flag += 1
		} else if arr[i] == ')' {
			flag -= 1
		} else if flag == 0 {
			c = i
		} else {
			return 0
		}
	}
	return c
}
