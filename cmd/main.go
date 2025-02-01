package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mostafizur-raahman/go-with-gone/internal/handlers"
	"github.com/mostafizur-raahman/go-with-gone/internal/repositories"
	"github.com/mostafizur-raahman/go-with-gone/internal/services"
)

func main() {
	// connect databae in future
	fmt.Println("Connect with db unsuccesfulyy... :) ")

	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	router := mux.NewRouter()

	router.HandleFunc("/user/create", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/users", userHandler.Users).Methods("GET")

	http.ListenAndServe(":8080", router)
}
