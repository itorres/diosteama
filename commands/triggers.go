package commands

import (
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// Triggers checks if some text triggers a response
func Triggers(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	var msg tgbotapi.MessageConfig
	var reply string

	if strings.Contains(strings.ToLower(update.Message.Text), "almeida") {
		reply = "¡¡CARAPOLLA!!"
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		bot.Send(msg)
		return
	}

	if strings.Contains(strings.ToLower(update.Message.Text), "ayudita") {
		cmdW00g(update, bot, nil)
		return
	}

	if strings.Contains(strings.ToLower(update.Message.Text), "carme") {
		reply = "PUTAAAAAAAAAA"
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		bot.Send(msg)
		return
	}

	if strings.Contains(strings.ToLower(update.Message.Text), "gamba") {
		reply = "MARIPURI!"
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		bot.Send(msg)
		return
	}
}
