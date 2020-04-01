# Covid 19 Telegram Bot (It's Corona Time)
Covid 19 Telegram Bot is a simple project to get latest updates about Covid 19.

This project inspired by [COVID-19-Telegram-bot](https://github.com/kasramp/COVID-19-Telegram-bot) made by [Kasra Madadipouya](https://github.com/kasramp/) and uses data from [COVID-19 API](https://github.com/mathdroid/covid-19-api) by [Odi](https://github.com/mathdroid).

## Prerequisites
* Golang 1.12.17 or newer with Go Module Support

## Setup The Project

### Create bot
Create new bot on telegram using [@BotFather](https://telegram.me/BotFather) and copy your token (reference : [https://medium.com/shibinco/create-a-telegram-bot-using-botfather-and-get-the-api-token-900ba00e0f39](https://medium.com/shibinco/create-a-telegram-bot-using-botfather-and-get-the-api-token-900ba00e0f39)).

### Setup Ngrok
Setup ngrok from [https://ngrok.com/](https://ngrok.com/)

```./ngrok http 3000```

### Create Location IQ account and get Token
[https://locationiq.com/](https://locationiq.com/)

### Clone The Project
`git clone https://github.com/oniharnantyo/covid-19-telegram-bot.git`

### Edit Config file
Copy ```example.config.toml``` to ```.config.toml```

`cp example.config.toml .config.toml`

Fill in your Telegram bot token and Location IQ token
```
token = "YOUR_TELEGRAM_BOT_TOKEN"
port = "3000"
mathdroid_url = "https://covid19.mathdro.id/api"
locationiq_token = "YOUR_LOCATIONIQ_TOKEN"
```

### Run The Project
```go run main.go```

### Set url to Telegram Bot
```curl -F "url=https://YOUR_URL"  https://api.telegram.org/bot<YOUR_TELEGRAM_BOT>/setWebhook```

note : When you restart the program, you must hit this url again

## Contributing
Let's make pull request and you will help people around the world.

## Contact
* oni.harnantyo97@gmail.com
