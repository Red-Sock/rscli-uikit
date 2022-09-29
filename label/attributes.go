package label

import (
	rscliuitkit "github.com/Red-Sock/rscli-uikit"
	"github.com/nsf/termbox-go"
)

type Attribute func(l *Label)

func NextScreen(next func() rscliuitkit.UIElement) Attribute {
	return func(l *Label) {
		l.next = next
	}
}

func X(x int) Attribute {
	return func(l *Label) {
		l.x = x
	}
}

func Y(y int) Attribute {
	return func(l *Label) {
		l.y = y
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
