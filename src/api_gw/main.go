package main

import (
	"Api_Gateway/message_receiver"
	"Api_Gateway/middleware"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	btcServiceOrigin, _ := url.Parse(os.Getenv("BTC_SERVICE_ORIGIN"))
	userServiceOrigin, _ := url.Parse(os.Getenv("USER_SERVICE_ORIGIN"))

	router := mux.NewRouter()

	//Btc Proxy
	btcDirector := func(req *http.Request) {
		req.Header.Add("X-Forwarded-Host", req.Host)
		req.Header.Add("X-Origin-Host", btcServiceOrigin.Host)
		req.URL.Scheme = "http"
		req.URL.Host = btcServiceOrigin.Host
	}
	btcProxy := &httputil.ReverseProxy{Director: btcDirector}
	router.HandleFunc("/btcRate", func(w http.ResponseWriter, r *http.Request) {
		btcProxy.ServeHTTP(w, r)
	}).Methods("GET")

	//User Proxy
	userDirector := func(req *http.Request) {
		req.Header.Add("X-Forwarded-Host", req.Host)
		req.Header.Add("X-Origin-Host", userServiceOrigin.Host)
		req.URL.Scheme = "http"
		req.URL.Host = userServiceOrigin.Host
	}

	userProxy := &httputil.ReverseProxy{Director: userDirector}
	router.HandleFunc("/user/login", UserHandlerFunc(userProxy)).Methods("POST")
	router.HandleFunc("/user/create", UserHandlerFunc(userProxy)).Methods("POST")

	http.Handle("/", router)

	router.Use(middleware.JwtMiddleware)

	port := os.Getenv("PORT")

	receiver := message_receiver.CreateMessageReceiver(os.Getenv("RABBIT_MQ"), os.Getenv("EXCHANGE"))
	defer receiver.Close()

	go LogErrorMessages(receiver)

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		fmt.Print(err)
	}

	_, _ = fmt.Fprintln(os.Stdout, "Running on port: port")
}

func UserHandlerFunc(proxy *httputil.ReverseProxy) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	}
}

func LogErrorMessages(receiver *message_receiver.RabbitMQReceiver) {
	for d := range receiver.ErrorChannel {
		log.Printf("%s", d.Body)
	}
}
