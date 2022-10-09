package radio_select

import (
	rscliuitkit "github.com/Red-Sock/rscli-uikit"
	"github.com/Red-Sock/rscli-uikit/utils/common"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

const (
	defaultSeparator = '>'
)

type Box struct {
	header rscliuitkit.Labeler

	items         []string
	itemSeparator rune
	cursorPos     int

	pos common.Positioner

	defaultBG, defaultFG, // default item background and foreground
	cursorBG, cursorFG termbox.Attribute // currently selected with cursor item

	callback func(args string) rscliuitkit.UIElement
}

func New(callback func(args string) rscliuitkit.UIElement, atrs ...Attribute) *Box {
	sb := &Box{
		callback: callback,

		cursorFG: termbox.ColorLightGray,
		cursorBG: termbox.ColorLightCyan,

		defaultFG: termbox.ColorDefault,
		defaultBG: termbox.ColorDefault,

		itemSeparator: defaultSeparator,
		pos:           &common.AbsolutePositioning{},
	}

	for _, a := range atrs {
		a(sb)
	}

	if sb.header != nil {
		sb.header.SetPosition(sb.pos)
	}

	return sb
}

func (s *Box) Render() {
	cursorX, cursorY := s.pos.GetPosition()

	if s.header != nil {
		s.header.Render()

		_, h := s.header.GetSize()
		cursorY += h
	}

	cursorY += 1

	for idx := range s.items {
		fg, bg := s.getColors(idx)
		s.renderItem(s.items[idx], cursorX, cursorY, fg, bg)
		cursorY++
	}
}

func (s *Box) Process(e termbox.Event) rscliuitkit.UIElement {
	switch e.Key {
	case termbox.KeyArrowUp:
		if s.cursorPos > 0 {
			s.cursorPos--
		}
	case termbox.KeyArrowDown:
		if s.cursorPos < len(s.items)-1 {
			s.cursorPos++
		}
	case termbox.KeyEnter:

		if len(s.items) == 0 {
			return s.callback("")
		}
		return s.callback(s.items[s.cursorPos])

	default:
	}
	return s
}

func (s *Box) renderItem(text string, x, y int, fg, bg termbox.Attribute) {
	termbox.SetCell(x, y, s.itemSeparator, fg, bg)
	x++
	for _, r := range text {
		termbox.SetCell(x, y, r, fg, bg)
		x += runewidth.RuneWidth(r)
	}
}

func (s *Box) getColors(idx int) (termbox.Attribute, termbox.Attribute) {
	var fg, bg termbox.Attribute
	switch {
	case idx == s.cursorPos:
		fg, bg = s.cursorFG, s.cursorBG
	default:
		fg, bg = s.defaultFG, s.defaultBG
	}
	return fg, bg
}
