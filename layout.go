// nolint: mnd
package chart

import (
	"log/slog"

	"github.com/tdewolff/canvas"
)

type Layout struct {
	Size  canvas.Size
	Title canvas.Rect
	Chart canvas.Rect
}

func DefaultLayout(width, height float64) *Layout {
	slog.Info("default layout", "width", width*0.9, "height", height*0.2)

	return &Layout{
		Size: canvas.Size{W: width, H: height},
		Title: canvas.Rect{
			X: 0, Y: height * 0.9,
			W: width, H: height * 0.1,
		},
		Chart: canvas.Rect{
			X: width * 0.1, Y: height * 0.8,
			W: width * 0.9, H: height * 0.2,
		},
	}
}
