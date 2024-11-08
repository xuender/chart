package fonts

import (
	_ "embed"

	"github.com/tdewolff/canvas"
	"github.com/xuender/kit/los"
	"github.com/xuender/kit/oss"
)

//go:embed CangErYuMoW02-2.ttf
var _font []byte

// nolint: gochecknoglobals
var DefaultFonts = NewDefaultFonts()

type Fonts struct {
	fonts map[string]*canvas.FontFamily
}

func NewFonts() *Fonts {
	fonts := &Fonts{
		fonts: map[string]*canvas.FontFamily{},
	}

	return fonts
}

func NewDefaultFonts() *Fonts {
	fonts := NewFonts()

	font := fonts.Font("default")
	los.Must0(font.LoadFont(_font, 0, canvas.FontRegular))

	return fonts
}

func (p *Fonts) Font(name string) *canvas.FontFamily {
	font, has := p.fonts[name]
	if !has {
		font = canvas.NewFontFamily(name)
		p.fonts[name] = font
	}

	return font
}

func (p *Fonts) Load(name, file string, style canvas.FontStyle) error {
	font := p.Font(name)

	if oss.Exist(file) {
		return font.LoadFontFile(file, style)
	}

	return font.LoadSystemFont(file, style)
}

func (p *Fonts) Face(name string, size float64, args ...any) *canvas.FontFace {
	font := p.Font(name)

	return font.Face(size, args...)
}
