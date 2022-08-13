package couple

func Couple(s string) []string {
	var r []string

	for s += "*"; len(s) > 1; s = s[2:]{
		r = append(r, s[:2])
	}

	return r
}