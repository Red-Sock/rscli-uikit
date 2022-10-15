package main

import (
	rscliuitkit "github.com/Red-Sock/rscli-uikit"
	"github.com/Red-Sock/rscli-uikit/basic/label"
	"github.com/Red-Sock/rscli-uikit/composit-items/radioselect"
	"github.com/Red-Sock/rscli-uikit/utils/common"
)

func main() {
	f := func(text string) rscliuitkit.UIElement {
		return label.New(text)
	}

	sc := radioselect.New(
		f,
		radioselect.HeaderLabel(
			label.New(
				"some header",
				label.Anchor(label.Right),
			)),
		radioselect.Items("item 1", "item 2", "item 3"),
		radioselect.Position(common.NewRelativePositioning(0.35, 0.35)),
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
