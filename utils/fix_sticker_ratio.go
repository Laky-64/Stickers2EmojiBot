package utils

import (
	"Stickers2Emoji/types"
	"golang.org/x/image/draw"
	"image"
)

func FixStickerRatio(src image.Image) image.Image {
	edges := AnalyzeImage(src)
	content := src.(types.SubImage).SubImage(edges.GetContentBox())
	maxSize := content.Bounds().Max
	var size int
	if content.Bounds().Max.X > content.Bounds().Max.Y {
		size = maxSize.X
	} else {
		size = maxSize.Y
	}
	dst := image.NewRGBA(image.Rect(0, 0, size, size))
	x := (size - maxSize.X) >> 1
	y := (size - maxSize.Y) >> 1
	draw.Draw(
		dst,
		image.Rect(x,y, maxSize.X+x, maxSize.Y+y),
		content, image.Point{}, draw.Src,
	)
	return dst
}

func AnalyzeImage(src image.Image) *types.Edge {
	bounds := src.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	srcRgba := image.NewRGBA(bounds)
	draw.Copy(srcRgba, image.Point{}, src, bounds, draw.Src, nil)
	edges := types.NewEdge(width, height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			idxS := (y*width + x) * 4
			pixel := srcRgba.Pix[idxS : idxS+4]
			edges.Trace(x, y, pixel)
		}
	}
	return edges
}
