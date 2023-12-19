package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/labstack/gommon/log"
	"log/slog"
	"net/http"
)

func main() {
	slog.Info("Initializing router")
	//      10s test    port            -c 50                   -c 200                  -c 1000
	go localGorilla() //8085       554316 | 10,63MB/s      572607 | 10,98MB/s      446838 | 8,56MB/s
	var forever chan struct{}
	slog.Info("Infinite loop")
	<-forever
}

func localGorilla() {
	slog.Info("Initializing gorilla")

	router := mux.NewRouter()
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "{\"status\": \"OK\"}")
	})

	router.Use(func(handler http.Handler) http.Handler {
		//do whatever you need here
		return handler
	})

	testRoute := router.PathPrefix("/test").Subrouter()
	testRoute.HandleFunc("/ok", func(w http.ResponseWriter, request *http.Request) {
		fmt.Fprint(w, "{\"test\": \"OK\"}")
	})
	testRoute.HandleFunc("/ko", func(w http.ResponseWriter, request *http.Request) {
		fmt.Fprint(w, "{\"test\": \"KO\"}")
	})

	log.Fatal(http.ListenAndServe("localhost:8081", router))
}
