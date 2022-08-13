package main

import "fmt"

func main()  {
	println("Hello, World!")

	fmt.Println("squareArea: ", squareArea(2))
}

func squareArea(a float64) float64  {
	return a * a
}