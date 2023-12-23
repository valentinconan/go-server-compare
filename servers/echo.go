package servers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
)

type Echo struct {
}

func NewEcho() *Echo {
	return &Echo{}
}

func (e *Echo) Init() {
	slog.Info("Initializing echo")
	server := echo.New()

	server.Use(func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		//do whatever you need here
		return handlerFunc
	})
	//QueryParam along with PathVariable Get API
	server.GET("/health*", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("{\"status\": \"OK\"}"))
	})
	testGroup := server.Group("test")
	testGroup.GET("/ok", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("{\"test\": \"OK\"}"))
	})
	testGroup.GET("/ko", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("{\"test\": \"KO\"}"))
	})
	server.Start(":8082")
}
