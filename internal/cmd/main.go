package main

import (
	rscliuitkit "github.com/Red-Sock/rscli-uikit"
	"github.com/Red-Sock/rscli-uikit/input"
	"github.com/Red-Sock/rscli-uikit/label"
)

func main() {
	f := func(text string) rscliuitkit.UIElement {
		return label.New(text)
	}

	sc := input.New(
		f,
		input.TextAbove("choose menu"),
		input.TextBelow("hello world"),

		input.Expandable(),
	)

	q := make(chan struct{})
	rscliuitkit.NewHandler(sc).Start(q)
}
