package main

import (
	rscliuitkit "github.com/Red-Sock/rscli-uikit"
	"github.com/Red-Sock/rscli-uikit/multiselectbox"
	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	sc, _ := multiselectbox.New(
		nil,
		multiselectbox.HeaderAttribute("choose menu"),
		multiselectbox.ItemsAttribute("hello world", "text-box", "other point", "that other thing"),
	)

	q := make(chan struct{})
	rscliuitkit.NewHandler(sc).Start(q)
}
