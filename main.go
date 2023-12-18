package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-martini/martini"
	"github.com/labstack/echo"
	"goji.io"
	"goji.io/pat"
	"log/slog"
	"net/http"
)

func main() {
	slog.Info("Initializing router")
	//                     -c 50           -c 200
	go localGoGin()   //50 794.50KB/s       847.27KB/s
	go localEcho()    //2 7.32MB/s          8.22MB/s
	go localMux()     //1 7.05MB/s			8.31MB/s
	go localMartini() //4 1.13MB/s
	go localGoji()    //3 6.48MB/s
	go localHttp()    //  6.22MB/s
	var forever chan struct{}
	slog.Info("Infinite loop")
	<-forever
}

func localHttp() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "{\"status\": \"OK\"}")
	})

	if err := http.ListenAndServe(":8085", nil); err != nil {
		panic(err)
	}
}

func localGoji() {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/health"), func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "{\"status\": \"OK\"}")
	})

	http.ListenAndServe("localhost:8084", mux)
}

func localMartini() {
	m := martini.Classic()
	m.Get("/health", func() string {
		return "{\"status\": \"OK\"}"
	})
	m.RunOnAddr("localhost:8083")
}
func localMux() {
	slog.Info("Initializing mux")
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "{\"status\": \"OK\"}")
	})

	http.ListenAndServe("localhost:8081", mux)
}
func localGoGin() {
	slog.Info("Initializing go gin")
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "{\"status\": \"OK\"}")
	})

	if err := router.Run(":8080"); err != nil {
		slog.Error("Erreur lors du lancement du serveur Gin: %v", err)
	}
}

func localEcho() {
	slog.Info("Initializing echo")
	e := echo.New()
	//QueryParam along with PathVariable Get API
	e.GET("/health*", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("{\"status\": \"OK\"}"))
	})
	e.Start("localhost:8082")
}
