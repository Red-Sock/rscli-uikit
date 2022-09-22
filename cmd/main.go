package main

import (
	rscliuitkit "github.com/Red-Sock/rscli-uikit"
	"github.com/Red-Sock/rscli-uikit/textbox"
	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	sc := textbox.NewTextBox(
		textbox.NewAttributeBG(termbox.ColorLightMagenta),
		textbox.NewAttributeFG(termbox.ColorRed),
	)

	sc.Y = 2
	sc.X = 3

	sc.W, sc.H = 20, 1
	q := make(chan struct{})
	rscliuitkit.NewHandler(sc).Start(q)
}
