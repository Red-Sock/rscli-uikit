package label

import (
	rscliuitkit "github.com/Red-Sock/rscli-uikit"
	"github.com/Red-Sock/rscli-uikit/utils/common"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

type Label struct {
	text string

	pos    common.Positioner
	fg, bg termbox.Attribute

	anchorType AnchorType

	w, h int

	next func() rscliuitkit.UIElement
}

func New(text string, attrs ...Attribute) *Label {
	l := &Label{
		pos: &common.AbsolutePositioning{},
	}

	for _, a := range attrs {
		a(l)
	}
	l.text = text
	// todo
	switch l.anchorType {
	case Centered:
		textLen := 0
		for _, r := range l.text {
			textLen += runewidth.RuneWidth(r)
		}

	case Right:

	}
	return l
}

func (t *Label) Render() {
	x, y := t.pos.GetPosition()
	for _, c := range t.text {
		termbox.SetCell(x, y, c, t.fg, t.bg)
		x += runewidth.RuneWidth(c)
	}
}

func (t *Label) Process(_ termbox.Event) rscliuitkit.UIElement {
	if t.next == nil {
		return nil
	}
	return t.next()
}

func (t *Label) GetSize() (w, h int) {
	return t.w, t.h
}

func (t *Label) SetPosition(p common.Positioner) {
	t.pos = p
}
