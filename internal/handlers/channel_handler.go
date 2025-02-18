package handlers

import (
	"net/http"

	"github.com/pedropassos06/thumbnail-previewer-backend/internal/responses"
	"github.com/pedropassos06/thumbnail-previewer-backend/internal/services"
)

func GetChannelHandler(w http.ResponseWriter, r *http.Request) {
	channelHandle := r.URL.Query().Get("handle")
	if channelHandle == "" {
		responses.SendError(w, http.StatusBadRequest, "channel handle is required")
		return
	}

	channel, err := services.FetchChannel(channelHandle)
	if err != nil {
		responses.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responses.SendSuccess(w, channel)
}
