package couple

func Couple(str string) []string {
	var r []string
	s := []rune(str)

	for s = append(s, []rune("*")...); len(s) > 1; s = s[2:]{
		r = append(r, string(s[:2]))
	}

	return r
}