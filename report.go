package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func getReportMessage() string {
	records, err := getLiveData()
	if err != nil {
		return fmt.Sprintf("Failed to get live data for report: %v", err)
	}
	list := ""
	for _, station := range records[1:] {
		list += fmt.Sprintf("\n%s: %s°C, %s%% (%s, %s%%)", station[1], station[2], station[3], station[0], station[5])
	}
	return "🫡 Report 🫡\n" + list

}

// main loop
func reportLoop(bot *telego.Bot, ctx context.Context) {
	for {
		bot.SendMessage(ctx, tu.Message(Config.GroupID, getReportMessage()))

		time.Sleep(Config.ReportInterval)
	}
}
