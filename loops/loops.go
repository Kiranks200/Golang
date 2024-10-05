package main

import (
	"fmt"
)

func main() {
	fmt.Println("loops of golang")
	days := []string{"sun", "mon", "tue", "wed", "thu", "fri", "sat"}

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for i := range days {
		fmt.Println(i)
	}

	for index, day := range days {
		fmt.Println(index, day)
	}

	rougr := 1
	for rougr < 10 {
		if rougr == 2 {
			goto loc
		}
		if rougr == 5 {
			rougr++
			continue
		}
		fmt.Println(rougr)
		rougr++
	}

loc:
	fmt.Print("goto location : execution stopped at rougr==2")
}
