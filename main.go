package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi/v2"
	"github.com/go-martini/martini"
	"github.com/gofiber/fiber/v2"
	gorilla "github.com/gorilla/mux"
	"github.com/labstack/echo"
	"goji.io"
	"goji.io/pat"
	"log/slog"
	"net/http"
)

func main() {
	slog.Info("Initializing router")
	//      10s test    port            -c 50                   -c 200                  -c 1000
	go localGoGin()   //8080       305409 |  5,86MB/s      301142 |  5,76MB/s      305546 | 5,83MB/s
	go localHttpMux() //8081       539316 | 10,34MB/s      584001 | 11,19MB/s      470207 | 9,00MB/s
	go localEcho()    //8082       557608 | 10,69MB/s      509854 |  9,77MB/s      463519 | 8,89MB/s
	go localMartini() //8083       251361 |  4,82MB/s      286005 |  5,48MB/s      271453 | 5,20MB/s
	go localGoji()    //8084       552464 | 10,59MB/s      584054 | 11,19MB/s      450340 | 8,62MB/s
	go localGorilla() //8085       554316 | 10,63MB/s      572607 | 10,98MB/s      446838 | 8,56MB/s
	go localFiber()   //8086       608215 | 11,66MB/s      555507 | 10,65MB/s      498507 | 9,55MB/s
	go localChi()     //8087       457159 |  8,76MB/s      439088 |  8,41MB/s      373024 | 7,14MB/s
	var forever chan struct{}
	slog.Info("Infinite loop")
	<-forever
}

func localFiber() {
	fib := fiber.New()
	fib.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.SendString("{\"status\": \"OK\"}")
	})
	fib.Listen(":8086")
}
func localChi() {
	chichi := chi.NewRouter()
	chichi.Get("/health", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("{\"status\": \"OK\"}"))
	})
	http.ListenAndServe(":8087", chichi)
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

func localGorilla() {
	slog.Info("Initializing gorilla")
	router := gorilla.NewRouter()
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "{\"status\": \"OK\"}")
	})

	http.ListenAndServe("localhost:8085", router)
}
func localHttpMux() {
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
