package servers

import (
	"github.com/julienschmidt/httprouter"
	"log/slog"
	"net/http"
)

type HttpRouter struct {
}

func NewHttpRouter() *HttpRouter {
	return &HttpRouter{}
}

func (h *HttpRouter) Init() {
	slog.Info("Initializing httprouter")
	router := httprouter.New()

	router.GET("/health", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("{\"status\": \"OK\"}"))
	})

	router.GET("/test/ok", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("{\"test\": \"OK\"}"))
	})

	router.GET("/test/ko", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("{\"test\": \"KO\"}"))
	})

	if err := http.ListenAndServe(":8088", router); err != nil {
		slog.Error("Erreur lors du lancement du serveur httprouter: %v", err)
	}
}