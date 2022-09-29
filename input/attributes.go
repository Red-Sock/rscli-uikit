package input

import "github.com/nsf/termbox-go"

type Attribute func(box *TextBox)

func FG(fg termbox.Attribute) Attribute {
	return func(box *TextBox) {
		box.fgInput = fg
	}
}

func BG(bg termbox.Attribute) Attribute {
	return func(box *TextBox) {
		box.bgInput = bg
	}
}

// NewAttributeSideSymbols overrides default symbols for rendering TextBox
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

func TextBelow(text string) Attribute {
	return func(box *TextBox) {
		box.textBelowBox = text
	}
}
