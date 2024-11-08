package chart

import (
	"cmp"
	"iter"
	"maps"
	"math"
	"slices"

	"github.com/tdewolff/canvas"
	"github.com/xuender/flow"
	"github.com/xuender/flow/seq"
)

func Lines[K cmp.Ordered, V Number](lines []iter.Seq2[K, V]) ([]K, [][]canvas.Point) {
	tmp := make([]map[K]V, len(lines))
	keySeq := make([]iter.Seq[K], len(lines))

	for idx, line := range lines {
		tmp[idx] = maps.Collect(line)
		keySeq[idx] = maps.Keys(tmp[idx])
	}

	keys := slices.Collect(flow.Chain(
		seq.Concat(keySeq...),
		flow.Distinct[K](),
		flow.Sort[K](),
	))

	points := make([][]canvas.Point, len(lines))

	for idx, line := range tmp {
		points[idx] = toLine(keys, line)
	}

	return keys, points
}

func toLine[K cmp.Ordered, V Number](keys []K, line map[K]V) []canvas.Point {
	points := make([]canvas.Point, 0, len(line))

	for key := range seq.Sorted(maps.Keys(line)) {
		points = append(points, canvas.Point{
			X: float64(slices.Index(keys, key)),
			Y: float64(line[key]),
		})
	}

	return points
}

func LinesRect(data [][]canvas.Point) canvas.Rect {
	idxMin, idxMax := 0.0, 0.0
	valMin := math.MaxFloat64
	valMax := math.MaxFloat64 * -1

	for _, items := range data {
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
