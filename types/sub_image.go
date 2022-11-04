package types

import "image"

type SubImage interface {
	SubImage(r image.Rectangle) image.Image
}
