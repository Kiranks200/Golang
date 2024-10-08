package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Kiranks200/mangoApi/router"
)

func main() {
	fmt.Println("MongoDB ATLAS")
	r := router.Router()
	fmt.Println("Server is running....")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("listing at port 4000..")
}
