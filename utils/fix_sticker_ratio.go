package utils

import (
	"Stickers2Emoji/consts"
	"Stickers2Emoji/types"
	"image"
)

type SSP types.SideScorePair

// This function finds the side of the image which contains more transparent
// pixels so not to crop the relevant parts
func findSide(src image.Rectangle, maxX, maxY, minX, minY int) int {
	sideScorePairs := make([]SSP, 4)

	sideScorePairs[consts.BoundCheckLeft] = SSP{Side: consts.BoundCheckLeft}
	sideScorePairs[consts.BoundCheckRight] = SSP{Side: consts.BoundCheckRight}
	sideScorePairs[consts.BoundCheckTop] = SSP{Side: consts.BoundCheckTop}
	sideScorePairs[consts.BoundCheckBottom] = SSP{Side: consts.BoundCheckBottom}

	isEmptyAt := func(side int, absPos int) bool {
		switch side {
		case consts.BoundCheckLeft:
			_, _, _, a := src.At(minX, absPos).RGBA()

			return a > 0
		case consts.BoundCheckRight:
			_, _, _, a := src.At(maxX, absPos).RGBA()

			return a > 0
		case consts.BoundCheckTop:
			_, _, _, a := src.At(absPos, minY).RGBA()

			return a > 0
		case consts.BoundCheckBottom:
			_, _, _, a := src.At(absPos, maxY).RGBA()

			return a > 0
		}

		panic("This should never happen")
	}

	for y := minY; y <= maxY; y++ {
		if isEmptyAt(consts.BoundCheckLeft, y) {
			sideScorePairs[consts.BoundCheckLeft].Score++
		}

		if isEmptyAt(consts.BoundCheckRight, y) {
			sideScorePairs[consts.BoundCheckRight].Score++
		}
	}

	for x := minX; x <= maxX; x++ {
		if isEmptyAt(consts.BoundCheckTop, x) {
			sideScorePairs[consts.BoundCheckTop].Score++
		}

		if isEmptyAt(consts.BoundCheckBottom, x) {
			sideScorePairs[consts.BoundCheckBottom].Score++
		}
	}

	return Reduce(sideScorePairs, func(a, b SSP) SSP {
		if a.Score > b.Score {
			return b
		} else {
			return a
		}
	}, sideScorePairs[0]).Side
}

func FixStickerRatio(src image.Rectangle) image.Rectangle {
	maxX := src.Bounds().Max.X
	maxY := src.Bounds().Max.Y
	minX := src.Bounds().Min.X
	minY := src.Bounds().Min.Y

	// No crop to be made
	if maxX-minX == consts.EmojiSize && maxY-minY == consts.EmojiSize {
		return src
	}

	side := findSide(src, maxX, maxY, minX, minY)

	switch side {
	case consts.BoundCheckLeft:
		return image.Rect(minX+1, src.Bounds().Min.X, maxX, maxY)
	case consts.BoundCheckRight:
		return image.Rect(minX, minY, maxX-1, maxY)
	case consts.BoundCheckTop:
		return image.Rect(minX, minY+1, maxX, maxY)
	case consts.BoundCheckBottom:
		return image.Rect(minX, minY, maxX, maxY-1)
	}

	panic("This should never happen")
}
