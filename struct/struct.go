package main

import (
	"fmt"
)

type User struct { //caps first letter as they are exported
	Name  string
	Email string
	Stat  bool
	Age   int
}

func main() {
	fmt.Println("Struct in golang")
	//no inheritance and no super or parent
	kiran := User{"Kiran", "kiran@gmail", true, 19}
	fmt.Println(kiran)
	fmt.Println(kiran.Email)
	fmt.Printf("%v \n", kiran)//{Kiran kiran@gmail true 19}
	fmt.Printf("%+v ", kiran)//{Name:Kiran Email:kiran@gmail Stat:true Age:19}

}
