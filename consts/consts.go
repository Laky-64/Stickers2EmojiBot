package consts

import (
	"github.com/Squirrel-Network/gobotapi/filters"
	"time"
)

var (
	BotToken            string
	DefaultAntiFlood    = filters.AntiFlood(5, time.Second*5, time.Second*10)
	AliasList           = []string{"/", ";", ".", "!"}
	EmojiSize           = 100
	ConvertingProcesses = make(map[int64]bool)
)
