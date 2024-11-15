package tile

import (
	"image/color"
	"math"

	"github.com/tdewolff/canvas"
)

func NewCairo() *Tile {
	ret := newTile()
	ret.colors = []color.Color{
		canvas.Sandybrown,
		canvas.Steelblue,
		canvas.Lightskyblue,
		canvas.Peachpuff,
	}
	ret.pen = Cairo

	return ret
}

func Cairo(ctx *canvas.Context, clip *canvas.Path, tile *Tile) {
	size := tile.getSize(clip)
	move := size * math.Tan(_angle30/_angle180*math.Pi)
	pentagon := &canvas.Path{}
	pentagon.MoveTo(size-move, 0)
	pentagon.LineTo(size, size)
	pentagon.LineTo(0, size+move)
	pentagon.LineTo(-size, size)
	pentagon.LineTo(-size+move, 0)
	pentagon.Close()

	cell := canvas.PrimitiveCell(
		canvas.Point{X: size + size, Y: 0 - size - size},
		canvas.Point{X: size + size, Y: size + size},
	)
	matrices := []canvas.Matrix{
		canvas.Identity,
		canvas.Identity.RotateAbout(_angle90, size, size),
		canvas.Identity.RotateAbout(_angle180, size, size),
		canvas.Identity.RotateAbout(_angle270, size, size),
	}

	ctx.SetStrokeColor(tile.strokeColor)
	ctx.SetStrokeWidth(tile.strokeWidth)

	for idx, matrix := range matrices {
		path := pentagon.Transform(matrix).Tile(clip, cell)

		ctx.SetFillColor(tile.getColor(idx))
		ctx.DrawPath(0.0, 0.0, path)
	}
}
