package main

import (
	"Api_Gateway/authentication"
	"Api_Gateway/middleware"
	"Api_Gateway/utils"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/urfave/negroni"
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

	authService := authentication.AuthService{}
	userProxy := &httputil.ReverseProxy{Director: userDirector}
	router.HandleFunc("/user/login", UserHandlerFunc(userProxy, authService)).Methods("POST")
	router.HandleFunc("/user/create", UserHandlerFunc(userProxy, authService)).Methods("POST")

	http.Handle("/", router)

	router.Use(middleware.JwtMiddleware)

	port := os.Getenv("PORT")

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		fmt.Print(err)
	}

	_, _ = fmt.Fprintln(os.Stdout, "Running on port: port")
}

func UserHandlerFunc(proxy *httputil.ReverseProxy, authServ authentication.AuthService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		lrw := negroni.NewResponseWriter(w)

		_, _ = w.Write([]byte(""))
		proxy.ServeHTTP(lrw, r)
		statusCode := lrw.Status()

		if statusCode == 200 {
			token, _ := authServ.GenerateToken()
			utils.Respond(lrw, map[string]interface{}{"securityToken": token})
		}
	}
}
