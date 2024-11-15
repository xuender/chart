package tile

import (
	"image/color"
	"math"

	"github.com/tdewolff/canvas"
)

const RightAngle = 90.0

type Cairo struct {
	density float64
	colors  []color.Color
}

func NewCairo() *Cairo {
	defaultDensity := 0.1

	return &Cairo{
		density: defaultDensity,
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

func (p *Cairo) WithColor(colors ...color.Color) *Cairo {
	p.colors = append(colors, p.colors...)

	return p
}

func (p *Cairo) Tiling(ctx *canvas.Context, clip *canvas.Path) {
	size := clip.Bounds().W * p.density
	move := size * math.Tan(30.0/180.0*math.Pi)

	pentagon := &canvas.Path{}
	pentagon.MoveTo(size-move, 0.0)
	pentagon.LineTo(size, size)
	pentagon.LineTo(0.0, size+move)
	pentagon.LineTo(-size, size)
	pentagon.LineTo(-size+move, 0.0)
	pentagon.Close()

	cell := canvas.PrimitiveCell(
		canvas.Point{X: size + size, Y: 0 - size - size},
		canvas.Point{X: size + size, Y: size + size},
	)
	matrices := []canvas.Matrix{
		canvas.Identity,
		canvas.Identity.RotateAbout(RightAngle, size, size),
		canvas.Identity.RotateAbout(RightAngle+RightAngle, size, size),
		canvas.Identity.RotateAbout(RightAngle+RightAngle+RightAngle, size, size),
	}

	ctx.SetStrokeColor(canvas.Transparent)

	for idx, matrix := range matrices {
		path := pentagon.Transform(matrix).Tile(clip, cell)

		ctx.SetFillColor(p.colors[idx])
		ctx.DrawPath(0.0, 0.0, path)
	}
}
