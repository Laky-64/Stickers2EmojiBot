package types

import (
	"Stickers2Emoji/consts"
	"image"
)

type Edge struct {
	pixels     map[int]*uint8
	Width      int
	Height     int
	Horizontal bool
}

func NewEdge(width, height int) *Edge {
	return &Edge{
		pixels:     make(map[int]*uint8),
		Width:      width,
		Height:     height,
		Horizontal: float64(width)/float64(height) > 1,
	}
}

func (e *Edge) Trace(x, y int, pixel []uint8) {
	cursor := e.cursor(x, y)
	alpha := 255 - pixel[3]
	e.pixels[cursor] = &alpha
}

func (e *Edge) cursor(x, y int) int {
	return y*e.Width + x
}

func (e *Edge) GetContentBox() image.Rectangle {
	top, bottom := e.getStartEndEdge(e.isColumnTransparent, e.Height)
	left, right := e.getStartEndEdge(e.isLineTransparent, e.Width)
	return image.Rect(left, top, right, bottom)
}

func (e *Edge) getStartEndEdge(f func(int) bool, size int) (int, int) {
	var start, end int
	for i := 0; i < size; i++ {
		if !f(i) {
			start = i
			break
		}
	}
	for i := size - 1; i >= 0; i-- {
		if !f(i) {
			end = i
			break
		}
	}
	return start, end
}

func (e *Edge) isLineTransparent(x int) bool {
	var score int
	for y := 0; y < e.Height; y++ {
		if pixel := e.pixels[e.cursor(x, y)]; pixel != nil {
			score += int(*pixel)
		}
	}
	score /= e.Height
	return float64(score)*100/255 > float64(consts.SensibilityBorder)
}

func (e *Edge) isColumnTransparent(y int) bool {
	var score int
	for x := 0; x < e.Width; x++ {
		if pixel := e.pixels[e.cursor(x, y)]; pixel != nil {
			score += int(*pixel)
		}
	}
	score /= e.Width
	return float64(score)*100/255 > float64(consts.SensibilityBorder)
}
