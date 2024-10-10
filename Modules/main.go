package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux" //it is a midddleware which defines routes for the web applications
	
	//github.com/go-chi/chi/middleware we can also use this for middleware routing
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
	w.Write([]byte("<iframe src=https://www.youtube.com/embed/dQw4w9WgXcQ>click</iframe>")) //writing somthing onto the webpage

}
