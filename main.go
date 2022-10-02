package main

import (
	"Stickers2Emoji/cfilters"
	"Stickers2Emoji/consts"
	"Stickers2Emoji/modules"
	"github.com/Squirrel-Network/gobotapi"
	"github.com/Squirrel-Network/gobotapi/filters"
)

func main() {
	consts.LoadEnv()
	client := gobotapi.NewClient(consts.BotToken)
	client.SleepThreshold = 60
	client.OnAnyMessageEvent(filters.Filter(modules.Start, filters.And(filters.Command("start", consts.AliasList...), filters.Private(), consts.DefaultAntiFlood)))
	client.OnMessage(filters.Filter(modules.ListenerBeta, filters.And(cfilters.Sticker(), filters.Private(), consts.DefaultAntiFlood)))
	_ = client.Run()
}
