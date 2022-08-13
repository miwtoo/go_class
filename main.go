package main

import (
	"fmt"
	"os"

	"github.com/miwtoo/go_class/foo" // suggest
	// any_name "github.com/miwtoo/go_class/foo"
	// _ "github.com/miwtoo/go_class/foo"
	// . "github.com/miwtoo/go_class/foo"

	"github.com/miwtoo/go_class/couple"
)

var name string = "Miwtoo"

func main() {
	// NAME=miwtoo go run main.go 
	n := os.Getenv("NAME")
	if n != "" {
		name = n
	}
	fmt.Println("Hello, ", foo.Title , name)

	fmt.Println(couple.Couple("abcdef"))
	fmt.Println(couple.Couple("abcdefg"))

}

// func main()  {
// 	println("Hello, World!")

// 	fmt.Println("squareArea: ", squareArea(2))

// 	a, b := swap(1, 2)
// 	fmt.Println("swap(1, 2): ", a, b)

// 	if ok := isCorrect(); ok {
// 		println("it' correct")
// 	}

// 	println("power(2, 3): ", power(2, 3) )
// }

// func power(b, x int) int {
// 	r := 1
// 	for i := 0; i < x; i++ {
// 		r *= b
// 	}
// 	return r
// }

// func isCorrect() bool {
// 	return true
// }

// func swap(a,b int) (int, int) {
// 	return b, a
// }

// func squareArea(a float64) float64  {
// 	return a * a
// }