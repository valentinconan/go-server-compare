package servers

import (
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
)

type Chi struct {
}

func NewChi() *Chi {
	return &Chi{}
}

func (g *Chi) Init() {
	slog.Info("Initializing chi")
	chichi := chi.NewRouter()

	chichi.Use(func(handler http.Handler) http.Handler {
		//do whatever you need here
		return handler
	})

	chichi.Get("/health", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("{\"status\": \"OK\"}"))
	})

	chichi.Route("/test", func(r chi.Router) {
		r.Get("/ok", func(writer http.ResponseWriter, request *http.Request) {
			writer.Write([]byte("{\"test\": \"OK\"}"))
		})

		r.Get("/ko", func(writer http.ResponseWriter, request *http.Request) {
			writer.Write([]byte("{\"test\": \"KO\"}"))
		})
	})

	http.ListenAndServe(":8087", chichi)
}
