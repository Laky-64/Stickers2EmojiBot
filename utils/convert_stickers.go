package utils

import (
	"Stickers2Emoji/consts"
	"Stickers2Emoji/types"
	"bytes"
	"golang.org/x/image/draw"
	"image"
	"sync"

	"github.com/chai2010/webp"
)

func ConvertStickers(stickersReader []types.StickerReader) ([]types.StickerBytes, error) {
	var convertedStickers []types.StickerBytes
	var failedError error
	var waitGroup sync.WaitGroup
	for _, sticker := range stickersReader {
		waitGroup.Add(1)
		go func(sticker types.StickerReader) {
			defer waitGroup.Done()
			if failedError == nil {
				src, err := webp.Decode(sticker.Data)
				if err != nil {
					failedError = err
					return
				}
				rectImage := image.Rect(0, 0, consts.EmojiSize, consts.EmojiSize)
				src = FixStickerRatio(src)
				dst := image.NewRGBA(rectImage)
				draw.CatmullRom.Scale(dst, rectImage, src, src.Bounds(), draw.Over, nil)
				var output bytes.Buffer
				err = webp.Encode(&output, dst, &webp.Options{
					Lossless: true,
					Exact:    true,
				})
				convertedStickers = append(convertedStickers, types.StickerBytes{
					Emoji: sticker.Emoji,
					Data:  output.Bytes(),
				})
			}
		}(sticker)
	}
	waitGroup.Wait()
	if failedError != nil {
		return nil, failedError
	}
	return convertedStickers, nil
}
