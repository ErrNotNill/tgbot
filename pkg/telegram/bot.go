package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"html/template"
	"log"
)

type Bot struct {
	bot         *tgbotapi.BotAPI
	site        *template.Template
	redirectUrl string
}

func NewBot(bot *tgbotapi.BotAPI, site *template.Template, redirectUrl string) *Bot {
	return &Bot{bot: bot, site: site, redirectUrl: redirectUrl}
}

func (b *Bot) StartBot() error {
	log.Printf("Authorized on account #%s name is #%v inline is:%v", b.bot.Self.UserName, b.bot.Self.FirstName, b.bot.Self.SupportsInlineQueries)
	updates := b.initUpdatesChannel()
	b.handleUpdates(updates)
	return nil
}

func (b *Bot) initUpdatesChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := b.bot.GetUpdatesChan(u)

	return updates
}
