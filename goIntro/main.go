package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(value(2))
	fmt.Println(add(12, 12))

}

func value(x float64) float64 {
	return math.Pow(x, 100)
}
// function to add
func add(x, y int) int {
	return x + y
}
