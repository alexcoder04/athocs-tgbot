package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

type AthocsTgbotConfig struct {
	ApiUrl             string
	BatteryAlarmThresh int
	GroupID            telego.ChatID
	RefreshInterval    time.Duration
	Token              string
}

var Config = GetConfig()

func GetConfig() AthocsTgbotConfig {
	id, err := strconv.ParseInt(os.Getenv("ATHOCS_TGBOT_GROUPID"), 10, 64)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return AthocsTgbotConfig{
		ApiUrl:             os.Getenv("ATHOCS_TGBOT_APIURL"),
		BatteryAlarmThresh: 40,
		GroupID:            tu.ID(id),
		RefreshInterval:    30 * time.Minute,
		Token:              os.Getenv("ATHOCS_TGBOT_TOKEN"),
	}
}
