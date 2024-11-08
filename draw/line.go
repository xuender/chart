package draw

import (
	"image/color"

	"github.com/tdewolff/canvas"
)

func Line(ctx *canvas.Context, strokeColor color.Color, strokeWidth float64, points []canvas.Point) {
	ctx.SetStrokeColor(strokeColor)
	ctx.SetStrokeWidth(strokeWidth)

	for idx, point := range points {
		if idx == 0 {
			ctx.MoveTo(point.X, point.Y)
		} else {
			ctx.LineTo(point.X, point.Y)
		}
	}

	ctx.Stroke()
}
