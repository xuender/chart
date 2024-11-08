package chart

import "github.com/tdewolff/canvas"

const (
	DefaultBorderWidth = 2
	DefaultLineWidth   = 2
	DefaultTitleSize   = 80
	DefaultWidth       = 800
	DefaultHeight      = 600
)

// nolint: gochecknoglobals, mnd
var (
	DebugTitle = canvas.RGBA(255, 0, 0, 0.2)
	DebugChart = canvas.RGBA(0, 0, 255, 0.2)
)
