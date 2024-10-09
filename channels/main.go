package main

import (
	"fmt"
	//"sync"
)

// func main() {
// 	fmt.Println("Channels in golang")

// 	mychannel := make(chan int, 2)
// 	wg := &sync.WaitGroup{}

// 	wg.Add(2)
// 	go func(ch <-chan int, wg *sync.WaitGroup) {////chan<- receive only
// 		val, isChannelOpen := <-mychannel
// 		//fmt.Println(<-mychannel)
// 		fmt.Println(isChannelOpen)
// 		fmt.Println(val)
// 		wg.Done()
// 	}(mychannel, wg)

// 	go func(ch chan<- int, wg *sync.WaitGroup) { //chan<- send only
// 		mychannel <- 56
// 		mychannel <- 22
// 		close(mychannel)
// 		wg.Done()
// 	}(mychannel, wg)

// 	wg.Wait()
// }

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
