package servers

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type Gin struct {
}

func NewGin() *Gin {
	return &Gin{}
}

func (g *Gin) Init() {
	slog.Info("Initializing go gin")
	router := gin.New()
	gin.SetMode(gin.ReleaseMode)

	router.Use(func(context *gin.Context) {
		//do whatever you need here
	})

	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "{\"status\": \"OK\"}")
	})

	testGroup := router.Group("/test")
	testGroup.GET("/ok", func(c *gin.Context) {
		c.String(http.StatusOK, "{\"test\": \"OK\"}")
	})
	testGroup.GET("/ko", func(c *gin.Context) {
		c.String(http.StatusOK, "{\"test\": \"KO\"}")
	})

	if err := router.Run(":8080"); err != nil {
		slog.Error("Erreur lors du lancement du serveur Gin: %v", err)
	}
}
