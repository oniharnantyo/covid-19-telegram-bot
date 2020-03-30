package main

import (
	config "covid19-bot/config"
	"covid19-bot/handlers"
	"covid19-bot/mathdroid"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
)

func main() {
	config, err := config.Init(".config.toml")
	if err != nil {
		log.Fatal(err)
	}

	mathdroid.Init(config.MathdroidURL)

	webhook := &tb.Webhook{
		Listen: ":" + config.Port,
		Endpoint: &tb.WebhookEndpoint{
			PublicURL: config.PublicUrl,
		},
	}

	pref := tb.Settings{
		Token:  config.Token,
		Poller: webhook,
	}

	b, err := tb.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	handlers.Init(b, config.LocationiqToken)

	b.Start()
}
