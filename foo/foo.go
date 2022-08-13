package foo

import "strconv"

var Title = "Hell Title"

func Say(n int) string {

	if n % 5 == 0 {
		return "Bar"
	}

	if n % 3 == 0 {
		return "Foo"
	}


	return strconv.Itoa(n)
}