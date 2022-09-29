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

type Attribute func(sb *Box)

func Header(header string) Attribute {
	return func(sb *Box) {
		if !strings.HasSuffix(header, ":") {
			header += ":"
		}

		sb.header = header
	}
}

func Items(items ...string) Attribute {
	return func(sb *Box) {
		sb.items = make([]string, len(items))
		for idx := range items {
			sb.items[idx] = items[idx]
		}

	}
}

func X(x int) Attribute {
	return func(sb *Box) {
		sb.x = x
	}
}

func Y(y int) Attribute {
	return func(sb *Box) {
		sb.y = y
	}
}

func ColorBG(fg, bg termbox.Attribute, type_ ColorSurface) Attribute {
	return func(sb *Box) {
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

func SeparatorSymbol(r []rune) Attribute {
	return func(sb *Box) {
		sb.itemSeparator = r
	}
}
func SeparatorUnderCursor(r []rune) Attribute {
	return func(sb *Box) {
		sb.itemSeparatorUnderCursor = r
	}
}
func SeparatorCheckedCursor(r []rune) Attribute {
	return func(sb *Box) {
		sb.itemSeparatorUnderCursor = r
	}
}

func SubmitText(text string) Attribute {
	return func(sb *Box) {
		sb.submitText = text
	}
}
