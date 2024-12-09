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

	fmt.Println("Rate the Golang between 1 to 5 ")
	input, _ := reader.ReadString('\n')
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

	defer fmt.Println(value(2)) // defer - to execute at last(LIFO).
	defer fmt.Println(adder(12, 12))
	params := Numbers{2, 3}
	fmt.Println("Methods: value in x,y,addition()", params.x, params.y, params.addition())
	fmt.Println(proadder(2, 3, 4, 5, 6, 7, 8, 8, 6, 5, 4, 4, 4))
}

func value(x float64) float64 { //declaring a function
	return math.Pow(x, 100)
}

func adder(x, y int) int { //function to add two numbers
	return x + y
}

type Numbers struct {
	x, y int
}

func (n Numbers) addition() int { // Method with Pointer Receivers - func addition(n Numbers) int {#code}.
	return n.x + n.y
}

func proadder(values ...int) int { // ... - we can give more paramertric values as input
	total := 0

	for _, a := range values {
		total += a
	}
	return total
}
