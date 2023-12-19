package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

func main() {
	slog.Info("Initializing router")
	//      10s test    port            -c 50                   -c 200                  -c 1000
	go localHttpMux() //8081       539316 | 10,34MB/s      584001 | 11,19MB/s      470207 | 9,00MB/s
	var forever chan struct{}
	slog.Info("Infinite loop")
	<-forever
}

func localHttpMux() {
	slog.Info("Initializing mux")
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "{\"status\": \"OK\"}")
	})

	testGroup := http.NewServeMux()
	testGroup.HandleFunc("/ok", func(w http.ResponseWriter, request *http.Request) {
		fmt.Fprint(w, "{\"test\": \"OK\"}")
	})
	testGroup.HandleFunc("/ko", func(w http.ResponseWriter, request *http.Request) {
		fmt.Fprint(w, "{\"test\": \"KO\"}")
	})

	mux.Handle("/", http.StripPrefix("/test", testGroup))

	http.ListenAndServe("localhost:8081", mux)
}
