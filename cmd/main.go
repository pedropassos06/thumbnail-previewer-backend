package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/golang/groupcache/lru"
	"github.com/pedropassos06/thumbnail-previewer-backend/internal/config"
	"github.com/pedropassos06/thumbnail-previewer-backend/internal/router"
)

func init() {
	config.LoadEnv()

}

func main() {
	fmt.Println("Starting server...")

	// Initialize LRU cache
	cache := lru.New(1000)

	// Initialize router with cache
	r := chi.NewRouter()
	router.InitRoutes(r, cache)

	// Get port from env
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		log.Fatalf("PORT not set in .env file")
	}

	log.Printf("Starting server on port %s", port)
	http.ListenAndServe(":"+port, r)
}
