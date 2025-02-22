package router

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/golang/groupcache/lru"
	"github.com/pedropassos06/thumbnail-previewer-backend/internal/handlers"
	"github.com/pedropassos06/thumbnail-previewer-backend/internal/middleware"
	"golang.org/x/time/rate"
)

func InitRoutes(r *chi.Mux, cache *lru.Cache) {
	r.Use(middleware.CorsMiddleware)
	r.Get("/channel/ping", handlers.PingHandler)

	// Initialize rate limiter
	limiter := rate.NewLimiter(rate.Every(time.Minute/3), 3)

	// Initialize route with rate limiter
	r.With(middleware.RateLimitMiddleware(limiter)).Get("/channel", handlers.GetChannelHandler(cache))

}
