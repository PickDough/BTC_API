package main

import (
	"BTC_Service/controllers"
	"BTC_Service/dal"
	"BTC_Service/message_broker"
	"BTC_Service/services"
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
	controllers.BtcServ = &services.BtcService{RateProvider: &dal.CoindeskRateProvider{}}
	controllers.MsgBroker = message_broker.CreateMessageBroker(os.Getenv("RABBIT_MQ"), os.Getenv("EXCHANGE"))
	defer controllers.MsgBroker.Close()

	router := mux.NewRouter()
	router.HandleFunc("/btcRate", controllers.Rate).Methods("GET")
	http.Handle("/", router)

	port := os.Getenv("PORT")

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		fmt.Print(err)
	}

	_, _ = fmt.Fprintln(os.Stdout, "Running on port: port")
}
