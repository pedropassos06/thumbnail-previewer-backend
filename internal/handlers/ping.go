package handlers

import (
	"net/http"

	"github.com/pedropassos06/thumbnail-previewer-backend/internal/responses"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	responses.SendSuccess(w, "pong")
}
