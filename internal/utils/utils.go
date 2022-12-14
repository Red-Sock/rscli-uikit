package utils

import "github.com/nsf/termbox-go"

func FillArea(x, y, w, h int, char rune, fg, bg termbox.Attribute) {
	for ly := 0; ly < h; ly++ {
		for lx := 0; lx < w; lx++ {
			termbox.SetCell(x+lx, y+ly, char, fg, bg)
		}
	}
}

func RemoveFromSlice[T comparable](slice []T, idx int) []T {
	if idx >= len(slice) {
		return slice
	}

	out := slice[:idx]
	if idx < len(slice)-1 {
		out = append(out, slice[idx+1:]...)
	}
	return out
}

func RemoveItemFromSlice[T comparable](items []T, item T) []T {
	for idx, is := range items {
		if is == item {
			return RemoveFromSlice(items, idx)
		}
	}
	return items
}

func Contains[T comparable](items []T, item T) bool {
	for _, si := range items {
		if si == item {
			return true
		}
	}

	return false
}

func AddOrMax(startValue, addValue, maxValue int) int {
	startValue += addValue
	if startValue < maxValue {
		return startValue
	}

	return maxValue
}

func SubtractOrMin(startValue, subValue, minValue int) int {
	startValue -= subValue
	if startValue > minValue {
		return startValue
	}

	return minValue
}

func MaxInt(items ...int) int {
	var max int

	if len(items) == 0 {
		return max
	}

	for _, item := range items {
		if item > max {
			max = item
		}
	}

	return max
}
