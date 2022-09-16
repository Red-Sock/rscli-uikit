package main

import (
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
	"unicode/utf8"
)

func tBPrint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}

func fill(x, y, w, h int, cell termbox.Cell) {
	for ly := 0; ly < h; ly++ {
		for lx := 0; lx < w; lx++ {
			termbox.SetCell(x+lx, y+ly, cell.Ch, cell.Fg, cell.Bg)
		}
	}
}

func runeAdvanceLen(r rune, pos int) int {
	if r == '\t' {
		return tabstop_length - pos%tabstop_length
	}
	return runewidth.RuneWidth(r)
}

func vOffsetCOffset(text []byte, boffset int) (voffset, coffset int) {
	text = text[:boffset]
	for len(text) > 0 {
		r, size := utf8.DecodeRune(text)
		text = text[size:]
		coffset += 1
		voffset += runeAdvanceLen(r, voffset)
	}
	return
}

func byteSliceGrow(s []byte, desiredCap int) []byte {
	if cap(s) < desiredCap {
		ns := make([]byte, len(s), desiredCap)
		copy(ns, s)
		return ns
	}
	return s
}

func byte_slice_remove(text []byte, from, to int) []byte {
	size := to - from
	copy(text[from:], text[to:])
	text = text[:len(text)-size]
	return text
}

func byte_slice_insert(text []byte, offset int, what []byte) []byte {
	n := len(text) + len(what)
	text = byteSliceGrow(text, n)
	text = text[:n]
	copy(text[offset+len(what):], text[offset:])
	copy(text[offset:], what)
	return text
}

const preferred_horizontal_threshold = 5
const tabstop_length = 8

// Draw draws the TextBox in the given location, 'h' is not used at the moment
func (eb *TextBox) Draw(x, y, w, h int) {
	eb.AdjustVOffset(w)

	const coldef = termbox.ColorDefault
	const colred = termbox.ColorRed

	fill(x, y, w, h, termbox.Cell{Ch: ' '})

	t := eb.text
	lx := 0
	tabstop := 0
	for {
		rx := lx - eb.lineVOffset
		if len(t) == 0 {
			break
		}

		if lx == tabstop {
			tabstop += tabstop_length
		}

		if rx >= w {
			termbox.SetCell(x+w-1, y, arrowRight,
				colred, coldef)
			break
		}

		r, size := utf8.DecodeRune(t)
		if r == '\t' {
			for ; lx < tabstop; lx++ {
				rx = lx - eb.lineVOffset
				if rx >= w {
					goto next
				}

				if rx >= 0 {
					termbox.SetCell(x+rx, y, ' ', coldef, coldef)
				}
			}
		} else {
			if rx >= 0 {
				termbox.SetCell(x+rx, y, r, coldef, coldef)
			}
			lx += runewidth.RuneWidth(r)
		}
	next:
		t = t[size:]
	}

	if eb.lineVOffset != 0 {
		termbox.SetCell(x, y, arrowLeft, colred, coldef)
	}
}

// AdjustVOffset Adjusts line visual offset to a proper value depending on width
func (eb *TextBox) AdjustVOffset(width int) {
	ht := preferred_horizontal_threshold
	maxHThreshold := (width - 1) / 2
	if ht > maxHThreshold {
		ht = maxHThreshold
	}

	threshold := width - 1
	if eb.lineVOffset != 0 {
		threshold = width - ht
	}
	if eb.cursorVOffset-eb.lineVOffset >= threshold {
		eb.lineVOffset = eb.cursorVOffset + (ht - width + 1)
	}

	if eb.lineVOffset != 0 && eb.cursorVOffset-eb.lineVOffset < ht {
		eb.lineVOffset = eb.cursorVOffset - ht
		if eb.lineVOffset < 0 {
			eb.lineVOffset = 0
		}
	}
}

func (eb *TextBox) MoveCursorTo(bOffset int) {
	eb.cursorBOffset = bOffset
	eb.cursorVOffset, eb.cursorCOffset = vOffsetCOffset(eb.text, bOffset)
}

func (eb *TextBox) RuneUnderCursor() (rune, int) {
	return utf8.DecodeRune(eb.text[eb.cursorBOffset:])
}

func (eb *TextBox) RuneBeforeCursor() (rune, int) {
	return utf8.DecodeLastRune(eb.text[:eb.cursorBOffset])
}

