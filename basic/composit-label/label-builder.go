package composit_label

import (
	rscliuitkit "github.com/Red-Sock/rscli-uikit"
	"github.com/Red-Sock/rscli-uikit/utils/common"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

type ComplexLabel struct {
	pos        common.Positioner
	text       []TextPart
	nextScreen rscliuitkit.UIElement
}

type TextPart struct {
	r      []rune
	fg, bg termbox.Attribute
}

func New(attrs ...Attribute) *ComplexLabel {
	cl := &ComplexLabel{
		pos: &common.AbsolutePositioning{},
	}

	for _, a := range attrs {
		a(cl)
	}

	return cl
}

func (b *ComplexLabel) Render() {
	cursorX, cursorY := b.pos.GetPosition()

	for _, item := range b.text {
		for _, r := range item.r {
			termbox.SetCell(cursorX, cursorY, r, item.fg, item.bg)
			cursorX += runewidth.RuneWidth(r)
		}
	}
}

func (b *ComplexLabel) Process(_ termbox.Event) rscliuitkit.UIElement {
	return b.nextScreen
}

func (b *ComplexLabel) AddText(t TextPart) {
	b.text = append(b.text, t)
}
