package multiselect

import (
	"github.com/nsf/termbox-go"
	"strings"
)

type ColorSurface int

const (
	ColorSurfaceDefault ColorSurface = iota
	ColorSurfaceUnderCursor
	ColorSurfaceChecked
	ColorSurfaceHeader
	ColorSurfaceSubmit
)

type Attribute func(sb *MultiSelectBox)

func HeaderAttribute(header string) Attribute {
	return func(sb *MultiSelectBox) {
		if !strings.HasSuffix(header, ":") {
			header += ":"
		}

		sb.header = header
	}
}

func ItemsAttribute(items ...string) Attribute {
	return func(sb *MultiSelectBox) {
		sb.items = make([]string, len(items))
		for idx := range items {
			sb.items[idx] = items[idx]
		}

	}
}

func CoordinatesAttribute(x, y int) Attribute {
	return func(sb *MultiSelectBox) {
		sb.x = x
		sb.y = y
	}
}

func ColorBGAttribute(fg, bg termbox.Attribute, type_ ColorSurface) Attribute {
	return func(sb *MultiSelectBox) {
		switch type_ {
		case ColorSurfaceDefault:
			sb.defaultFG = fg
			sb.defaultBG = bg
		case ColorSurfaceUnderCursor:
			sb.cursorFG = fg
			sb.cursorBG = bg
		case ColorSurfaceChecked:
			sb.checkedFG = fg
			sb.checkedBG = bg
		case ColorSurfaceHeader:
			sb.headerFG = fg
			sb.headerBG = bg
		case ColorSurfaceSubmit:
			sb.submitFG = fg
			sb.submitBG = bg
		}
	}
}

func SeparatorSymbolAttribute(r rune) Attribute {
	return func(sb *MultiSelectBox) {
		sb.itemSeparator = r
	}
}

func SubmitTextAttribute(text string) Attribute {
	return func(sb *MultiSelectBox) {
		sb.submitText = text
	}
}
