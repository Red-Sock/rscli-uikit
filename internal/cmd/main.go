package main

import (
	rscliuitkit "github.com/Red-Sock/rscli-uikit"
	"github.com/Red-Sock/rscli-uikit/common"
	"github.com/Red-Sock/rscli-uikit/label"
	"github.com/Red-Sock/rscli-uikit/selectone"
)

func main() {
	f := func(text string) rscliuitkit.UIElement {
		return label.New(text)
	}

	sc := selectone.New(
		f,
		selectone.HeaderLabel(
			label.New(
				"some header",
				label.Position(common.NewRelativePositioning(0.35, 0.35)),
			)),
		selectone.Items("item 1", "item 2", "item 3"),
	)

	//	input.New(
	//	f,
	//	input.TextAbove("choose menu"),
	//	input.TextBelow("hello world"),
	//
	//	input.Expandable(),
	//
	//	input.Position(common.NewRelativePositioning(0.5, 0.5)),
	//)

	q := make(chan struct{})
	rscliuitkit.NewHandler(sc).Start(q)
}
