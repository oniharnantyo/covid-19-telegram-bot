package handlers

import (
	"context"
	"covid19-bot/mathdroid"
	"fmt"
	"github.com/codingsince1985/geo-golang/locationiq"
	"gopkg.in/tucnak/telebot.v2"
	"log"
	"time"
)

var (
	START_RESOURCES = `
Covid-19 Updates #ItsCoronaTime
				
Hi %s, This bot provide latest information about Covid-19 Updates

Worldwide summary : 
😷 Confirmed	: %d
😃 Recovered	: %d
😇 Death			: %d

👉 /latest  Get latest updates
👉 /death  Get latest deaths summary

Data source : [covid-19 API](https://github.com/mathdroid/covid-19-api)
Created By : [Oni Harnantyo](https://www.linkedin.com/in/oniharnantyo/)

Keep fight againts Covid-19 and Stay Safe!!!`

	LOCATION_RESOURCE = `
%s summary : 
😷 Confirmed	: %d
😃 Recovered	: %d
😇 Death			: %d

Last Update : %s
`
)

func Init(bot *telebot.Bot, locationiqToken string) *telebot.Bot {
	ctx := context.Background()

	summary, err := mathdroid.GetSummary(ctx)
	if err != nil {
		log.Println(err)
	}

	bot.Handle("/start", func(m *telebot.Message) {
		bot.Send(m.Sender, fmt.Sprintf(START_RESOURCES, m.Sender.FirstName, summary.Confirmed.Value, summary.Recovered.Value, summary.Deaths.Value), &telebot.SendOptions{
			ParseMode: telebot.ModeMarkdown,
		})
	})

	bot.Handle("/latest", func(m *telebot.Message) {
		ctx := context.Background()

		confirmeds, err := mathdroid.GetConfirmed(ctx)
		if err != nil {
			log.Println(err)
		}

		confirmeds = confirmeds[0:20]

		bodyTables := `
| Country | 😷 Confirmed | 😃 Recovered | 😇 Death |
`
		for _, confirmed := range confirmeds {
			fmt.Println("row", confirmed)
			bodyTables = fmt.Sprintf(`
%s
| %s | %d | %d | %d |
		`, bodyTables, confirmed.CountryRegion, confirmed.Confirmed, confirmed.Recovered, confirmed.Deaths)
		}

		bot.Send(m.Sender, fmt.Sprintf(`
%s
`, bodyTables), &telebot.SendOptions{
			ParseMode: telebot.ModeMarkdown,
		})
	})

	bot.Handle("/around-me", func(m *telebot.Message) {
		bot.Send(m.Sender, fmt.Sprintf(`
Please, send your location to get information.
`))
	})

	bot.Handle(telebot.OnLocation, func(q *telebot.Message) {
		location, err := locationiq.Geocoder(locationiqToken, 18).ReverseGeocode(float64(q.Location.Lat), float64(q.Location.Lng))
		if err != nil {
			fmt.Println(err)
		}

		if location == nil {
			bot.Send(q.Sender, `Sorry, your location not found`,
				&telebot.SendOptions{ParseMode: telebot.ModeMarkdown})
		} else {

			summary, err := mathdroid.GetSummaryPerCountry(ctx, location.CountryCode)
			if err != nil {
				log.Println(err)
			}

			bot.Send(q.Sender, fmt.Sprintf(LOCATION_RESOURCE, location.Country, summary.Confirmed.Value, summary.Recovered.Value, summary.Deaths.Value, summary.LastUpdate.Format(time.RFC850)),
				&telebot.SendOptions{ParseMode: telebot.ModeMarkdown})
		}
	})

	return bot
}
