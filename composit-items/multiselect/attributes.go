package multiselect

import (
	rscliuitkit "github.com/Red-Sock/rscli-uikit"
	"github.com/Red-Sock/rscli-uikit/basic/label"
	"github.com/Red-Sock/rscli-uikit/utils/common"
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

		sb.header = label.New(header)
	}
}

func HeaderLabel(header *label.Label) Attribute {
	return func(sb *Box) {
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
func SeparatorChecked(r []rune) Attribute {
	return func(sb *Box) {
		sb.itemSeparatorChecked = r
	}
}

func SubmitText(text string) Attribute {
	return func(sb *Box) {
		sb.submitText = text
	}
}

func Position(pos common.Positioner) Attribute {
	return func(sb *Box) {
		sb.pos = pos
	}
}

func PreviousScreen(element rscliuitkit.UIElement) Attribute {
	return func(box *Box) {
		box.PreviousScreen = element
	}
}

func Checked(checked []int) Attribute {
	return func(sb *Box) {
		sb.checkedIdx = checked
	}
}
