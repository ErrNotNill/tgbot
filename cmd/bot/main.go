package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"tgbot/pkg/telegram"
	"tgbot/server"
)

const Token = "1953280162:AAEWa1cqSHlpHCTi-GEu7CjGyrIt2k7jAAs"
const redirectUrl = "localhost:8000"

func main() {
	go func() {
		server.RunServer()
	}()

	bot, err := tgbotapi.NewBotAPI(Token)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	newbot := telegram.NewBot(bot)
	if err := newbot.StartBot(); err != nil {
		log.Fatal(err)
	}
}

//socket := new(websocket.Conn)
//	scon := new(websocket.Dialer)
//
//	scon.Dial("")
//	fmt.Println(saddr)
