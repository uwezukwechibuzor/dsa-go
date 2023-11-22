package main

// importing fmt package
import (
	"fmt"
)

// gets the power series of integer a and returns tuple of square of a
// and cube of a
func powerSeries(a int) (int, int) {
	return a * a, a * a * a
}

func powerSeries_(a int) (square int, cube int) {
	square = a * a
	cube = square * a
	return
}

// main method
func main() {
	var square int
	var cube int
	square, cube = powerSeries(3)
	fmt.Println("Square ", square, "Cube", cube)

	square, cube = powerSeries_(4)
	fmt.Println("Square ", square, "Cube", cube)
}
