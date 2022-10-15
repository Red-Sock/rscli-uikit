package label

import (
	rscliuitkit "github.com/Red-Sock/rscli-uikit"
	"github.com/Red-Sock/rscli-uikit/utils/common"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
	"strings"
)

type Label struct {
	lines [][]rune

	pos    common.Positioner
	fg, bg termbox.Attribute

	anchorType AnchorType

	w, h int

	wDelta int

	wDelta2 float32

	next func() rscliuitkit.UIElement
}

func New(text string, attrs ...Attribute) *Label {
	l := &Label{
		h: 1 + strings.Count(text, "\n"),
	}

	for _, a := range attrs {
		a(l)
	}

	maxWidth, currentLength := 0, 0
	currentLine := make([]rune, 0, len(text)/2)

	for _, r := range text {
		if r == '\n' {
			if maxWidth < currentLength {
				maxWidth = currentLength
			}
			currentLength = 0

			l.lines = append(l.lines, currentLine)
			currentLine = make([]rune, len(text)/2)
			continue
		}
		currentLength += runewidth.RuneWidth(r)
		currentLine = append(currentLine, r)
	}

	l.lines = append(l.lines, currentLine)

	l.w = maxWidth

	return l
}

func (t *Label) Render() {
	x, y := t.pos.GetPosition()

	for _, line := range t.lines {
		x += int(float32(len(line)) * t.wDelta2)
		for _, c := range line {

			termbox.SetCell(x, y, c, t.fg, t.bg)
			x += runewidth.RuneWidth(c)
		}
		y++
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
