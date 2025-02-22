package handler

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

var (
	cache *lru.Cache
	r     *chi.Mux
)

func init() {
	config.LoadEnv()

	// Initialize LRU cache
	cache = lru.New(1000)

	// Initialize router with cache
	r = chi.NewRouter()
	router.InitRoutes(r, cache)
}

func main() {
	fmt.Println("Starting server...")

	// Get port from env
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)
	http.ListenAndServe(":"+port, r)
}

// Handler is the entry point for Vercel requests
func Handler(w http.ResponseWriter, req *http.Request) {
	r.ServeHTTP(w, req)
}
