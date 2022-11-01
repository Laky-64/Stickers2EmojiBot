package utils

import (
	"Stickers2Emoji/consts"
	"Stickers2Emoji/types"
	"bytes"
	"image"
	"sync"

	"github.com/chai2010/webp"
	"golang.org/x/image/draw"
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
				var newW, newH int
				ratio := float64(src.Bounds().Max.X) / float64(src.Bounds().Max.Y)
				var rectImage image.Rectangle
				if ratio > 1 {
					newW = int(float64(consts.EmojiSize) * ratio)
					newH = consts.EmojiSize
					delta := (newW - consts.EmojiSize) / 2
					rectImage = image.Rect(delta, 0, newW-delta, newH)
				} else {
					newW = consts.EmojiSize
					newH = int(float64(consts.EmojiSize) / ratio)
					delta := (newH - consts.EmojiSize) / 2
					rectImage = image.Rect(0, delta, newW, newH-delta)
				}

				rectImage = FixStickerRatio(rectImage)

				dst := image.NewRGBA(rectImage)
				draw.CatmullRom.Scale(dst, image.Rect(0, 0, newW, newH), src, src.Bounds(), draw.Over, nil)
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
