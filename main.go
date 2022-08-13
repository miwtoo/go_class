package main

import "fmt"

func main()  {
	println("Hello, World!")

	fmt.Println("squareArea: ", squareArea(2))

	a, b := swap(1, 2)
	fmt.Println("swap(1, 2): ", a, b)
}

func swap(a,b int) (int, int) {
	return b, a
}

func squareArea(a float64) float64  {
	return a * a
}