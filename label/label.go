package label

import (
	rscliuitkit "github.com/Red-Sock/rscli-uikit"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

type Label struct {
	text string

	x, y   int
	fg, bg termbox.Attribute

	next func() rscliuitkit.UIElement
}

func New(text string, attrs ...Attribute) rscliuitkit.UIElement {
	l := &Label{}

	for _, a := range attrs {
		a(l)
	}
	l.text = text
	return l
}

func (t *Label) Render() {
	x := t.x
	for _, c := range t.text {
		termbox.SetCell(x, t.y, c, t.fg, t.bg)
		x += runewidth.RuneWidth(c)
	}
}

func (t *Label) Process(e termbox.Event) rscliuitkit.UIElement {
	if t.next == nil {
		return nil
	}
	return t.next()
}
