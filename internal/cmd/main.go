package main

import (
	rscliuitkit "github.com/Red-Sock/rscli-uikit"
	"github.com/Red-Sock/rscli-uikit/selectone"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	f := func(text string) rscliuitkit.Screen {
		return &testScreen{text: text}
	}

	sc, _ := selectone.New(
		f,
		selectone.HeaderAttribute("choose menu"),
		selectone.ItemsAttribute("hello world", "text-box", "other point", "that other thing"),
	)

	q := make(chan struct{})
	rscliuitkit.NewHandler(sc).Start(q)
}

type testScreen struct {
	text string

	x, y   int
	fg, bg termbox.Attribute
}

func (t *testScreen) Render() {
	x := t.x
	for _, c := range t.text {
		termbox.SetCell(x, t.y, c, t.fg, t.bg)
		x += runewidth.RuneWidth(c)
	}
}

func (t *testScreen) Process(e termbox.Event) rscliuitkit.Screen {
	return nil
}
