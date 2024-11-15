package main

import "github.com/xuender/chart"

func main() {
	ch1 := chart.New[int, int]().
		WidthDebug().
		WithSize(400, 300).
		Build()

	ch1.WriteFile("debug.jpg")
}
