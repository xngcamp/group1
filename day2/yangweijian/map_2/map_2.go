package map_2

func Translate(str string) string {
	if len(str) % 3 != 0 {
		return ""
	}
	codons := []byte(str)

	var res string
	for i := 0; i < len(str) / 3; i ++ {
		name := ProteinName[string(codons[i * 3: (i + 1) * 3])]
		if name == "STOP" {
			break
		}
		res += name
	}
	return res
}
