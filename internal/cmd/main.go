package main

import (
	rscliuitkit "github.com/Red-Sock/rscli-uikit"
	"github.com/Red-Sock/rscli-uikit/label"
	"github.com/Red-Sock/rscli-uikit/selectone"
)

func main() {
	f := func(text string) rscliuitkit.UIElement {
		return label.New(text)
	}

	sc, _ := selectone.New(
		f,
		selectone.HeaderAttribute("choose menu"),
		selectone.ItemsAttribute("hello world", "text-box", "other point", "that other thing"),
	)

	q := make(chan struct{})
	rscliuitkit.NewHandler(sc).Start(q)
}
