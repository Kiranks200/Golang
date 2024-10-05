package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://localhost:9000/learn?couresname=reactjs&paymentid=kjfdihai22&paymentmode=upi"

func main() {
	fmt.Println("URL handling")
	fmt.Println(myurl)
	result, _ := url.Parse(myurl)

	fmt.Println("scheme:", result.Scheme)
	fmt.Println("host :", result.Host)
	fmt.Println("path :", result.Path)
	fmt.Println("port:", result.Port())

	fmt.Println("Raw-query :", result.RawQuery) // raw query is the query string
	fmt.Println("Query :", result.Query())      // query is the map of query string

	for key, val := range result.Query() {
		fmt.Println(key, val)
	}
	partsofurl := &url.URL{
		Scheme:   "https",
		Host:     "localhost:9000",
		Path:     "tut",
		RawQuery: "user=john_Doe",
	}

	anotherutl := partsofurl.String()
	fmt.Println(anotherutl)
}
