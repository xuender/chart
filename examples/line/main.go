// nolint: mnd
package main

import "github.com/xuender/chart"

func main() {
	ch1 := chart.New[int](800, 600, chart.WithDisplayBorder)

	ch1.Column("one", []int{1, 3, 2, 5, 4})
	ch1.Column("two", []int{10, 30, 20, 50, 40})
	ch1.Column("three", []int{13, 10, 26, 15, 24})
	ch1.WriteFile("line.jpg")
}
