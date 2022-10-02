package complelabel

import (
	rscliuitkit "github.com/Red-Sock/rscli-uikit"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

type ComplexLabel struct {
	x, y       int
	text       []TextPart
	nextScreen rscliuitkit.UIElement
}

type TextPart struct {
	r      []rune
	fg, bg termbox.Attribute
}

func New() *ComplexLabel {
	return &ComplexLabel{}
}

func (b *ComplexLabel) Render() {
	cursorX, cursorY := b.x, b.y

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
