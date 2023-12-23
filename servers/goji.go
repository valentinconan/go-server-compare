package servers

import (
	"fmt"
	goji "goji.io"
	"goji.io/pat"
	"log/slog"
	"net/http"
)

type Goji struct {
}

func NewGoji() *Goji {
	return &Goji{}
}

func (g *Goji) Init() {
	slog.Info("Initializing goji")
	mux := goji.NewMux()

	mux.Use(func(handler http.Handler) http.Handler {
		//do whatever you need here
		return handler
	})

	mux.HandleFunc(pat.Get("/health"), func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "{\"status\": \"OK\"}")
	})

	testGroup := goji.SubMux()
	mux.Handle(pat.New("/test/*"), testGroup)
	testGroup.HandleFunc(pat.Get("/ok"), func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "{\"test\": \"OK\"}")
	})
	testGroup.HandleFunc(pat.Get("/ko"), func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "{\"test\": \"KO\"}")
	})

	http.ListenAndServe(":8084", mux)
}
