package modules

import (
	"Stickers2Emoji/consts"
	"Stickers2Emoji/utils"
	"archive/zip"
	"bytes"
	"github.com/Squirrel-Network/gobotapi"
	"github.com/Squirrel-Network/gobotapi/methods"
	"github.com/Squirrel-Network/gobotapi/types"
)

func ListenerBeta(client *gobotapi.Client, update types.Message) {
	if !consts.ConvertingProcesses[update.Chat.ID] {
		consts.ConvertingProcesses[update.Chat.ID] = true
		rawMess, _ := client.Invoke(&methods.SendMessage{
			ChatID:    update.Chat.ID,
			Text:      "The <b>emoji pack</b> making is in progress. Please wait until it's finished.",
			ParseMode: "HTML",
		})
		message := rawMess.Result.(types.Message)
		res, _ := client.Invoke(&methods.GetStickerSet{
			Name: update.Sticker.SetName,
		})
		stickerSet := res.Result.(types.StickerSet)
		stickers, err := utils.DownloadStickers(client, stickerSet)
		if err != nil {
			_, _ = client.Invoke(&methods.EditMessageText{
				ChatID:    message.Chat.ID,
				MessageID: message.MessageID,
				Text:      "<b>Error while downloading stickers</b>",
			})
			return
		}
		convertedStickers, err := utils.ConvertStickers(stickers)
		if err != nil {
			_, _ = client.Invoke(&methods.EditMessageText{
				ChatID:    message.Chat.ID,
				MessageID: message.MessageID,
				Text:      "<b>Error while converting stickers</b>",
			})
			return
		}
		var zipFile bytes.Buffer
		zipWriter := zip.NewWriter(&zipFile)
		for _, sticker := range convertedStickers {
			create, _ := zipWriter.Create(sticker.Emoji + ".webp")
			_, _ = create.Write(sticker.Data)
		}
		_ = zipWriter.Close()
		_, _ = client.Invoke(&methods.SendDocument{
			ChatID: update.Chat.ID,
			Document: types.InputBytes{
				Name: "emoji.zip",
				Data: zipFile.Bytes(),
			},
		})
		_, _ = client.Invoke(&methods.DeleteMessage{
			ChatID:    message.Chat.ID,
			MessageID: message.MessageID,
		})
		delete(consts.ConvertingProcesses, update.Chat.ID)
	} else {
		_, _ = client.Invoke(&methods.SendMessage{
			ChatID: update.Chat.ID,
			Text:   "You are already converting a sticker pack. Wait until it's done.",
		})
	}
}
