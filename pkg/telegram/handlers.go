package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil { // If we got a message
			continue
		}
		if update.Message.IsCommand() {
			err := b.handleCommand(update.Message)
			if err != nil {
				log.Fatal(err.Error())
			}
			continue
		}

		err := b.handleMessage(update.Message)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

}

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case "start":
		msg := tgbotapi.NewMessage(message.Chat.ID, "You used command start.")
		msg.ReplyMarkup = b.KeyBoard9Buttons()
		_, err := b.bot.Send(msg)
		if err != nil {
			log.Printf(err.Error())
		}
		return err
	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, "I dont know this command")
		_, err := b.bot.Send(msg)
		return err
	}
	// And even more text...
}

func (b *Bot) handleVideo(video *tgbotapi.Video) {
}
func (b *Bot) initInlines(tg tgbotapi.InlineQuery) {

}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	//msg.Text = "Hello"
	//msg.ReplyToMessageID = message.MessageID

	_, err := b.bot.Send(msg)
	if err != nil {
		log.Fatal("The message was not sent", err.Error())
	}
	return err
}

var Keyboard []*tgbotapi.InlineKeyboardMarkup

type DeleteMessageConfig struct {
	ChannelUsername string
	ChatID          int64
	MessageID       int
}

func (config DeleteMessageConfig) Parameters() (tgbotapi.Params, error) {
	params := make(tgbotapi.Params)
	params.AddFirstValid("chat_id", config.ChatID, config.ChannelUsername)
	params.AddNonZero("message_id", config.MessageID)
	return params, nil
}

func (b *Bot) KeyBoard9Buttons() tgbotapi.InlineKeyboardMarkup {
	//firstkeyboard :=
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("Localhost", "http://127.0.0.1:8000"),
			tgbotapi.NewInlineKeyboardButtonData("2", "2"),
			tgbotapi.NewInlineKeyboardButtonData("3", "3"),
		))
	//secondkeyboard :=
	tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("4", "4"),
			tgbotapi.NewInlineKeyboardButtonData("5", "5"),
			tgbotapi.NewInlineKeyboardButtonData("6", "6"),
		))
	//thirdkeyboard :=
	tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("4", "4"),
			tgbotapi.NewInlineKeyboardButtonData("5", "5"),
			tgbotapi.NewInlineKeyboardButtonData("6", "6"),
		))
	return keyboard
}
