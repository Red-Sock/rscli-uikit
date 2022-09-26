package selectone

import (
	"errors"
	rscliuitkit "github.com/Red-Sock/rscli-uikit"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

var (
	ErrNoItems = errors.New("no items provide")
)

const (
	defaultSeparator = '>'
)

type SelectBox struct {
	header        string
	items         []string
	itemSeparator rune

	cursorPos int

	x, y int

	defaultBG, defaultFG, // default item background and foreground
	cursorBG, cursorFG, // currently selected with cursor item
	headerBG, headerFG termbox.Attribute

	callback func(args string) rscliuitkit.UIElement
}

func New(
	callback func(args string) rscliuitkit.UIElement,
	atrs ...Attribute) (*SelectBox, error) {

	sb := &SelectBox{
		callback: callback,

		headerFG: termbox.ColorDefault,
		headerBG: termbox.ColorDarkGray,

		cursorFG: termbox.ColorLightGray,
		cursorBG: termbox.ColorLightCyan,

		defaultFG: termbox.ColorDefault,
		defaultBG: termbox.ColorDefault,

		itemSeparator: defaultSeparator,
	}

	for _, a := range atrs {
		a(sb)
	}

	if len(sb.items) == 0 {
		return nil, ErrNoItems
	}

	return sb, nil
}

func (s *SelectBox) Render() {
	cursorX, cursorY := s.x, s.y
	for _, r := range s.header {
		termbox.SetCell(cursorX, cursorY, r, s.headerFG, s.headerBG)
		cursorX += runewidth.RuneWidth(r)
	}
	cursorX = s.x
	cursorY = s.y + 1

	for idx := range s.items {
		fg, bg := s.getColors(idx)
		s.renderItem(s.items[idx], cursorX, cursorY, fg, bg)
		cursorY++
	}
}

func (s *SelectBox) Process(e termbox.Event) rscliuitkit.UIElement {
	switch e.Key {
	case termbox.KeyArrowUp:
		if s.cursorPos > 0 {
			s.cursorPos--
		}
	case termbox.KeyArrowDown:
		if s.cursorPos < len(s.items) {
			s.cursorPos++
		}
	case termbox.KeyEnter:

		return s.callback(s.items[s.cursorPos])

	default:
	}
	return s
}

func (s *SelectBox) renderItem(text string, x, y int, fg, bg termbox.Attribute) {
	termbox.SetCell(x, y, s.itemSeparator, fg, bg)
	x++
	for _, r := range text {
		termbox.SetCell(x, y, r, fg, bg)
		x += runewidth.RuneWidth(r)
	}
}

func (s *SelectBox) getColors(idx int) (termbox.Attribute, termbox.Attribute) {
	var fg, bg termbox.Attribute
	switch {
	case idx == s.cursorPos:
		fg, bg = s.cursorFG, s.cursorBG
	default:
		fg, bg = s.defaultFG, s.defaultBG
	}
	return fg, bg
}
