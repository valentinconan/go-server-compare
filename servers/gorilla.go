package servers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/labstack/gommon/log"
	"log/slog"
	"net/http"
)

type Gorilla struct {
}

func NewGorilla() *Gorilla {
	return &Gorilla{}
}

func (g *Gorilla) Init() {
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

	log.Fatal(http.ListenAndServe(":8085", router))
}
