package tile

import (
	"image/color"
	"math"

	"github.com/tdewolff/canvas"
)

const RightAngle = 90.0

type Cairo struct {
	density float64
	size    float64
	colors  []color.Color
}

func NewCairo() *Cairo {
	defaultDensity := 0.1
	defaultSize := 0.0

	return &Cairo{
		density: defaultDensity,
		size:    defaultSize,
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
	if p.size == 0.0 {
		p.size = clip.Bounds().W * p.density
	}

	move := p.size * math.Tan(30.0/180.0*math.Pi)

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
		canvas.Identity.RotateAbout(RightAngle, p.size, p.size),
		canvas.Identity.RotateAbout(RightAngle+RightAngle, p.size, p.size),
		canvas.Identity.RotateAbout(RightAngle+RightAngle+RightAngle, p.size, p.size),
	}

	ctx.SetStrokeColor(canvas.Transparent)

	for idx, matrix := range matrices {
		path := pentagon.Transform(matrix).Tile(clip, cell)

		ctx.SetFillColor(p.colors[idx])
		ctx.DrawPath(0.0, 0.0, path)
	}
}
