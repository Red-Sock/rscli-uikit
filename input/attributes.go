package input

import "github.com/nsf/termbox-go"

type Attribute func(box *TextBox)

func X(x int) Attribute {
	return func(box *TextBox) {
		box.X = x
	}
}
func Y(y int) Attribute {
	return func(box *TextBox) {
		box.Y = y
	}
}
func Width(w int) Attribute {
	return func(box *TextBox) {
		box.W = w
	}
}
func Height(h int) Attribute {
	return func(box *TextBox) {
		box.H = h
	}
}

func ContentFG(fg termbox.Attribute) Attribute {
	return func(box *TextBox) {
		box.fgInput = fg
	}
}
func ContentBG(bg termbox.Attribute) Attribute {
	return func(box *TextBox) {
		box.bgInput = bg
	}
}

// SideSymbols overrides default symbols for rendering TextBox
// lu - left upper corner, ld - left down
// ru - right upper, rd - right down
// vs - vertical side
// hs - horizontal side
func SideSymbols(lu, ld, ru, rd, vs, hs rune) Attribute {
	return func(tb *TextBox) {
		tb.lu = lu
		tb.ld = ld
		tb.ru = ru
		tb.rd = rd
		tb.vs = vs
		tb.hs = hs
	}
}

func TextAbove(text string) Attribute {
	return func(box *TextBox) {
		box.textAboveBox = text
		box.Y++
	}
}
func TextAboveColor(fg, bg termbox.Attribute) Attribute {
	return func(box *TextBox) {
		box.textAboveFg = fg
		box.textAboveBg = bg
	}
}

func TextBelow(text string) Attribute {
	return func(box *TextBox) {
		box.textBelowBox = text
	}
}
func TextBelowColor(fg, bg termbox.Attribute) Attribute {
	return func(box *TextBox) {
		box.textBelowFg = fg
		box.textBelowBg = bg
	}
}
