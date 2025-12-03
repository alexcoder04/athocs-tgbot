package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

// get stations data
func checkLowBatteries() ([]string, error) {
	records, err := getLiveData()
	if err != nil {
		return nil, err
	}

	var low []string

	// skip header
	for i := 1; i < len(records); i++ {
		row := records[i]
		if len(row) < 6 {
			continue
		}

		station := row[1]
		batteryStr := row[5]

		batt, err := strconv.Atoi(batteryStr)
		if err != nil {
			continue
		}

		if batt < Config.BatteryAlarmThresh {
			low = append(low, fmt.Sprintf("%s, %d%%", station, batt))
		}
	}

	return low, nil
}

// main loop
func monitorBatteryLoop(bot *telego.Bot, ctx context.Context) {
	for {
		lowStations, err := checkLowBatteries()
		if err != nil {
			bot.SendMessage(ctx, tu.Message(Config.GroupID, fmt.Sprintf("Failed to check battery status: %v", err)))
		} else if len(lowStations) > 0 {
			msg := "⚠ WARNING! ⚠\nFollowing stations have low battery:\n" + strings.Join(lowStations, "\n")
			bot.SendMessage(ctx, tu.Message(Config.GroupID, msg))
		}

		time.Sleep(Config.RefreshInterval)
	}
}
