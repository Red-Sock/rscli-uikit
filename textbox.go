package main

type TextBox struct {
	text          []byte
	lineVOffset   int
	cursorBOffset int // cursor offset in bytes
	cursorVOffset int // visual cursor offset in termbox cells
	cursorCOffset int // cursor offset in unicode code points

	// TODO
	isFullscreen bool
}

func NewTextBox() {

}
