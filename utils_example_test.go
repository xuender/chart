package chart_test

import (
	"fmt"

	"github.com/xuender/chart"
)

func ExampleShuffleSlice() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(chart.ShuffleSlice(slice, 2))
	fmt.Println(chart.ShuffleSlice(slice, 3))
	fmt.Println(chart.ShuffleSlice(slice, 5))

	// Output:
	// [0 2 4 6 8 1 3 5 7 9]
	// [0 3 6 9 1 4 7 2 5 8]
	// [0 5 1 6 2 7 3 8 4 9]
}
