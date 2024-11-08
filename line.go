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

func Lines[K cmp.Ordered, V Number](lines map[string]map[K]V) ([]K, map[string][]canvas.Point) {
	keySeq := []iter.Seq[K]{}
	for _, line := range lines {
		keySeq = append(keySeq, maps.Keys(line))
	}

	keys := slices.Collect(flow.Chain(
		seq.Concat(keySeq...),
		flow.Distinct[K](),
		flow.Sort[K](),
	))

	points := map[string][]canvas.Point{}

	for name, line := range lines {
		points[name] = toLine(keys, line)
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

func LinesRect(data iter.Seq[[]canvas.Point]) canvas.Rect {
	idxMin, idxMax := 0.0, 0.0
	valMin := math.MaxFloat64
	valMax := math.MaxFloat64 * -1

	for items := range data {
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
