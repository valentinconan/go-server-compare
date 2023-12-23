package servers

import (
	"fmt"
	"log/slog"
	"net/http"
)

type HttpMux struct {
}

func NewHttpMux() *HttpMux {
	return &HttpMux{}
}

func (h *HttpMux) Init() {
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

	http.ListenAndServe(":8081", Use(mux))
}

func Use(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//do whatever you need here
		next.ServeHTTP(w, r)
	})
}
