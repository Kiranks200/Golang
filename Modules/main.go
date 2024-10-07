package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux" //it is a midddleware which defines routes for the web applications
	//github.com/go-chi/chi/middleware ww can also use this gor middleware
)

func main() {
	fmt.Println("mod in golang")
	greeter()

	r := mux.NewRouter()
	r.HandleFunc("/", ServerHome).Methods("Get")

	log.Fatal(http.ListenAndServe(":4000", r)) //localhost server
}

func greeter() {
	fmt.Println("Hey there mod user")
}

func ServerHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Code with Golang</h1>"))
}
