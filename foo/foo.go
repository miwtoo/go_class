package foo

import "strconv"

var Title = "Hell Title"

func Say(n int) string {
	if n == 10 {
		return "Bar"
	}

	if n == 9 {
		return "Foo"
	}

	if n == 8 {
		return strconv.Itoa(n)
	}

	if n == 7 {
		return strconv.Itoa(n)
	}

	if n == 6 {
		return "Foo"
	}

	if n == 5 {
		return "Bar"
	}

	if n == 4 {
		return strconv.Itoa(n)
	}

	if n == 3 {
		return "Foo"
	}

	if n == 2 {
		return strconv.Itoa(n)
	}

	return strconv.Itoa(n)
}