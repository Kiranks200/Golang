package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Cname    string `json:"coursename"` //something like alias
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("JSON handling using Golang")
	JsonEncoding()
}

func JsonEncoding() {
	Newcourse := []course{
		{"DSA", 2000, "edex", "alpha22", []string{"web_dev", "java"}},
		{"Java", 3000, "Coursera", "445566", nil},
		{"ML", 3000, "Udemy", "789654", []string{"python", "pandas", "Numpy"}},
	}

	//package the data as json

	finalJosn, err := json.MarshalIndent(Newcourse, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", (finalJosn))

}
