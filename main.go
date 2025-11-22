package main

import (
	"context"
	"fmt"
	"os"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {
	// create bot
	ctx := context.Background()
	bot, err := telego.NewBot(Config.Token, telego.WithDefaultLogger(false, true))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// start loop
	go monitorBatteryLoop(bot, ctx)

	// check for telegram updates
	updates, _ := bot.UpdatesViaLongPolling(ctx, nil)
	bh, _ := th.NewBotHandler(bot, updates)
	defer func() { _ = bh.Stop() }()

	// commands
	{
		// /start
		bh.Handle(func(ctx *th.Context, update telego.Update) error {
			_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
				tu.ID(update.Message.Chat.ID),
				"I'm up and running!\nWill let you know if any of the stations' batteries are empty.",
			))
			return nil
		}, th.CommandEqual("start"))

		// other
		bh.Handle(func(ctx *th.Context, update telego.Update) error {
			_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
				tu.ID(update.Message.Chat.ID),
				"Unknown command, use /start",
			))
			return nil
		}, th.AnyCommand())

	}

	_ = bh.Start()
}
