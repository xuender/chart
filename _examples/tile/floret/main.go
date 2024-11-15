package main

import (
	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers"
	"github.com/xuender/chart/tile"
)

func main() {
	W := 29.0
	H := 29.0
	c := canvas.New(W, H)
	ctx := canvas.NewContext(c)
	ctx.SetFillColor(canvas.White)
	ctx.DrawPath(0, 0, canvas.Rectangle(W, H))

	cairo := tile.NewFloret()
	clip := canvas.RoundedRectangle(10.0, 10.0, 2.0).Translate(3, 16)
	cairo.Tile(ctx, clip)

	clip = canvas.RoundedRectangle(10.0, 10.0, 2.0).Translate(16, 16)
	cairo.WithStrokeWidth(0.1).WithStrokeColor(canvas.Black)
	cairo.Tile(ctx, clip)

	clip = canvas.RoundedRectangle(10.0, 10.0, 2.0).Translate(3, 3)
	cairo.WithColor(canvas.Red, canvas.Blue, canvas.Green)
	cairo.Tile(ctx, clip)

	clip = canvas.RoundedRectangle(10.0, 10.0, 2.0).Translate(16, 3)
	cairo.WithDensity(0.2).WithStrokeColor(canvas.Transparent)
	cairo.Tile(ctx, clip)

	renderers.Write("floret.png", c, canvas.DPMM(20.0))
}
