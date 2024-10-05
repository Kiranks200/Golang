package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("File operations in golang")

	content := "inserting into a file"

	file, err := os.Create("./creating_file.txt")
	errchecknil(err)

	length, err := io.WriteString(file, content)
	errchecknil(err)

	fmt.Println("length is", length)

	defer file.Close()
	Readfile("./creating_file.txt")

}

func Readfile(filename string) {
	databyte, err := os.ReadFile(filename)
	errchecknil(err)
	fmt.Println("text inside the file in byte format :", databyte)
	fmt.Println("text inside the file is :", string(databyte))
}

func errchecknil(err error) {
	if err != nil {
		panic(err)
	}
}
