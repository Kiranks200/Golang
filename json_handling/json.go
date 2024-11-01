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
	DecodeJSon()
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

func DecodeJSon() {
	josonDataFromWeb := []byte(`
	{
                "coursename": "DSA",
                "Price": 2000,
                "website": "edex",
				"tags":["interview","Java"]
				}
				`)

	var Newcourse course
	Checvalidjson := json.Valid(josonDataFromWeb)
	if Checvalidjson {
		fmt.Println("Is my Json Decodeing valid : ", Checvalidjson)
		json.Unmarshal(josonDataFromWeb, &Newcourse)
		fmt.Printf("%#v\n", Newcourse)
	} else {
		fmt.Println("Invalid json")
	}

	//add data to the key value pair
	var myonlineData map[string]interface{} //Key is string but value can be anything as we are interacting with the web
	json.Unmarshal(josonDataFromWeb, &myonlineData)
	fmt.Printf("%#v\n", myonlineData)

}
