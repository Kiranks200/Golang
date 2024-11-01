package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Get,Post,PostForm handling for index.js using Golang")

	Perform_getrequests()
	// Perform_Jsonpost()
	// Perform_postFormRequst()

}

//run the index.js in webserver folder

func Perform_getrequests() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status code :", response.StatusCode)
	fmt.Println("content length :", response.ContentLength)

	var response_string strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := response_string.Write(content)

	fmt.Println("byteCount is :", byteCount)
	fmt.Println("response using strings.Builder : ", response_string.String())
	fmt.Println("response using io.Readall : ", string(content))

}

func Perform_Jsonpost() {
	const myurl = "http://localhost:8000/post"

	//fake json payload
	requestBody := strings.NewReader(`
	{
		"Cousename":"reactjs",
		"amount":4000,
		"website":"udemy"
	}
	`)
	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println("content of post", string(content))

}

func Perform_postFormRequst() {
	const myurl = "http://localhost:8000/postform"

	//form-data

	data := url.Values{}
	data.Add("Firstname", "Kiran")
	data.Add("Lastname", "K S")
	data.Add("Email", "post_form@test.com")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))

}
