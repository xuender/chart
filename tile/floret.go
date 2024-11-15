package tile

import (
	"image/color"
	"math"

	"github.com/tdewolff/canvas"
)

func NewFloret() *Tile {
	ret := newTile()
	ret.colors = []color.Color{canvas.Hex("#c5f0e9")}
	ret.strokeColor = canvas.Black
	ret.strokeWidth = 0.02
	ret.pen = Floret

	return ret
}

func Floret(ctx *canvas.Context, clip *canvas.Path, tile *Tile) {
	size := tile.getSize(clip)
	pdx, pdy := math.Sincos(_angle30 * math.Pi / _angle180)
	pdx *= size
	pdy *= size

	pentagon := &canvas.Path{}
	pentagon.LineTo(pdx, pdy)
	pentagon.LineTo(_quarter*size, _oneAndHalf*pdy)
	pentagon.LineTo(-1*_quarter*size, _oneAndHalf*pdy)
	pentagon.LineTo(-pdx, pdy)
	pentagon.Close()

	width := size*_double + pdx*_half
	height := pdy * _half
	distance := math.Sqrt(width*width + height*height)
	cell := canvas.PrimitiveCell(
		canvas.Point{X: math.Sqrt(_threeQuarters) * distance, Y: -1 * _half * distance},
		canvas.Point{X: math.Sqrt(_threeQuarters) * distance, Y: _half * distance},
	)

	ctx.SetStrokeColor(tile.strokeColor)
	ctx.SetStrokeWidth(tile.strokeWidth)

	theta0 := _angle30 + math.Atan(height/width)*_angle180/math.Pi
	place := 6

	for idx := range place {
		path := pentagon.Transform(canvas.Identity.Rotate(theta0+_angle60*float64(idx))).
			Tile(clip, cell)

		ctx.SetFillColor(tile.getColor(idx))
		ctx.DrawPath(0.0, 0.0, path)
	}
}
