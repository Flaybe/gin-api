package api

import (
	"time"

	"gin-api/config"
	"gin-api/internal/api/transport"
	"gin-api/internal/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	app := gin.Default()

	// CORS by environment
	if config.Env == "prod" {
		app.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:3000"},
			AllowMethods:     []string{"GET", "POST", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))
	} else {
		app.Use(cors.Default())
	}

	// Wire services and handlers
	ai := services.NewAIService(config.AILatency)
	handler  := transport.NewHandler(ai)

	// Routes
	app.GET("/health", handler.Health)
	app.POST("/ask", handler.Ask)

	return app
}
