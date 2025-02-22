package main

import (
	"fmt"
	"os"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {
	bot, err := telego.NewBot(Config.Token, telego.WithDefaultLogger(false, true))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	bot.SendMessage(tu.Message(Config.GroupID, "Bot powered up"))

	updates, _ := bot.UpdatesViaLongPolling(&telego.GetUpdatesParams{})

	for update := range updates {
		if update.Message != nil {
			fmt.Printf("Got message '%s' from %s in chat '%s'\n", update.Message.Text, update.Message.From.Username, update.Message.Chat.Title)
		}
	}
}
