package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// connect databae in future
	fmt.Println("Connect with db unsuccesfulyy... :) ")

	router := mux.NewRouter("/create/user")

	http.ListenAndServe(":8080", router)
}
