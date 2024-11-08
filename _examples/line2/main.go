// nolint: mnd
package main

import (
	"slices"

	"github.com/xuender/chart"
)

func main() {
	ch1 := chart.New[int, int]().
		WithBorderWidth(20).
		WithDisplayBorder().
		Build()

	ch1.Title = "Line2 Demo"

	ch1.Column("line1", slices.All([]int{3, 5, 7, 11, 13}))
	ch1.Column("line2", slices.All([]int{2, 4, 6, 8, 9}))
	ch1.WriteFile("line2.png")
}
