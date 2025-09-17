package services

import (
	"context"
	"fmt"
	"time"
)

// AIService defines the interface for the AI service
type AIService interface {
	Answer(ctx context.Context, question string) (string, error)
}

// aiService is a stub implementation of AIService
type aiService struct {
	latency time.Duration
}

// NewAIService creates a new AI service with the given simulated latency
func NewAIService(latency time.Duration) AIService {
	return &aiService{latency: latency}
}

func (s *aiService) Answer(ctx context.Context, question string) (string, error) {
	// Respect caller context (timeouts/cancellations)
	select {
	case <-time.After(s.latency):
		return fmt.Sprintf("This is a test answer from the AI for the question: %s", question), nil
	// exit early if caller's context is done
	case <-ctx.Done():
		return "", ctx.Err()
	}
}
