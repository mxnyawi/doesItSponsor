package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/juju/ratelimit"
	"github.com/mxnyawi/doesItSponsor/internal/db"
	"github.com/mxnyawi/doesItSponsor/internal/handler"
	"github.com/rs/cors"
)

func StartServer() {
	fmt.Println("Connecting to database")
	// Initialize database connection
	database, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.Close()

	fmt.Println("Connected to database")
	// Initialize handler with database dependency
	handler := &handler.Handler{
		DB: database,
	}

	router := mux.NewRouter()
	// Create a rate limiter allowing 100 requests per second with burst limit of 10
	limiter := ratelimit.NewBucketWithRate(100, 100)

	// Define routes with rate limiting middleware
	SetRoutes(router, limiter, handler)

	// Enable CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	})
	fullHandler := c.Handler(router)

	fmt.Println("Starting server")
	// Start the server
	log.Fatal(http.ListenAndServe(":8080", fullHandler))
}

// limitMiddleware applies rate limiting and JWT authentication to the HTTP handler
func limitMiddleware(limiter *ratelimit.Bucket, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Limit requests based on client IP address
		if limiter.TakeAvailable(1) <= 0 {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	}
}
