package chart

import "slices"

func ShuffleSlice[T any](slice []T, step int) []T {
	length := len(slice)
	if length <= 1 {
		return slice
	}

	step %= length
	if step <= 1 {
		return slice
	}

	result := make([]T, length)
	group := slices.Collect(slices.Chunk(slice, step))
	idxResult := 0

	for idx := range length {
		for _, items := range group {
			if idx >= len(items) {
				continue
			}

			result[idxResult] = items[idx]
			idxResult++
		}
	}

	return result
}
