package chart

type Option[V Number] func(*Chart[V])

func WithDisplayBorder[V Number](chart *Chart[V]) {
	chart.displayBorder = true
}

func WithBorderWidth[V Number](width float64) func(chart *Chart[V]) {
	return func(chart *Chart[V]) {
		chart.borderWidth = width
	}
}
