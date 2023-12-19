package main

import (
	"fmt"
	"github.com/labstack/echo"
	"log/slog"
	"net/http"
)

func main() {
	slog.Info("Initializing router")
	//      10s test    port            -c 50                   -c 200                  -c 1000
	go localEcho() //8082       557608 | 10,69MB/s      509854 |  9,77MB/s      463519 | 8,89MB/s
	var forever chan struct{}
	slog.Info("Infinite loop")
	<-forever
}
func localEcho() {
	slog.Info("Initializing echo")
	e := echo.New()

	e.Use(func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		//do whatever you need here
		return handlerFunc
	})
	//QueryParam along with PathVariable Get API
	e.GET("/health*", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("{\"status\": \"OK\"}"))
	})
	testGroup := e.Group("test")
	testGroup.GET("/ok", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("{\"test\": \"OK\"}"))
	})
	testGroup.GET("/ko", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("{\"test\": \"KO\"}"))
	})
	e.Start("localhost:8082")
}
