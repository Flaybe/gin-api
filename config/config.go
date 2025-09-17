package config

import "time"

// Some should variables should be in a .env but for clarity are placed here
const (
	ServerPort     = ":8080"
	// Simulated AI response latency
	AILatency      = 1 * time.Second
	// timeout for /ask requests
	DefaultTimeout = 5 * time.Second
	// App environment: "dev" or "prod"
	Env = "prod"
	// Allowed origins (for CORS)
	AllowedOrigins = "http://localhost:3000"
)
