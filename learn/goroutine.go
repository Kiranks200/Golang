package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{}

var waitgo sync.WaitGroup //pointers
var mut sync.Mutex //pointer

func main() {
	// go greeter("hello") //In standard we use sync package
	// greeter("world")
	websitelist := []string{
		"https://www.google.com",
		"https://www.bing.com",
		"https://github.com",
		"https://fb.com",
		"https://go.dev"}

	for _, web := range websitelist {
		go getStatusCode(web)
		waitgo.Add(1) //Add(1) because only one call happening
	}
	waitgo.Wait() //wait for all routines to commplete and is usually written at the last of the main function.
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 4; i++ {
// 		time.Sleep(3 * time.Millisecond) //wait to run greeter("hello")
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer waitgo.Done() //signal to indicate the completion of the routines

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Error in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}
