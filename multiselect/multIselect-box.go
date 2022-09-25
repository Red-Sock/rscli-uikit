package multiselect

import (
	"errors"
	rscliuitkit "github.com/Red-Sock/rscli-uikit"
	"github.com/Red-Sock/rscli-uikit/internal/common"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

var (
	ErrNoItems = errors.New("no items provide")
)

const (
	defaultSeparator = '>'
)

type MultiSelectBox struct {
	header        string
	items         []string
	itemSeparator rune

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
	atrs ...Attribute) (*MultiSelectBox, error) {

	sb := &MultiSelectBox{
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

		itemSeparator: defaultSeparator,
	}

	for _, a := range atrs {
		a(sb)
	}

	if len(sb.items) == 0 {
		return nil, ErrNoItems
	}

	if sb.submitText == "" {
		sb.submitText = "submit"
	}

	return sb, nil
}

func (s *MultiSelectBox) Render() {
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

	s.renderSubmitButton(cursorX, cursorY)
}

func (s *MultiSelectBox) Process(e termbox.Event) rscliuitkit.UIElement {
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
		if !common.Contains(s.checkedIdx, s.cursorPos) {
			s.checkedIdx = append(s.checkedIdx, s.cursorPos)
		} else {
			s.checkedIdx = common.RemoveItemFromSlice(s.checkedIdx, s.cursorPos)
		}
	default:
	}
	return s
}

func (s *MultiSelectBox) renderItem(text string, x, y int, fg, bg termbox.Attribute) {
	termbox.SetCell(x, y, s.itemSeparator, fg, bg)
	x++
	for _, r := range text {
		termbox.SetCell(x, y, r, fg, bg)
		x += runewidth.RuneWidth(r)
	}
}

func (s *MultiSelectBox) renderSubmitButton(cursorX, cursorY int) {
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

func (s *MultiSelectBox) getColors(idx int) (termbox.Attribute, termbox.Attribute) {
	var fg, bg termbox.Attribute
	switch {
	case idx == s.cursorPos:
		fg, bg = s.cursorFG, s.cursorBG
	case common.Contains(s.checkedIdx, idx):
		fg, bg = s.checkedFG, s.checkedBG
	default:
		fg, bg = s.defaultFG, s.defaultBG
	}
	return fg, bg
}
