package handlers

import (
	"net/http"

	"github.com/golang/groupcache/lru"
	"github.com/pedropassos06/thumbnail-previewer-backend/internal/responses"
	"github.com/pedropassos06/thumbnail-previewer-backend/internal/services"
)

func GetChannelHandler(cache *lru.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		channelHandle := r.URL.Query().Get("handle")
		if channelHandle == "" {
			responses.SendError(w, http.StatusBadRequest, "channel handle is required")
			return
		}

		channel, err := services.FetchChannel(channelHandle, cache)
		if err != nil {
			responses.SendError(w, http.StatusInternalServerError, err.Error())
			return
		}

		responses.SendSuccess(w, channel)
	}
}
