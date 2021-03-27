package bot

import (
	"fmt"

	"github.com/Amazeful/AmazefulBot-Twitter/helix"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type Auth struct {
	TwitterConsumerKey    string
	TwitterConsumerSecret string
	TwitterAccessToken    string
	TwitterAccessSecret   string
	TwitchClientID        string
	TwitchClientSecret    string
	TwitchAccessToken     string
	TwitchRefreshToken    string
}

type Bot struct {
	HelixClient   *helix.Client
	TwitterClient *twitter.Client
	Auth          *Auth
}

func CreateBot(auth *Auth) (*Bot, error) {
	config := oauth1.NewConfig(auth.TwitterConsumerKey, auth.TwitterConsumerSecret)
	token := oauth1.NewToken(auth.TwitterAccessToken, auth.TwitterAccessSecret)

	httpClient := config.Client(oauth1.NoContext, token)

	var bot *Bot = new(Bot)
	bot.Auth = auth

	bot.TwitterClient = twitter.NewClient(httpClient)

	var err error
	bot.HelixClient, err = helix.NewClient(&helix.Options{
		ClientID:        auth.TwitchClientID,
		ClientSecret:    auth.TwitchClientSecret,
		UserAccessToken: auth.TwitchAccessToken,
	})

	if err != nil {
		return nil, err
	}

	err = verify(bot)
	if err != nil {
		return nil, err
	}

	return bot, nil
}

func verify(bot *Bot) error {
	verify := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	_, _, err := bot.TwitterClient.Accounts.VerifyCredentials(verify)
	if err != nil {
		return err
	}
	userAccessToken := bot.HelixClient.GetUserAccessToken()

	isValid, _, err := bot.HelixClient.ValidateToken(userAccessToken)
	if err != nil {
		return err
	}

	if isValid {
		fmt.Printf("%s access token is valid!\n", userAccessToken)
	}
	return nil
}
