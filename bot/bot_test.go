package bot

import (
	"fmt"
	"os"
	"testing"

	"github.com/Amazeful/AmazefulBot-Twitter/helix"
)

func TestHelix(t *testing.T) {
	auth := &Auth{
		TwitterConsumerKey:    os.Getenv("CONSUMER_KEY"),
		TwitterConsumerSecret: os.Getenv("CONSUMER_SECRET"),
		TwitterAccessToken:    os.Getenv("TWITTER_ACCESS_TOKEN"),
		TwitterAccessSecret:   os.Getenv("TWITTER_ACCESS_SECRET"),
		TwitchClientID:        os.Getenv("TWITCH_CLIENT_ID"),
		TwitchClientSecret:    os.Getenv("TWITCH_CLIENT_SECRET"),
		TwitchAccessToken:     os.Getenv("TWITCH_ACCESS_TOKEN"),
		TwitchRefreshToken:    os.Getenv("TWITCH_REFRESH_TOKEN"),
	}
	bot, err := CreateBot(auth)
	if err != nil {
		t.Fatalf("ERROR %s", err)
	}
	refreshToken := auth.TwitchRefreshToken
	res1, err := bot.HelixClient.RefreshUserAccessToken(refreshToken)
	if err != nil {
		t.Fatalf("ERROR %s", err)
	}
	fmt.Printf("Status code: %+v\n", res1)

	resp, err := bot.HelixClient.GetUsers(&helix.UsersParams{
		IDs:    []string{"26301881", "18074328"},
		Logins: []string{"summit1g", "lirik"},
	})

	if err != nil {
		t.Fatalf("ERROR %s", err)
	}

	fmt.Printf("Rate limit: %d\n", resp.GetRateLimit())
	fmt.Printf("Rate limit remaining: %d\n", resp.GetRateLimitRemaining())
	fmt.Printf("Rate limit reset: %d\n\n", resp.GetRateLimitReset())

	for _, user := range resp.Data.Users {
		fmt.Printf("ID: %s Name: %s\n", user.ID, user.DisplayName)
	}

}
