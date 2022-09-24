package input

import "github.com/nsf/termbox-go"

type Attribute func(box *TextBox)

func NewAttributeFG(fg termbox.Attribute) Attribute {
	return func(box *TextBox) {
		box.fg = fg
	}
}

func NewAttributeBG(bg termbox.Attribute) Attribute {
	return func(box *TextBox) {
		box.bg = bg
	}
}

// NewAttributeSideSymbols overrides default symbols for rendering TextBox
// lu - left upper corner, ld - left down
// ru - right upper, rd - right down
// vs - vertical side
// hs - horizontal side
func NewAttributeSideSymbols(lu, ld, ru, rd, vs, hs rune) Attribute {
	return func(tb *TextBox) {
		tb.lu = lu
		tb.ld = ld
		tb.ru = ru
		tb.rd = rd
		tb.vs = vs
		tb.hs = hs
	}
}
