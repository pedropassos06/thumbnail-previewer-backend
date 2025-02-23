package handlers

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/pedropassos06/thumbnail-previewer-backend/internal/responses"
)

func PingHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return responses.SendSuccess(map[string]string{"message": "pong"})
}
