package router

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/pedropassos06/thumbnail-previewer-backend/internal/handlers"
	"github.com/pedropassos06/thumbnail-previewer-backend/internal/middleware"
)

func InitRoutes(r *chi.Mux) {
	r.Use(middleware.CorsMiddleware)
	r.Get("/channel/ping", handlers.PingHandler)
	r.Get("/channel", handlers.GetChannelHandler)
}

func StartServer() {
	// init new router and routes
	r := chi.NewRouter()
	InitRoutes(r)

	// get port from env
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		log.Fatalf("PORT not set in .env file")
	}

	log.Printf("Starting server on port %s", port)
	// start server
	http.ListenAndServe(":"+port, r)
}
