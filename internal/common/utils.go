package common

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
