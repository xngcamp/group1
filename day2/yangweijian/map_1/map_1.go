package map_1

func NoDuplicates(str string) string {
	start, end := 0, 0
	var charMap map[byte]int
	arr := []byte(str)

	for i := 0; i < len(str); i ++ {
		charMap = make(map[byte]int)
		for j := i; j < len(str); j ++ {
			_, exist := charMap[arr[j]]
			if !exist {
				charMap[arr[j]] = 1
			} else {
				if j - i > end - start {
					start, end = i, j
				}
				break
			}
		}
	}
	return string(arr[start: end])
}
