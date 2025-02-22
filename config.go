package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

type AthocsTgbotConfig struct {
	Token   string
	GroupID telego.ChatID
}

var Config = GetConfig()

func GetConfig() AthocsTgbotConfig {
	id, err := strconv.ParseInt(os.Getenv("ATHOCS_TGBOT_GROUPID"), 10, 64)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return AthocsTgbotConfig{
		Token:   os.Getenv("ATHOCS_TGBOT_TOKEN"),
		GroupID: tu.ID(id),
	}
}
