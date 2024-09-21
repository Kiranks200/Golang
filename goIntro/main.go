package main

import ("fmt"
	"math")

func main() {
	fmt.Println(value(2))

}

func value(x float64) float64{
	return math.Pow(x, 100)
}