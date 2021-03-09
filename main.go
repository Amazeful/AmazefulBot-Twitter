package main

import (
	"fmt"
	"log"

	"amazefulbot.com/Amazeful/Amazefulbot-Twitter/bot"
)

func main()  {
	fmt.Printf("Hello?")

	bot, err := bot.NewBot(); if err != nil {
		log.Fatal("failed to create a bot instance")
	}
}