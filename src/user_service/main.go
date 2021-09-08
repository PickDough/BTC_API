package main

import (
	"User_Service/controllers"
	"User_Service/dal"
	"User_Service/message_broker"
	"User_Service/services"
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
	userServ := services.UserService{Repo: &dal.FileRepository{FileLocation: "user.data"}}
	controllers.UserServ = &userServ
	controllers.MsgBroker = message_broker.CreateMessageBroker(os.Getenv("RABBIT_MQ"), os.Getenv("EXCHANGE"))
	defer controllers.MsgBroker.Close()

	router := mux.NewRouter()
	router.HandleFunc("/user/create", controllers.Create).Methods("POST")
	router.HandleFunc("/user/login", controllers.Login).Methods("POST")
	http.Handle("/", router)

	port := os.Getenv("PORT")

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		fmt.Print(err)
	}
}
