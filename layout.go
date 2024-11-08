// nolint: mnd
package chart

import (
	"log/slog"

	"github.com/tdewolff/canvas"
)

type Layout struct {
	Size  canvas.Size
	Title canvas.Rect
	Main  canvas.Rect
	Chart canvas.Rect
}

func DefaultLayout(width, height, borderWidth float64) *Layout {
	slog.Info("default layout", "width", width, "height", height)

	ret := &Layout{
		Size: canvas.Size{W: width, H: height},
		Title: canvas.Rect{
			X: 0, Y: height * 0.9,
			W: width, H: height * 0.1,
		},
		Main: canvas.Rect{
			X: width * 0.1, Y: height * 0.1,
			W: width * 0.8, H: height * 0.8,
		},
	}

	ret.Chart = canvas.Rect{
		X: ret.Main.X + borderWidth,
		Y: ret.Main.Y + borderWidth,
		W: ret.Main.W - borderWidth - borderWidth,
		H: ret.Main.H - borderWidth - borderWidth,
	}

	return ret
}
