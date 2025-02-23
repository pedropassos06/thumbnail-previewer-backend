package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pedropassos06/thumbnail-previewer-backend/internal/handlers"
	"github.com/pedropassos06/thumbnail-previewer-backend/internal/responses"
)

func router(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch req.Path {
	case "/ping":
		return handlers.PingHandler(ctx, req)
	case "/channel":
		return handlers.GetChannelHandler(ctx, req)
	default:
		return responses.SendError(http.StatusNotFound, "Not Found")
	}
}

func main() {
	lambda.Start(router)
}
