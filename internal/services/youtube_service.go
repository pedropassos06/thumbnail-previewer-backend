package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/pedropassos06/thumbnail-previewer-backend/internal/models"
	"github.com/pedropassos06/thumbnail-previewer-backend/internal/responses"
)

func FetchChannel(handle string) (*models.Channel, error) {
	apiKey, err := getAPIKey()
	if err != nil {
		return nil, err
	}

	url := buildURL(apiKey, handle)
	channel, err := callYoutubeAPI(url)
	if err != nil {
		return nil, err
	}

	if len(channel.Items) == 0 {
		return nil, fmt.Errorf("channel not found")
	}

	return models.NewChannel(
		channel.Items[0].Snippet.Title,
		channel.Items[0].Snippet.Thumbnails.Default.URL,
	), nil
}

func getAPIKey() (string, error) {
	apiKey := os.Getenv("YOUTUBE_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("YOUTUBE_API_KEY is not set")
	}

	return apiKey, nil
}

func buildURL(apiKey, handle string) string {
	return fmt.Sprintf("https://www.googleapis.com/youtube/v3/channels?part=snippet&fields=items/snippet/title,items/snippet/thumbnails/default&key=%s&forHandle=%s", apiKey, handle)
}

func callYoutubeAPI(url string) (*responses.YoutubeResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch channel: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("youtube api returned status code %d", resp.StatusCode)
	}

	var result responses.YoutubeResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode youtube response: %w", err)
	}

	return &result, nil
}
