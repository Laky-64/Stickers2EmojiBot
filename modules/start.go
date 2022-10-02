package modules

import (
	"Stickers2Emoji/utils"
	"fmt"
	"github.com/Squirrel-Network/gobotapi"
	"github.com/Squirrel-Network/gobotapi/methods"
	"github.com/Squirrel-Network/gobotapi/types"
)

func Start(client *gobotapi.Client, update types.Message) {
	startMessage := fmt.Sprintf("Hello %s, I'm Stickers2Emoji Bot! I can convert your sticker pack to emoji pack or viceversa.", utils.Mention(*update.From))
	startMessage += "\n\n"
	startMessage += "To convert your sticker pack to emoji pack, send me the sticker you want to convert the entire sticker pack. (<a href='https://core.telegram.org/stickers#static-stickers'>works only with static sticker</a>)"
	_, _ = client.Invoke(&methods.SendMessage{
		ChatID:                update.Chat.ID,
		Text:                  startMessage,
		ParseMode:             "HTML",
		DisableWebPagePreview: true,
	})
}
