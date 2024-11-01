package main

import (
	"fmt"
	//"net"
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
	partsofurl := &url.URL{ //creating url
		Scheme:   "https",
		Host:     "localhost:9000",
		Path:     "tutorial",
		RawQuery: "user=john_Doe",
	}

	anotherutl := partsofurl.String()
	fmt.Println(anotherutl)
}

// func main() {

// 	s := "postgres://user:pass@host.com:5432/path?k=v#f"

// 	u, err := url.Parse(s)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(u.Scheme)

// 	fmt.Println(u.User)
// 	fmt.Println(u.User.Username())
// 	p, _ := u.User.Password()
// 	fmt.Println(p)

// 	fmt.Println(u.Host)
// 	host, port, _ := net.SplitHostPort(u.Host)
// 	fmt.Println(host)
// 	fmt.Println(port)

// 	fmt.Println(u.Path)
// 	fmt.Println(u.Fragment)

// 	fmt.Println(u.RawQuery)
// 	m, _ := url.ParseQuery(u.RawQuery)
// 	fmt.Println(m)
// 	fmt.Println(m["k"][0])
// }
