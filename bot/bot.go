package bot

import (
	"os"

	"github.com/nicklaw5/helix"
)

// Bot provides different bot clients
type Bot struct {
	twitchClientID string
	twitchClientSecret string
	twichRefreshToken string
	HelixClient *helix.Client;
}

// NewBot creates a new bot instance
func  NewBot() (*Bot, error) {
	var err error
	bot := new(Bot)
	bot.twitchClientID = os.Getenv("TWITCH_CLIENT_ID")
	bot.twitchClientSecret = os.Getenv("TWITCH_CLIENT_SECRET")
	bot.twichRefreshToken = os.Getenv("TWITCH_REFRESH_TOKEN")
	bot.HelixClient, err = helix.NewClient(&helix.Options{
		ClientID: bot.twitchClientID,
		ClientSecret: bot.twitchClientSecret,
	}); if err != nil {
		return nil, err
	}

	return bot, nil
}