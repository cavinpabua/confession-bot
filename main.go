package main

import (
	bot "confession_bot/bot"
	"flag"
)

var (
	BotToken = flag.String("token", "", "Bot access token")
)

func main() {
	flag.Parse()

	bot.BotToken = *BotToken
	bot.Run()
}
