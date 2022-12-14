package label

import (
	rscliuitkit "github.com/Red-Sock/rscli-uikit"
	screenDiscovery "github.com/Red-Sock/rscli-uikit/basic/screen-discovery"
	"github.com/Red-Sock/rscli-uikit/internal/utils"
	"github.com/Red-Sock/rscli-uikit/utils/common"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
	"strings"
)

type Label struct {
	screenDiscovery.ScreenDiscovery

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
		h:   1 + strings.Count(text, "\n"),
		pos: &common.AbsolutePositioning{},
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
		currentX := x + int(float32(len(line))*t.wDelta2)
		for _, c := range line {

			termbox.SetCell(currentX, y, c, t.fg, t.bg)
			currentX += runewidth.RuneWidth(c)
		}
		y++
	}
}

func (t *Label) Process(e termbox.Event) rscliuitkit.UIElement {
	if utils.Contains([]termbox.Key{termbox.KeyEsc, termbox.KeyBackspace, termbox.KeyBackspace2}, e.Key) {
		return t.PreviousScreen
	}
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
