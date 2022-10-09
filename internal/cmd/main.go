package main

import (
	rscliuitkit "github.com/Red-Sock/rscli-uikit"
	label2 "github.com/Red-Sock/rscli-uikit/basic/label"
	selectone2 "github.com/Red-Sock/rscli-uikit/composit-items/radio-select"
	"github.com/Red-Sock/rscli-uikit/utils/common"
)

func main() {
	f := func(text string) rscliuitkit.UIElement {
		return label2.New(text)
	}

	sc := selectone2.New(
		f,
		selectone2.HeaderLabel(
			label2.New(
				"some header",
				label2.Position(common.NewRelativePositioning(0.35, 0.35)),
			)),
		selectone2.Items("item 1", "item 2", "item 3"),
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
