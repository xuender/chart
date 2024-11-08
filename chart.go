package chart

import (
	"cmp"
	"fmt"
	"image/color"
	"iter"
	"log/slog"
	"maps"

	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers"
	"github.com/xuender/chart/draw"
)

type Chart[K cmp.Ordered, V Number] struct {
	Title  string
	Fonts  *Fonts
	canvas *canvas.Canvas
	ctx    *canvas.Context
	data   map[string]iter.Seq2[K, V]
	Colors []color.RGBA
	colors map[string]int
	layout *Layout

	width, height float64
	displayBorder bool
	borderWidth   float64
	lineWidth     float64
	titleSize     float64
	debug         bool
}

func (p *Chart[K, V]) draw() error {
	p.drawBackgroup()

	if p.debug {
		p.drawDebug()
	}

	if p.Title != "" {
		p.drawTitle()
	}

	if p.displayBorder {
		p.drawBorder()
	}

	p.drawData()

	return nil
}

func (p *Chart[K, V]) drawBorder() {
	p.ctx.SetStrokeColor(canvas.Black)
	p.ctx.SetStrokeWidth(p.borderWidth)

	border := p.borderWidth / 2 // nolint

	p.ctx.DrawPath(p.layout.Main.X+border, p.layout.Main.Y+border,
		canvas.Rectangle(p.layout.Main.W-p.borderWidth, p.layout.Main.H-p.borderWidth),
	)
}

func (p *Chart[K, V]) drawData() {
	lineMap := map[string]map[K]V{}
	for name, data := range p.data {
		lineMap[name] = maps.Collect(data)
	}

	keys, lines := Lines(lineMap)

	slog.Info("draw", "keys", keys)

	rect := LinesRect(maps.Values(lines))
	rect.W = float64(len(keys) - 1)

	slog.Info("rect", "rect", rectString(rect))
	slog.Info("chart", "chart", rectString(p.layout.Chart))

	idxRatio := p.layout.Chart.W / rect.W
	valRatio := p.layout.Chart.H / (rect.H - rect.Y)

	slog.Info("data", "lines", lines)

	for name, line := range lines {
		for idx, val := range line {
			val.Y -= rect.Y
			val.X *= idxRatio
			val.Y *= valRatio
			val.X += p.layout.Chart.X
			val.Y += p.layout.Chart.Y
			line[idx] = val
		}

		p.drawLine(name, line)
	}

	slog.Info("data", "lines", lines)
}

func (p *Chart[K, V]) drawLine(name string, data []canvas.Point) {
	idx, has := p.colors[name]
	if !has {
		idx = len(p.colors)
		p.colors[name] = idx
	}

	draw.Line(p.ctx, p.Colors[idx], p.lineWidth, data)
}

func (p *Chart[K, V]) drawBackgroup() {
	background := canvas.Rectangle(p.ctx.Width(), p.ctx.Height())

	p.ctx.SetFillColor(canvas.White)
	p.ctx.DrawPath(0, 0, background)
}

func (p *Chart[K, V]) drawTitle() {
	face := p.Fonts.Face("default", p.titleSize)
	draw.Text(p.ctx, p.Title, face, p.layout.Title)
}

func rectString(rect canvas.Rect) string {
	return fmt.Sprintf("(%.0f,%.0f) (%.0f,%.0f)", rect.X, rect.Y, rect.W, rect.H)
}

func (p *Chart[K, V]) drawDebug() {
	slog.Info("debug", "title", rectString(p.layout.Title), "chart", rectString(p.layout.Main))
	p.ctx.SetFillColor(DebugTitle)
	p.ctx.DrawPath(p.layout.Title.X, p.layout.Title.Y, canvas.Rectangle(p.layout.Title.W, p.layout.Title.H))

	p.ctx.SetFillColor(DebugChart)

	p.ctx.DrawPath(
		p.layout.Chart.X, p.layout.Chart.Y,
		canvas.Rectangle(p.layout.Chart.W, p.layout.Chart.H),
	)
}

func (p *Chart[K, V]) Column(name string, seq iter.Seq2[K, V]) {
	p.data[name] = seq
}

func (p *Chart[K, V]) WriteFile(file string) error {
	if err := p.draw(); err != nil {
		return err
	}

	return renderers.Write(file, p.canvas)
}
