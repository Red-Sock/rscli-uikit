# CLI UIKIT for golang

[Golang termbox implementation](http://github.com/nsf/termbox-go) based uikit for building diverse cli tools

## Package is in development

## Released features

- Select one screen
- Multiselect screen
- Input box
- Simple label

## Paradigm

For now - callback are the main and only option to shift between screens.
After some condition is completed a function gets called.
Function must return a new screen or nothing (nil) 
if after completing action program has to be stopped   

## Example

```go
package main

import (
	rscliuitkit "github.com/Red-Sock/rscli-uikit"
	"github.com/Red-Sock/rscli-uikit/common"
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

		input.Position(common.NewRelativePositioning(0.5, 0.5)),
	)

	q := make(chan struct{})
	rscliuitkit.NewHandler(sc).Start(q)
}
```

### Common features
EVERYTHING is customisable

### Select one screen
Allows to create menu where user has to choose one of given items

### Multiselect screen
Allows to create menu where user has to choose one or more of given items and submit his choice

### Label 
Prints text on the screen

### Input box
Simple input field with boundaries. 
Can be customized to support multirow input.