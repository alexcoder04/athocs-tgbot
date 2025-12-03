package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

var Started = time.Now().Format("2006-01-02_15:04:05")

func main() {
	// create bot
	ctx := context.Background()
	bot, err := telego.NewBot(Config.Token, telego.WithDefaultLogger(false, true))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// start loops
	go monitorBatteryLoop(bot, ctx)
	go reportLoop(bot, ctx)

	// check for telegram updates
	updates, _ := bot.UpdatesViaLongPolling(ctx, nil)
	bh, _ := th.NewBotHandler(bot, updates)
	defer func() { _ = bh.Stop() }()

	// commands
	{
		// /start
		bh.Handle(func(ctx *th.Context, update telego.Update) error {
			hostname, err := os.Hostname()
			if err != nil {
				hostname = "unknown"
			}
			_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
				tu.ID(update.Message.Chat.ID),
				fmt.Sprintf("I'm up and running on %s since %s!", hostname, Started),
			))
			return nil
		}, th.CommandEqual("start"))

		// /live
		bh.Handle(func(ctx *th.Context, update telego.Update) error {
			_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
				tu.ID(update.Message.Chat.ID),
				getReportMessage(),
			))
			return nil
		}, th.CommandEqual("report"))

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
