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
