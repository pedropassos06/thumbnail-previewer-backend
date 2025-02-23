package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/pedropassos06/thumbnail-previewer-backend/internal/models"
	"github.com/pedropassos06/thumbnail-previewer-backend/internal/responses"
)

type YoutubeAPIConsumer interface {
	FetchChannel(handle string) (*models.Channel, error)
}

func FetchChannel(handle string) (*models.Channel, error) {
	apiKey := os.Getenv("YOUTUBE_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("YOUTUBE_API_KEY is not set")
	}

	// Makes a single request to fetch both the channel name and profile picture,
	// reducing unnecessary API calls and improving performance.
	url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/channels?part=snippet&fields=items/snippet/title,items/snippet/thumbnails/default&key=%s&forHandle=%s", apiKey, handle)

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

	if len(result.Items) == 0 {
		return nil, fmt.Errorf("channel not found")
	}

	channel := models.NewChannel(
		result.Items[0].Snippet.Title,
		result.Items[0].Snippet.Thumbnails.Default.URL,
	)

	return channel, nil
}
