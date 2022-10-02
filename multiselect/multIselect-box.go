package multiselect

import (
	rscliuitkit "github.com/Red-Sock/rscli-uikit"
	"github.com/Red-Sock/rscli-uikit/internal/utils"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

const (
	defaultSeparator = '>'
)

type Box struct {
	header string

	items                    []string
	itemSeparator            []rune
	itemSeparatorUnderCursor []rune
	itemSeparatorChecked     []rune

	submitText string

	checkedIdx []int
	cursorPos  int

	x, y int

	defaultBG, defaultFG, // default item background and foreground
	cursorBG, cursorFG, // currently selected with cursor item
	checkedBG, checkedFG, // marked item, in case of multiselect
	headerBG, headerFG,
	submitBG, submitFG termbox.Attribute

	callback func(args []string) rscliuitkit.UIElement
}

func New(
	callback func(args []string) rscliuitkit.UIElement,
	atrs ...Attribute) *Box {

	sb := &Box{
		callback: callback,

		headerFG: termbox.ColorDefault,
		headerBG: termbox.ColorDarkGray,

		cursorFG: termbox.ColorLightGray,
		cursorBG: termbox.ColorLightCyan,

		defaultFG: termbox.ColorDefault,
		defaultBG: termbox.ColorDefault,

		checkedFG: termbox.ColorGreen,
		checkedBG: termbox.ColorGreen,

		submitFG: termbox.ColorDefault,
		submitBG: termbox.ColorDefault,

		itemSeparator: []rune{defaultSeparator},
	}

	for _, a := range atrs {
		a(sb)
	}

	if sb.submitText == "" {
		sb.submitText = "submit"
	}

	return sb
}

func (s *Box) Render() {
	cursorX, cursorY := s.x, s.y
	for _, r := range s.header {
		termbox.SetCell(cursorX, cursorY, r, s.headerFG, s.headerBG)
		cursorX += runewidth.RuneWidth(r)
	}
	cursorX = s.x
	cursorY = s.y + 1

	for idx := range s.items {
		separator, fg, bg := s.getColors(idx)
		s.renderItem(string(separator)+s.items[idx], cursorX, cursorY, fg, bg)
		cursorY++
	}

	s.renderSubmitButton(cursorX, cursorY)
}

func (s *Box) Process(e termbox.Event) rscliuitkit.UIElement {
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
		if s.cursorPos == len(s.items) {
			args := make([]string, 0, len(s.checkedIdx))
			for _, i := range s.checkedIdx {
				args = append(args, s.items[i])
			}
			return s.callback(args)
		}
		if !utils.Contains(s.checkedIdx, s.cursorPos) {
			s.checkedIdx = append(s.checkedIdx, s.cursorPos)
		} else {
			s.checkedIdx = utils.RemoveItemFromSlice(s.checkedIdx, s.cursorPos)
		}
	default:
	}
	return s
}

func (s *Box) renderItem(text string, x, y int, fg, bg termbox.Attribute) {
	for _, r := range text {
		termbox.SetCell(x, y, r, fg, bg)
		x += runewidth.RuneWidth(r)
	}
}

func (s *Box) renderSubmitButton(cursorX, cursorY int) {
	var fg, bg termbox.Attribute

	if s.cursorPos == len(s.items) {
		fg = s.cursorFG
		bg = s.cursorBG
	} else {
		fg = s.submitFG
		bg = s.submitBG
	}

	for _, r := range s.submitText {
		termbox.SetCell(cursorX, cursorY, r, fg, bg)
		cursorX += runewidth.RuneWidth(r)
	}
}

func (s *Box) getColors(idx int) ([]rune, termbox.Attribute, termbox.Attribute) {
	switch {
	case idx == s.cursorPos:
		return s.itemSeparatorUnderCursor, s.cursorFG, s.cursorBG
	case utils.Contains(s.checkedIdx, idx):
		return s.itemSeparatorChecked, s.checkedFG, s.checkedBG
	default:
		return s.itemSeparator, s.defaultFG, s.defaultBG
	}
}
