package core

func ReversedString(s string) string {
	charSlice := make([]string, len(s))
	p := -1
	// var output string = ""

	for _, char := range s {
		p++
		charSlice[p] = string(char)
	}

	o := ""
	// fmt.Println(charSlice)

	for p >= 0 {
		o = o + charSlice[p]
		p--
	}

	return o
}
