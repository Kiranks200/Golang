package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var a = "Hey"
	fmt.Print(a, "\n")

	reader := bufio.NewReader(os.Stdin)
	//comma ok | error syntax
	msg, _ := reader.ReadString('\n')
	fmt.Println("Learing Golang as a ", msg)

	input, _ := reader.ReadString('\n')
	fmt.Println("Rate the Golang between 1 to 5 ", input)
	inputconvert, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		if inputconvert > 5 {
			fmt.Println("invlid input")
		} else {
			fmt.Println("input converted to decimal", inputconvert)
		}
	}

	fmt.Println(value(2))
	fmt.Println(add(12, 23))

	//ep-18

}

func value(x float64) float64 {
	return math.Pow(x, 100)
}

// function to add
func add(x, y int) int {
	return x + y
}
