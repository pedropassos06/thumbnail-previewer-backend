package services

import (
	"testing"
)

func TestFetchChannel(t *testing.T) {
	// Define expected output
	expectedName := "Pedro Passos"
	expectedProfilePic := "https://yt3.ggpht.com/LfQxT63Q6ndKEzMysi3cemjdoiY8vep3L9b6O5M56ZKJtkyhZRpkFo-0pQggb5c06cdoWclp5A=s88-c-k-c0x00ffffff-no-rj"

	// Call the function
	channel, err := FetchChannel("@pedropassos_")
	if err != nil {
		t.Fatalf("FetchChannel returned an error: %v", err)
	}

	// Check if the response matches expected values
	if channel.ChannelName != expectedName {
		t.Errorf("Expected channel name %q, but got %q", expectedName, channel.ChannelName)
	}

	if channel.ChannelProfilePicURL != expectedProfilePic {
		t.Errorf("Expected profile pic %q, but got %q", expectedProfilePic, channel.ChannelProfilePicURL)
	}
}
