package chart

import (
	"iter"
	"math"

	"github.com/tdewolff/canvas"
)

func Line[V Number](data iter.Seq2[int, V]) []canvas.Point {
	nums := []canvas.Point{}

	for idx, val := range data {
		idxFloat := float64(idx)
		valFloat := float64(val)
		nums = append(nums, canvas.Point{X: idxFloat, Y: valFloat})
	}

	return nums
}

func LinesRect(data iter.Seq[[]canvas.Point]) canvas.Rect {
	idxMin, idxMax := 0.0, 0.0
	valMin := math.MaxFloat64
	valMax := math.MaxFloat64 * -1

	for items := range data {
		if float64(len(items)) > idxMax {
			idxMax = float64(len(items))
		}

		for _, item := range items {
			if item.Y < valMin {
				valMin = item.Y
			}

			if item.Y > valMax {
				valMax = item.Y
			}
		}
	}

	return canvas.Rect{X: idxMin, Y: valMin, W: idxMax, H: valMax}
}
