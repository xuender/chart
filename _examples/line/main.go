// nolint: mnd
package main

import (
	"slices"

	"github.com/xuender/chart"
)

func main() {
	ch1 := chart.New[int, int]().
		WithDisplayBorder().
		WidthDebug().
		WithSize(600, 480).
		Build()

	ch1.Title = "Line Demo"

	ch1.Column("line1", slices.All([]int{3, 10}))
	ch1.Column("line2", slices.All([]int{10, 3}))
	ch1.WriteFile("line.png")
}
