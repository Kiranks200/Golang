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
	fmt.Printf("%v \n", kiran)  //{Kiran kiran@gmail true 19}
	fmt.Printf("%+v \n", kiran) //{Name:Kiran Email:kiran@gmail Stat:true Age:19}
	kiran.Greetuser()
	fmt.Println(kiran.Newemail())
}

// methods in golang

func (u User) Greetuser() { // method name first letter capital as we have to export it
	fmt.Println("Hi", u.Name)

}

func (u User) Newemail() string {   // func (u *User) Newemail() string{} - pointers can be used to change the original value.
	u.Email = "test@gmail.com"     // func (u User) Newemail() string{} - this method only creates a copy doesnot change the original value
	return (u.Email)          
}