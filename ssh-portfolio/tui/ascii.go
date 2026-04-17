package tui

import (
	"bytes"
	"fmt"
	"image"
	_ "image/jpeg"
	"strings"

	_ "embed"
)

//go:embed portrait.jpg
var portraitJPG []byte

const portraitCols = 40

// GetAscii renders portrait.jpg as a half-block true-colour image.
// Each character row covers 2 pixel rows: top → background, bottom → foreground, char = ▄.
func GetAscii(_ Theme) string {
	img, _, err := image.Decode(bytes.NewReader(portraitJPG))
	if err != nil {
		return ""
	}

	b := img.Bounds()
	srcW := b.Max.X - b.Min.X
	srcH := b.Max.Y - b.Min.Y

	// Preserve aspect ratio assuming terminal chars are ~2:1 (h:w).
	targetW := portraitCols
	targetH := targetW * srcH / srcW / 2
	if targetH < 1 {
		targetH = 1
	}
	pixH := targetH * 2

	scaled := resizeNearest(img, targetW, pixH)

	var sb strings.Builder
	for row := 0; row < targetH; row++ {
		for col := 0; col < targetW; col++ {
			tR, tG, tB, _ := scaled.At(col, row*2).RGBA()
			bR, bG, bB, _ := scaled.At(col, row*2+1).RGBA()
			fmt.Fprintf(&sb, "\x1b[38;2;%d;%d;%dm\x1b[48;2;%d;%d;%dm▄",
				bR>>8, bG>>8, bB>>8,
				tR>>8, tG>>8, tB>>8)
		}
		if row < targetH-1 {
			sb.WriteString("\x1b[0m\n")
		} else {
			sb.WriteString("\x1b[0m")
		}
	}

	return sb.String()
}

func resizeNearest(src image.Image, w, h int) image.Image {
	dst := image.NewRGBA(image.Rect(0, 0, w, h))
	b := src.Bounds()
	sw := b.Max.X - b.Min.X
	sh := b.Max.Y - b.Min.Y
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			sx := x*sw/w + b.Min.X
			sy := y*sh/h + b.Min.Y
			dst.Set(x, y, src.At(sx, sy))
		}
	}
	return dst
}
