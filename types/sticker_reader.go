package types

import "bytes"

type StickerReader struct {
	Emoji string
	Data  *bytes.Reader
}
