// nolint: mnd
package main

import "github.com/xuender/chart"

func main() {
	ch1 := chart.New[int](800, 600, chart.WithDisplayBorder)

	ch1.Title = "你好 Chart"
	ch1.WriteFile("title.jpg")
}