func (eb *TextBox) MoveCursorOneRuneBackward() {
	if eb.cursorBOffset == 0 {
		return
	}
	_, size := eb.RuneBeforeCursor()
	eb.MoveCursorTo(eb.cursorBOffset - size)
}

func (eb *TextBox) MoveCursorOneRuneForward() {
	if eb.cursorBOffset == len(eb.text) {
		return
	}
	_, size := eb.RuneUnderCursor()
	eb.MoveCursorTo(eb.cursorBOffset + size)
}

func (eb *TextBox) MoveCursorToBeginningOfTheLine() {
	eb.MoveCursorTo(0)
}

func (eb *TextBox) MoveCursorToEndOfTheLine() {
	eb.MoveCursorTo(len(eb.text))
}

func (eb *TextBox) DeleteRuneBackward() {
	if eb.cursorBOffset == 0 {
		return
	}

	eb.MoveCursorOneRuneBackward()
	_, size := eb.RuneUnderCursor()
	eb.text = byte_slice_remove(eb.text, eb.cursorBOffset, eb.cursorBOffset+size)
}

func (eb *TextBox) DeleteRuneForward() {
	if eb.cursorBOffset == len(eb.text) {
		return
	}
	_, size := eb.RuneUnderCursor()
	eb.text = byte_slice_remove(eb.text, eb.cursorBOffset, eb.cursorBOffset+size)
}

func (eb *TextBox) DeleteTheRestOfTheLine() {
	eb.text = eb.text[:eb.cursorBOffset]
}

func (eb *TextBox) InsertRune(r rune) {
	var buf [utf8.UTFMax]byte
	n := utf8.EncodeRune(buf[:], r)
	eb.text = byte_slice_insert(eb.text, eb.cursorBOffset, buf[:n])
	eb.MoveCursorOneRuneForward()
}

// Please, keep in mind that cursor depends on the value of lineVOffset, which
// is being set on Draw() call, so.. call this method after Draw() one.
func (eb *TextBox) CursorX() int {
	return eb.cursorVOffset - eb.lineVOffset
}

var edit_box TextBox

const edit_box_width = 30

func redraw_all() {
	const coldef = termbox.ColorDefault
	termbox.Clear(coldef, coldef)
	w, h := termbox.Size()

	midy := h / 2
	midx := (w - edit_box_width) / 2

	// unicode box drawing chars around the edit box
	if runewidth.EastAsianWidth {
		termbox.SetCell(midx-1, midy, '|', coldef, coldef)
		termbox.SetCell(midx+edit_box_width, midy, '|', coldef, coldef)
		termbox.SetCell(midx-1, midy-1, '+', coldef, coldef)
		termbox.SetCell(midx-1, midy+1, '+', coldef, coldef)
		termbox.SetCell(midx+edit_box_width, midy-1, '+', coldef, coldef)
		termbox.SetCell(midx+edit_box_width, midy+1, '+', coldef, coldef)
		fill(midx, midy-1, edit_box_width, 1, termbox.Cell{Ch: '-'})
		fill(midx, midy+1, edit_box_width, 1, termbox.Cell{Ch: '-'})
	} else {
		termbox.SetCell(midx-1, midy, '│', coldef, coldef)
		termbox.SetCell(midx+edit_box_width, midy, '│', coldef, coldef)
		termbox.SetCell(midx-1, midy-1, '┌', coldef, coldef)
		termbox.SetCell(midx-1, midy+1, '└', coldef, coldef)
		termbox.SetCell(midx+edit_box_width, midy-1, '┐', coldef, coldef)
		termbox.SetCell(midx+edit_box_width, midy+1, '┘', coldef, coldef)
		fill(midx, midy-1, edit_box_width, 1, termbox.Cell{Ch: '─'})
		fill(midx, midy+1, edit_box_width, 1, termbox.Cell{Ch: '─'})
	}

	edit_box.Draw(midx, midy, edit_box_width, 1)
	termbox.SetCursor(midx+edit_box.CursorX(), midy)

	tBPrint(midx+6, midy+3, coldef, coldef, "Press ESC to quit")
	termbox.Flush()
}

var arrowLeft = '←'
var arrowRight = '→'

func init() {
	if runewidth.EastAsianWidth {
		arrowLeft = '<'
		arrowRight = '>'
	}
}
