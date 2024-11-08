package chart

import (
	"image/color"
	"iter"
	"log/slog"
	"maps"
	"slices"

	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers"
	"github.com/xuender/chart/draw"
)

type Chart[V Number] struct {
	Title  string
	Fonts  *Fonts
	canvas *canvas.Canvas
	ctx    *canvas.Context
	data   map[string]iter.Seq2[int, V]
	Colors []color.RGBA
	colors map[string]int
	layout *Layout

	displayBorder bool
	borderWidth   float64
	lineWidth     float64
	titleSize     float64
}

func New[V Number](width, height float64, options ...Option[V]) *Chart[V] {
	can := canvas.New(width, height)
	chart := &Chart[V]{
		canvas:      can,
		ctx:         canvas.NewContext(can),
		Fonts:       DefaultFonts,
		data:        make(map[string]iter.Seq2[int, V]),
		Colors:      DefaultColors,
		colors:      make(map[string]int),
		layout:      DefaultLayout(width, height),
		borderWidth: DefaultBorderWidth,
		lineWidth:   DefaultLineWidth,
		titleSize:   DefaultTitleSize,
	}

	for _, option := range options {
		option(chart)
	}

	return chart
}

func (p *Chart[V]) draw() error {
	p.drawBackgroup()

	if p.Title != "" {
		p.drawTitle()
	}

	if p.displayBorder {
		p.drawBorder()
	}

	p.drawData()

	return nil
}

func (p *Chart[V]) drawBorder() {
	draw.Line(p.ctx, canvas.Black, p.borderWidth, []canvas.Point{
		{X: p.layout.Chart.X, Y: p.layout.Chart.Y},
		{X: p.layout.Chart.W, Y: p.layout.Chart.Y},
		{X: p.layout.Chart.W, Y: p.layout.Chart.H},
		{X: p.layout.Chart.X, Y: p.layout.Chart.H},
		{X: p.layout.Chart.X, Y: p.layout.Chart.Y},
	})
}

func (p *Chart[V]) drawData() {
	lines := map[string][]canvas.Point{}
	for name, data := range p.data {
		lines[name] = Line(data)
	}

	rect := LinesRect(maps.Values(lines))

	slog.Info("rect", "w", rect.W, "h", rect.H, "x", rect.X, "y", rect.Y)
	slog.Info("chart", "w", p.layout.Chart.W, "h", p.layout.Chart.H, "x", p.layout.Chart.X, "y", p.layout.Chart.Y)

	idxRatio := p.layout.Chart.W / rect.W
	valRatio := (p.layout.Chart.Y - p.layout.Chart.H) / rect.H

	slog.Info("data", "lines", lines)

	for name, line := range lines {
		for idx, val := range line {
			val.X *= idxRatio
			val.Y *= valRatio
			val.Y -= rect.Y
			val.X += (idxRatio / 4) // nolint: mnd
			val.X += p.layout.Chart.X
			val.Y += p.layout.Chart.H
			line[idx] = val
		}

		p.drawLine(name, line)
	}

	slog.Info("data", "lines", lines)
}

func (p *Chart[V]) drawLine(name string, data []canvas.Point) {
	idx, has := p.colors[name]
	if !has {
		idx = len(p.colors)
		p.colors[name] = idx
	}

	draw.Line(p.ctx, p.Colors[idx], p.lineWidth, data)
}

func (p *Chart[V]) drawBackgroup() {
	background := canvas.Rectangle(p.ctx.Width()+1, p.ctx.Height()+1)

	p.ctx.SetFillColor(canvas.White)
	p.ctx.DrawPath(0, 0, background)
}

func (p *Chart[V]) drawTitle() {
	face := p.Fonts.Face("default", p.titleSize)
	draw.Text(p.ctx, p.Title, face, p.layout.Title)
}

func (p *Chart[V]) Column(name string, data []V) {
	p.data[name] = slices.All(data)
}

func (p *Chart[V]) WriteFile(file string) error {
	if err := p.draw(); err != nil {
		return err
	}

	return renderers.Write(file, p.canvas)
}
