package game2048

import (
	"image/color"
)

var (
	backgroundColor      = color.RGBA{R: 0xfa, G: 0xf8, B: 0xef, A: 0xff}
	frameColor           = color.RGBA{R: 0xbb, G: 0xad, B: 0xa0, A: 0xff}
	tilesColor           = color.RGBA{R: 0xf9, G: 0xf6, B: 0xf2, A: 0xff}
	backgroundColor0     = color.NRGBA{R: 0xee, G: 0xe4, B: 0xda, A: 0x59}
	backgroundColor2     = color.RGBA{R: 0xee, G: 0xe4, B: 0xda, A: 0xff}
	backgroundColor4     = color.RGBA{R: 0xed, G: 0xe0, B: 0xc8, A: 0xff}
	backgroundColor8     = color.RGBA{R: 0xf2, G: 0xb1, B: 0x79, A: 0xff}
	backgroundColor16    = color.RGBA{R: 0xf5, G: 0x95, B: 0x63, A: 0xff}
	backgroundColor32    = color.RGBA{R: 0xf6, G: 0x7c, B: 0x5f, A: 0xff}
	backgroundColor64    = color.RGBA{R: 0xf6, G: 0x5e, B: 0x3b, A: 0xff}
	backgroundColor128   = color.RGBA{R: 0xed, G: 0xcf, B: 0x72, A: 0xff}
	backgroundColor256   = color.RGBA{R: 0xed, G: 0xcc, B: 0x61, A: 0xff}
	backgroundColor512   = color.RGBA{R: 0xed, G: 0xc8, B: 0x50, A: 0xff}
	backgroundColor1024  = color.RGBA{R: 0xed, G: 0xc5, B: 0x3f, A: 0xff}
	backgroundColor2048  = color.RGBA{R: 0xed, G: 0xc2, B: 0x2e, A: 0xff}
	backgroundColor4096  = color.NRGBA{R: 0xa3, G: 0x49, B: 0xa4, A: 0x7f}
	backgroundColor8192  = color.NRGBA{R: 0xa3, G: 0x49, B: 0xa4, A: 0xb2}
	backgroundColor16384 = color.NRGBA{R: 0xa3, G: 0x49, B: 0xa4, A: 0xcc}
	backgroundColor32768 = color.NRGBA{R: 0xa3, G: 0x49, B: 0xa4, A: 0xe5}
	backgroundColor65536 = color.NRGBA{R: 0xa3, G: 0x49, B: 0xa4, A: 0xff}
)

func tileColor() color.Color {
	return tilesColor
}

func tileBackgroundColor(value int) color.Color {
	switch value {
	case 0:
		return backgroundColor0
	case 2:
		return backgroundColor2
	case 4:
		return backgroundColor4
	case 8:
		return backgroundColor8
	case 16:
		return backgroundColor16
	case 32:
		return backgroundColor32
	case 64:
		return backgroundColor64
	case 128:
		return backgroundColor128
	case 256:
		return backgroundColor256
	case 512:
		return backgroundColor512
	case 1024:
		return backgroundColor1024
	case 2048:
		return backgroundColor2048
	case 4096:
		return backgroundColor4096
	case 8192:
		return backgroundColor8192
	case 16384:
		return backgroundColor16384
	case 32768:
		return backgroundColor32768
	case 65536:
		return backgroundColor65536
	}
	panic("not reach")
}
