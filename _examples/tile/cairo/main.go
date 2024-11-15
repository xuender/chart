package main

import (
	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers"
	"github.com/xuender/chart/tile"
)

func main() {
	W := 16.0
	H := 16.0
	c := canvas.New(W, H)
	ctx := canvas.NewContext(c)
	ctx.SetFillColor(canvas.White)
	ctx.DrawPath(0, 0, canvas.Rectangle(W, H))

	clip := canvas.RoundedRectangle(10.0, 10.0, 2.0).Translate(3, 3)
	cairo := tile.NewCairo().WithDensity(0.2)
	cairo.Tiling(ctx, clip)

	renderers.Write("cairo.png", c, canvas.DPMM(20.0))
	renderers.Write("cairo.svg", c)
}
