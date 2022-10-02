package utils

import (
	botType "Stickers2Emoji/types"
	"bytes"
	"github.com/Squirrel-Network/gobotapi"
	"github.com/Squirrel-Network/gobotapi/types"
	"strings"
	"sync"
)

func DownloadStickers(client *gobotapi.Client, stickerSet types.StickerSet) ([]botType.StickerReader, error) {
	var stickers []botType.StickerReader
	var waitGroup sync.WaitGroup
	var failedError error
	for _, sticker := range stickerSet.Stickers {
		if sticker.IsAnimated || sticker.IsVideo {
			continue
		}
		waitGroup.Add(1)
		go func(sticker types.Sticker) {
			defer waitGroup.Done()
			if failedError == nil {
			tryDownload:
				data, err := client.DownloadBytes(sticker.FileID, nil)
				if err != nil {
					if strings.Contains(err.Error(), "GOAWAY") {
						goto tryDownload
					}
					failedError = err
				}
				stickers = append(stickers, botType.StickerReader{
					Emoji: sticker.Emoji,
					Data:  bytes.NewReader(data),
				})
			}
		}(sticker)
	}
	waitGroup.Wait()
	if failedError != nil {
		return nil, failedError
	}
	return stickers, nil
}
