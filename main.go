package main

import (
	"context"
	"gin-api/config"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	log.Printf("Starting server in %s mode on %s\n", config.Env, config.ServerPort)

	// Configure CORS
	if config.Env == "prod" {
		// Stricter config for production
		app.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:3000"}, // frontend url example
			AllowMethods:     []string{"GET", "POST", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))
	} else {
		// Default: allow everything (dev)
		app.Use(cors.Default())
	}

	// /health endpoint
	app.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// /ask endpoint
	app.POST("/ask", func(c *gin.Context) {
		// expected input
		var input struct {
			Question string `json:"question"`
		}

		err := c.ShouldBindJSON(&input)
		// if we couldn't bind the input to our struct (input doesn't contain {'question': 'question text'}
		// or i if the 'question' field is empy we throw a bad request error (400)
		if err != nil || input.Question == "" {
			// http.StatusBadRequest: 400
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing input (question:)"})
			return
		}
		// Create a context with timeout (e.g. 2 seconds)
		ctx, cancel := context.WithTimeout(c.Request.Context(), config.DefaultTimeout)
		defer cancel()

		result := make(chan string, 1)

		// Run answerAI in a goroutine
		go func() {
			result <- answerAI(input.Question)
		}()

		select {
		case <-ctx.Done():
			// Context expired → timeout
			c.JSON(http.StatusGatewayTimeout, gin.H{"error": "AI timed out"})
			return
		case answer := <-result:
			// Got result in time
			c.JSON(http.StatusOK, gin.H{
				"answer": answer,
				"source": "stubbed",
			})
		}
	})

	app.Run(config.ServerPort)
}

func answerAI(question string) string {
	// Simulate time it takes for response from AI
	time.Sleep(config.AILatency)
	return "This is a test answer from the AI for the question: " + question
}
