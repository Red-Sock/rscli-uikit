package rscliuitkit

//
//func tBPrint(x, y int, fg, bg termbox.Attribute, msg string) {
//	for _, c := range msg {
//		termbox.SetCell(x, y, c, fg, bg)
//		x += runewidth.RuneWidth(c)
//	}
//}
//
//func fillOld(x, y, w, h int, cell termbox.Cell) {
//	for ly := 0; ly < h; ly++ {
//		for lx := 0; lx < w; lx++ {
//			termbox.SetCell(x+lx, y+ly, cell.Ch, cell.Fg, cell.Bg)
//		}
//	}
//}
//
//func runeAdvanceLen(r rune, pos int) int {
//	if r == '\t' {
//		return textbox.tabStopLengthTB - pos%textbox.tabStopLengthTB
//	}
//	return runewidth.RuneWidth(r)
//}
//
//func vOffsetCOffset(text []byte, boffset int) (voffset, coffset int) {
//	text = text[:boffset]
//	for len(text) > 0 {
//		r, size := utf8.DecodeRune(text)
//		text = text[size:]
//		coffset += 1
//		voffset += runeAdvanceLen(r, voffset)
//	}
//	return
//}
//
//func byteSliceGrow(s []byte, desiredCap int) []byte {
//	if cap(s) < desiredCap {
//		ns := make([]byte, len(s), desiredCap)
//		copy(ns, s)
//		return ns
//	}
//	return s
//}
//
//func byte_slice_remove(text []byte, from, to int) []byte {
//	size := to - from
//	copy(text[from:], text[to:])
//	text = text[:len(text)-size]
//	return text
//}
//
//func byte_slice_insert(text []byte, offset int, what []byte) []byte {
//	n := len(text) + len(what)
//	text = byteSliceGrow(text, n)
//	text = text[:n]
//	copy(text[offset+len(what):], text[offset:])
//	copy(text[offset:], what)
//	return text
//}
//
//func (tb *textbox.TextBox) MoveCursorTo(bOffset int) {
//	tb.cursorBOffset = bOffset
//	tb.cursorVOffset, tb.cursorCOffset = vOffsetCOffset(tb.text, bOffset)
//}
//
//func (tb *textbox.TextBox) RuneUnderCursor() (rune, int) {
//	return utf8.DecodeRune(tb.text[tb.cursorBOffset:])
//}
//
//func (tb *textbox.TextBox) RuneBeforeCursor() (rune, int) {
//	return utf8.DecodeLastRune(tb.text[:tb.cursorBOffset])
//}
//
//func (tb *textbox.TextBox) MoveCursorOneRuneBackward() {
//	if tb.cursorBOffset == 0 {
//		return
//	}
//	_, size := tb.RuneBeforeCursor()
//	tb.MoveCursorTo(tb.cursorBOffset - size)
//}
//
//func (tb *textbox.TextBox) MoveCursorForward() {
//	if tb.cursorBOffset == len(tb.text) {
//		return
//	}
//	_, size := tb.RuneUnderCursor()
//	tb.MoveCursorTo(tb.cursorBOffset + size)
//}
//
//func (tb *textbox.TextBox) MoveCursorToBeginningOfTheLine() {
//	tb.MoveCursorTo(0)
//}
//
//func (tb *textbox.TextBox) MoveCursorToEndOfTheLine() {
//	tb.MoveCursorTo(len(tb.text))
//}
//
//func (tb *textbox.TextBox) DeleteRuneBackward() {
//	if tb.cursorBOffset == 0 {
//		return
//	}
//
//	tb.MoveCursorOneRuneBackward()
//	_, size := tb.RuneUnderCursor()
//	tb.text = byte_slice_remove(tb.text, tb.cursorBOffset, tb.cursorBOffset+size)
//}
//
//func (tb *textbox.TextBox) DeleteRuneForward() {
//	if tb.cursorBOffset == len(tb.text) {
//		return
//	}
//	_, size := tb.RuneUnderCursor()
//	tb.text = byte_slice_remove(tb.text, tb.cursorBOffset, tb.cursorBOffset+size)
//}
//
//func (tb *textbox.TextBox) DeleteTheRestOfTheLine() {
//	tb.text = tb.text[:tb.cursorBOffset]
//}
//
//// Please, keep in mind that cursor depends on the value of lineVOffset, which
//// is being set on Render() call, so.. call this method after Render() one.
//func (tb *textbox.TextBox) CursorX() int {
//	return tb.cursorVOffset - tb.lineVOffset
//}
//
//var edit_box textbox.TextBox
//
//const editBoxDefaultWidth = 30
//
//func redraw_all() {
//	const coldef = termbox.ColorDefault
//	termbox.Clear(coldef, coldef)
//	w, h := termbox.Size()
//
//	midy := h / 2
//	midx := (w - editBoxDefaultWidth) / 2
//
//	// unicode box drawing chars around the edit box
//	if runewidth.EastAsianWidth {
//		termbox.SetCell(midx-1, midy, '|', coldef, coldef)
//		termbox.SetCell(midx+editBoxDefaultWidth, midy, '|', coldef, coldef)
//		termbox.SetCell(midx-1, midy-1, '+', coldef, coldef)
//		termbox.SetCell(midx-1, midy+1, '+', coldef, coldef)
//		termbox.SetCell(midx+editBoxDefaultWidth, midy-1, '+', coldef, coldef)
//		termbox.SetCell(midx+editBoxDefaultWidth, midy+1, '+', coldef, coldef)
//		fillOld(midx, midy-1, editBoxDefaultWidth, 1, termbox.Cell{Ch: '-'})
//		fillOld(midx, midy+1, editBoxDefaultWidth, 1, termbox.Cell{Ch: '-'})
//	} else {
//		//termbox.SetCell(midx-1, midy, '│', coldef, coldef)
//		termbox.SetCell(midx+editBoxDefaultWidth, midy, '│', coldef, coldef)
//		termbox.SetCell(midx-1, midy-1, '┌', coldef, coldef)
//		termbox.SetCell(midx-1, midy+1, '└', coldef, coldef)
//		termbox.SetCell(midx+editBoxDefaultWidth, midy-1, '┐', coldef, coldef)
//		termbox.SetCell(midx+editBoxDefaultWidth, midy+1, '┘', coldef, coldef)
//		fillOld(midx, midy-1, editBoxDefaultWidth, 1, termbox.Cell{Ch: '─'})
//		fillOld(midx, midy+1, editBoxDefaultWidth, 1, termbox.Cell{Ch: '─'})
//	}
//
//	//edit_box.Render(midx, midy, editBoxDefaultWidth, 1)
//	termbox.SetCursor(midx+edit_box.CursorX(), midy)
//
//	tBPrint(midx+6, midy+3, coldef, coldef, "Press ESC to quit")
//	termbox.Flush()
//}
//
//var arrowLeft = '←'
//var arrowRight = '→'
