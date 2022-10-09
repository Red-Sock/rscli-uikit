package label

import (
	rscliuitkit "github.com/Red-Sock/rscli-uikit"
	"github.com/Red-Sock/rscli-uikit/common"
	"github.com/nsf/termbox-go"
)

type Attribute func(l *Label)

func NextScreen(next func() rscliuitkit.UIElement) Attribute {
	return func(l *Label) {
		l.next = next
	}
}

func Position(pos common.Positioner) Attribute {
	return func(box *Label) {
		box.pos = pos
	}
}

func Fg(fg termbox.Attribute) Attribute {
	return func(l *Label) {
		l.fg = fg
	}
}

func Bg(bg termbox.Attribute) Attribute {
	return func(l *Label) {
		l.bg = bg
	}
}

type AnchorType uint8

const (
	Left AnchorType = iota
	Centered
	Right
)

// Anchor sets a point where text will start
// Left - text starts at given coordinates
// Centered - given coordinates are the center of the text
// Right - text ends at given coordinates
func Anchor(at AnchorType) Attribute {
	return func(box *Label) {
		box.anchorType = at
	}
}
