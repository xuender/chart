package tile

import (
	"image/color"
	"math"

	"github.com/tdewolff/canvas"
)

const (
	Angle30  = 30.0
	Angle90  = 90.0
	Angle180 = 180.0
	Angle270 = 270.0
	NoneSize = 0.0
)

type Cairo struct {
	density float64
	size    float64
	colors  []color.Color
}

func NewCairo() *Cairo {
	defaultDensity := 0.1

	return &Cairo{
		density: defaultDensity,
		size:    NoneSize,
		colors: []color.Color{
			canvas.Sandybrown,
			canvas.Steelblue,
			canvas.Lightskyblue,
			canvas.Peachpuff,
		},
	}
}

func (p *Cairo) WithDensity(density float64) *Cairo {
	p.density = density

	return p
}

func (p *Cairo) WithSize(size float64) *Cairo {
	p.size = size

	return p
}

func (p *Cairo) WithColor(colors ...color.Color) *Cairo {
	p.colors = append(colors, p.colors...)

	return p
}

func (p *Cairo) Tiling(ctx *canvas.Context, clip *canvas.Path) {
	if p.size == NoneSize {
		p.size = clip.Bounds().W * p.density
	}

	move := p.size * math.Tan(Angle30/Angle180*math.Pi)

	pentagon := &canvas.Path{}
	pentagon.MoveTo(p.size-move, 0.0)
	pentagon.LineTo(p.size, p.size)
	pentagon.LineTo(0.0, p.size+move)
	pentagon.LineTo(-p.size, p.size)
	pentagon.LineTo(-p.size+move, 0.0)
	pentagon.Close()

	cell := canvas.PrimitiveCell(
		canvas.Point{X: p.size + p.size, Y: 0 - p.size - p.size},
		canvas.Point{X: p.size + p.size, Y: p.size + p.size},
	)
	matrices := []canvas.Matrix{
		canvas.Identity,
		canvas.Identity.RotateAbout(Angle90, p.size, p.size),
		canvas.Identity.RotateAbout(Angle180, p.size, p.size),
		canvas.Identity.RotateAbout(Angle270, p.size, p.size),
	}

	ctx.SetStrokeColor(canvas.Transparent)

	for idx, matrix := range matrices {
		path := pentagon.Transform(matrix).Tile(clip, cell)

		ctx.SetFillColor(p.colors[idx])
		ctx.DrawPath(0.0, 0.0, path)
	}
}
