// nolint: mnd
package main

import "github.com/xuender/chart"

func main() {
	ch1 := chart.New[int, int]().
		WithDisplayBorder().
		WithBorderWidth(40).
		WidthDebug().
		Build()

	ch1.WriteFile("border.jpg")
}
