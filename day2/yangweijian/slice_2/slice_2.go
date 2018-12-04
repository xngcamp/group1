package slice_2

import "strings"

func IsValidISBN(str string) bool {
	strs := strings.Split(str, "-")
	length := 0
	for _, s := range strs {
		length += len(s)
	}
	if length != 10 {
		return false
	}

	sum, multiplier := 0, 10
	arr := []rune(str)

	for i, val := range arr {
		if val == '-' {
			continue
		}
		switch {
			case val >= '0' && val <= '9':
				sum += int(val - '0') * multiplier
			case val == 'X' && i == len(arr) - 1:
				sum += 10
			default:
				return false
		}
		multiplier --
	}
	return sum % 11 == 0
}
