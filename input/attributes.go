package input

import (
	"github.com/Red-Sock/rscli-uikit/internal/common"
	"github.com/nsf/termbox-go"
)

type Attribute func(box *TextBox)

func Position(pos common.Positioner) Attribute {
	return func(box *TextBox) {
		box.pos = pos
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

func Expandable() Attribute {
	return func(box *TextBox) {
		box.isExpandable = true
		box.maxW = 0
	}
}
func ExpandableMax(value int) Attribute {
	return func(box *TextBox) {
		box.isExpandable = true
		box.maxW = value
	}
}
func ExpandableMin(value int) Attribute {
	return func(box *TextBox) {
		box.isExpandable = true
		box.minW = value
	}
}
func ExpandableStep(value int) Attribute {
	return func(box *TextBox) {
		box.expandingStep = value
	}
}
