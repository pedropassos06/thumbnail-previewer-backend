package handlers

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/pedropassos06/thumbnail-previewer-backend/internal/responses"
	"github.com/pedropassos06/thumbnail-previewer-backend/internal/services"
)

func GetChannelHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	channelHandle := req.QueryStringParameters["handle"]
	if channelHandle == "" {
		return responses.SendError(http.StatusBadRequest, "channel handle is required")
	}

	channel, err := services.FetchChannel(channelHandle)
	if err != nil {
		return responses.SendError(http.StatusInternalServerError, err.Error())
	}

	return responses.SendSuccess(channel)
}
