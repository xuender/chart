package draw

import (
	"image/color"

	"github.com/tdewolff/canvas"
)

func Line(ctx *canvas.Context, colo color.Color, width float64, points []canvas.Point) {
	ctx.SetStrokeColor(colo)
	ctx.SetStrokeWidth(width)

	for idx, point := range points {
		if idx == 0 {
			ctx.MoveTo(point.X, point.Y)
		} else {
			ctx.LineTo(point.X, point.Y)
		}
	}

	ctx.Stroke()
}
