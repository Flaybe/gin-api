package transport

import (
	"context"
	"net/http"

	"gin-api/config"
	"gin-api/internal/models"
	"gin-api/internal/services"

	"github.com/gin-gonic/gin"
)

// Handler holds the AI service and handles HTTP requests
type Handler struct {
	ai services.AIService
}

// NewHandler creates a new handler with the AI service injected
func NewHandler(ai services.AIService) *Handler {
	return &Handler{ai: ai}
}

// Health is a simple endpoint to check if the server is up
func (handler *Handler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// Ask handles the AI question request
func (handler *Handler) Ask(c *gin.Context) {
	var request models.AskRequest
	if err := c.ShouldBindJSON(&request); err != nil || request.Question == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing input (question:)"})
		return
	}

	// Create a context with timeout to avoid hanging requests
	timeoutCtx, cancel := context.WithTimeout(c.Request.Context(), config.DefaultTimeout)
	defer cancel()

	// Call the AI service to get an answer
	answer, err := handler.ai.Answer(timeoutCtx, request.Question)
	if err != nil {
		// Check if it timed out
		if timeoutCtx.Err() == context.DeadlineExceeded {
			c.JSON(http.StatusGatewayTimeout, gin.H{"error": "AI timed out"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "AI failed"})
		return
	}
	// Respond with the AI answer
	c.JSON(http.StatusOK, models.AskResponse{
		Answer: answer,
		Source: "stubbed",
	})
}
