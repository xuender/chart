package tile

import (
	"image/color"

	"github.com/tdewolff/canvas"
)

type Pen func(*canvas.Context, *canvas.Path, *Tile)

type Tile struct {
	density     float64
	size        float64
	colors      []color.Color
	strokeColor color.Color
	strokeWidth float64
	pen         Pen
}

func newTile() *Tile {
	return &Tile{
		density:     _defaultDensity,
		strokeColor: canvas.Transparent,
	}
}

func (p *Tile) WithDensity(density float64) *Tile {
	p.density = density

	return p
}

func (p *Tile) WithSize(size float64) *Tile {
	p.size = size

	return p
}

func (p *Tile) WithStrokeWidth(width float64) *Tile {
	p.strokeWidth = width

	return p
}

func (p *Tile) WithStrokeColor(color color.Color) *Tile {
	p.strokeColor = color

	return p
}

func (p *Tile) WithColor(colors ...color.Color) *Tile {
	if len(colors) == 0 {
		return p
	}

	p.colors = colors

	return p
}

func (p *Tile) getColor(idx int) color.Color {
	return p.colors[idx%len(p.colors)]
}

func (p *Tile) getSize(clip *canvas.Path) float64 {
	if p.size == _noneSize {
		return clip.Bounds().W * p.density
	}

	return p.size
}

func (p *Tile) Tile(ctx *canvas.Context, clip *canvas.Path) {
	if p.pen != nil {
		p.pen(ctx, clip, p)
	}
}
