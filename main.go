package main

import (
	"SE_School/controllers"
	"SE_School/dal"
	"SE_School/middleware"
	"SE_School/rate_providers"
	"SE_School/services"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	//Dependency Injection
	userServ := services.UserService{Repo: &dal.FileRepository{FileLocation: "users.data"}}
	controllers.UserServ = &userServ
	controllers.BtcServ = &services.BtcService{RateProvider: &rate_providers.CoindeskRateProvider{}}
	controllers.AuthServ = &services.AuthService{}

	router := mux.NewRouter()
	router.HandleFunc("/user/create", controllers.Create).Methods("POST")
	router.HandleFunc("/user/login", controllers.Login).Methods("POST")
	router.HandleFunc("/btcRate", controllers.Rate).Methods("GET")
	http.Handle("/", router)

	router.Use(middleware.JwtMiddleware)

	port := os.Getenv("PORT")

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		fmt.Print(err)
	}
}
