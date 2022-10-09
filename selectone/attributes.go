package selectone

import (
	"github.com/Red-Sock/rscli-uikit/common"
	"github.com/Red-Sock/rscli-uikit/label"
	"github.com/nsf/termbox-go"
)

type ColorSurface int

type Attribute func(sb *Box)

func Position(positioner common.Positioner) Attribute {
	return func(sb *Box) {
		sb.pos = positioner
	}
}

func HeaderLabel(header *label.Label) Attribute {
	return func(sb *Box) {
		sb.header = header
	}
}
func Header(header string) Attribute {
	return func(sb *Box) {
		sb.header = label.New(header)
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
func SeparatorSymbol(r rune) Attribute {
	return func(sb *Box) {
		sb.itemSeparator = r
	}
}

func ColorBGCursor(fg, bg termbox.Attribute, type_ ColorSurface) Attribute {
	return func(sb *Box) {
		sb.cursorFG = fg
		sb.cursorBG = bg
	}
}
func ColorBGDefault(fg, bg termbox.Attribute, type_ ColorSurface) Attribute {
	return func(sb *Box) {
		sb.defaultFG = fg
		sb.defaultBG = bg
	}
}
