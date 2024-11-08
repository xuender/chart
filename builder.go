package chart

import (
	"cmp"
	"iter"

	"github.com/tdewolff/canvas"
)

type Builder[K cmp.Ordered, V Number] struct {
	chart *Chart[K, V]
}

func New[K cmp.Ordered, V Number]() *Builder[K, V] {
	return &Builder[K, V]{
		chart: &Chart[K, V]{
			Fonts:       DefaultFonts,
			data:        make(map[string]iter.Seq2[K, V]),
			Colors:      DefaultColors,
			colors:      make(map[string]int),
			borderWidth: DefaultBorderWidth,
			lineWidth:   DefaultLineWidth,
			titleSize:   DefaultTitleSize,
			width:       DefaultWidth,
			height:      DefaultHeight,
		},
	}
}

func (p *Builder[K, V]) Build() *Chart[K, V] {
	p.chart.canvas = canvas.New(p.chart.width, p.chart.height)
	p.chart.ctx = canvas.NewContext(p.chart.canvas)
	p.chart.layout = DefaultLayout(
		p.chart.width,
		p.chart.height,
		p.chart.borderWidth,
	)

	return p.chart
}

func (p *Builder[K, V]) WidthDebug() *Builder[K, V] {
	p.chart.debug = true

	return p
}

func (p *Builder[K, V]) WithDisplayBorder() *Builder[K, V] {
	p.chart.displayBorder = true

	return p
}

func (p *Builder[K, V]) WithSize(width, height float64) *Builder[K, V] {
	p.chart.width = width
	p.chart.height = height

	return p
}

func (p *Builder[K, V]) WithBorderWidth(width float64) *Builder[K, V] {
	p.chart.borderWidth = width

	return p
}
