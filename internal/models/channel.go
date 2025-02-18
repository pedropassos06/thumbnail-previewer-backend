package models

type Channel struct {
	ChannelName          string `json:"channel_name"`
	ChannelProfilePicURL string `json:"channel_profile_pic"`
}

func NewChannel(channelName, ChannelProfilePicURL string) *Channel {
	return &Channel{
		ChannelName:          channelName,
		ChannelProfilePicURL: ChannelProfilePicURL,
	}
}
