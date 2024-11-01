package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://pandas.pydata.org/"

func main() {
	fmt.Println("Read webpage")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("response of type %T\n", response)
	//fmt.Println(response) //gives response about status code,content-type,date,last_modified,server-timing etc

	defer response.Body.Close()

	databytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("content of the website is : ", string(databytes)) //prints all the webcontent(html)

}
