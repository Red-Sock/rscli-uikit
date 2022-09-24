package multiselectbox

import (
	"github.com/nsf/termbox-go"
	"strings"
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

func ColorBGAttribute(bg termbox.Attribute) Attribute {
	return func(sb *MultiSelectBox) {
		sb.defaultBG = bg
	}
}

func PointerAttribute(r rune) Attribute {
	return func(sb *MultiSelectBox) {
		sb.itemSeparator = r
	}
}
