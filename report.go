package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

// main loop
func reportLoop(bot *telego.Bot, ctx context.Context) {
	for {
		records, err := getLiveData()
		if err != nil {
			bot.SendMessage(ctx, tu.Message(Config.GroupID, fmt.Sprintf("Failed to get live data for report: %v", err)))
		} else {
			list := ""
			for _, station := range records[1:] {
				list += fmt.Sprintf("\n%s: %s°C, %s%% (%s, %s%%)", station[1], station[2], station[3], station[0], station[5])
			}
			msg := "🫡 Report 🫡\n" + list
			bot.SendMessage(ctx, tu.Message(Config.GroupID, msg))
		}

		time.Sleep(Config.ReportInterval)
	}
}
