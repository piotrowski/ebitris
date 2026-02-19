package render

import (
	"bytes"
	"sync"

	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var defaultFontSource *text.GoTextFaceSource

var once sync.Once

type fontSize int

const (
	FontMedium fontSize = 14
	FontLarge  fontSize = 24
)

func GetDefaultFont(size fontSize) text.Face {
	once.Do(func() {
		fontSource, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
		if err != nil {
			panic(err)
		}
		defaultFontSource = fontSource
	})

	return &text.GoTextFace{
		Source: defaultFontSource,
		Size:   float64(size),
	}
}
