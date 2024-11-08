// nolint: mnd
package main

import (
	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers"
	"github.com/xuender/kit/los"
)

func main() {
	can := canvas.New(100, 100)
	ctx := canvas.NewContext(can)

	triangle := los.Must(canvas.ParseSVGPath("L60 0L30 60z"))
	ctx.SetFillColor(canvas.Mediumseagreen)
	ctx.DrawPath(20, 20, triangle)

	los.Must0(renderers.Write("getting-started.png", can, canvas.DPMM(3.2)))
}
