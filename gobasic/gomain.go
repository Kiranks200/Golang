package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	fmt.Println("time in golang")
	ptime := time.Now()
	fmt.Println(ptime) //output : 2024-09-26 12:41:11.8500648

	fmt.Println(ptime.Format("01-02-2006"))                 // output :09-26-2024
	fmt.Println(ptime.Format("01-02-2006 15:04:05 Monday")) // output:09-26-2024 12:45:36 Thursday

	// GOOS="windows" go build  - to build/convert the main file to windows exe file
	// GOOS="linux" go build - to build/convert the main file which can be run on the linux os

	//Pointers
	var a int = 10
	var b *int = &a //b is a pointer to a
	fmt.Println(b)  //output :0xc0000a8000
	fmt.Println(*b) //output :10
	*b = 20         //output :20
	fmt.Println(a)  //output :20

	// &->referance ,*->for the actual value of the pointer

	var listMix = []string{"wef", "23", "qreq", "erwer"}
	listMix = append(listMix[1:3])
	fmt.Println(listMix)

	score := make([]int, 4)
	score[0] = 66
	score[2] = 44
	score[1] = 89
	score[3] = 23
	fmt.Println(score)
	//no more elements can be added
	//but see this exaple below
	score = append(score, 34, 24234)
	fmt.Println(score)

	fmt.Println(sort.IntsAreSorted(score))
	sort.Ints(score)
	fmt.Println("sorted scores", score)

	index := 2
	score = append(score[:index], score[index+1:]...)
	//variadic parameter-It will accept zero or more string arguments, and reference them as a slice.
	//splitting the elements of the second slice so that can be added to first slice as individual elements by the append method
	fmt.Println("after removing 3rd element", score)

	// maps-hashtable-(key:val)
	languages := make(map[string]string)
	languages["python"] = "scripting"
	languages["java"] = "object oriented"
	languages["c"] = "procedural"
	languages["c++"] = "object oriented"

	fmt.Println("maps : ", languages)
	fmt.Println("c is mapped as : ", languages["c"])
	delete(languages, "c") //can be used to delete in maps and slices
	fmt.Println("map : ", languages)

	for key, value := range languages {
		fmt.Println(key, value)
	}

}
