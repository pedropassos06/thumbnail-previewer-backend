package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/golang/groupcache/lru"
	"github.com/pedropassos06/thumbnail-previewer-backend/internal/handlers"
	"github.com/pedropassos06/thumbnail-previewer-backend/internal/middleware"
)

func InitRoutes(r *chi.Mux, cache *lru.Cache) {
	r.Use(middleware.CorsMiddleware)
	r.Get("/channel/ping", handlers.PingHandler)
	r.Get("/channel", handlers.GetChannelHandler(cache))
}
