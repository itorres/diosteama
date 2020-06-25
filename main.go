package main

import (
	"log"
	"os"
	"strings"

	"encoding/json"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/pandacrew-net/diosteama/commands"
	"github.com/pandacrew-net/diosteama/database"
)

// Quote is the full quote
type Quote struct {
	Recnum   int
	Date     string
	Author   string
	Text     string
	Messages []*tgbotapi.Message
	From     tgbotapi.User
}

func main() {
	var err error
	token := os.Getenv("TELEGRAM_BOT_TOKEN")

	database.Init()
	database.Info(0)
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		j, _ := json.Marshal(update)
		log.Printf("%s", j)
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		response(update, bot)
		log.Printf("[%s] %s (%v)", update.Message.From.UserName, update.Message.Text, update.Message.IsCommand())

	}
}

func response(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	var msg tgbotapi.MessageConfig
	if commands.EvalAddquote(update) {
		// This is a forward part of an !addquote and has been processed. Return.
		return
	}
	if len(update.Message.Text) > 0 && (string(update.Message.Text[0]) == "!" || string(update.Message.Text[0]) == "/") {
		commands.Command(update, bot)
	} else if strings.Contains(strings.ToLower(update.Message.Text), "almeida") {
		reply := "¡¡CARAPOLLA!!"
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		bot.Send(msg)
	} else if strings.Contains(strings.ToLower(update.Message.Text), "ayudita") {
		reply := "Capitan castor, ayuditaaaaaaaaaaaaaaaa!!!"
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		bot.Send(msg)
	} else if strings.Contains(strings.ToLower(update.Message.Text), "carme") {
		reply := "PUTAAAAAAAAAA"
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		bot.Send(msg)
	} else if strings.Contains(strings.ToLower(update.Message.Text), "gamba") {
		reply := "MARIPURI!"
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		bot.Send(msg)
	}
}
