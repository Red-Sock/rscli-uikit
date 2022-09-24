package textbox

import (
	"github.com/Red-Sock/rscli-uikit"
	"github.com/Red-Sock/rscli-uikit/internal/common"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

type TextBox struct {
	X, Y, H, W int

	// new fields
	rText []rune

	showTextStartCursor int
	editTextCursor      int

	fg, bg termbox.Attribute

	lu, ld, ru, rd, vs, hs rune
}

func NewTextBox(atrs ...Attribute) *TextBox {
	tb := &TextBox{}
	for _, a := range atrs {
		a(tb)
	}

	if tb.fg == 0 {
		NewAttributeFG(termbox.ColorDefault)(tb)
	}

	if tb.bg == 0 {
		NewAttributeBG(termbox.ColorDefault)(tb)
	}

	if tb.lu == 0 {
		NewAttributeSideSymbols('┌', '└', '┐', '┘', '│', '─')(tb)
	}

	return tb
}

func (tb *TextBox) Render() {
	tb.drawBounds()
	tb.drawContent()
	tb.drawCursor()
}

func (tb *TextBox) Process(e termbox.Event) rscliuitkit.Screen {
	switch e.Key {
	case termbox.KeyEsc:
		return nil
	case termbox.KeyArrowLeft:
		if tb.editTextCursor > 0 {
			tb.editTextCursor--
			if tb.editTextCursor < tb.showTextStartCursor {
				tb.showTextStartCursor--
			}
		}
	case termbox.KeyArrowRight:
		if tb.editTextCursor < len(tb.rText) {
			tb.editTextCursor++
			if tb.editTextCursor > tb.showTextStartCursor+tb.W-1 {
				tb.showTextStartCursor++
			}
		}
	case termbox.KeyBackspace, termbox.KeyBackspace2:
		tb.DeleteRune()
	case termbox.KeyDelete, termbox.KeyCtrlD:
		tb.DeleteRune()
	case termbox.KeyTab:
		tb.InsertRune('\t')
	case termbox.KeySpace:
		tb.InsertRune(' ')
	default:
		if e.Ch != 0 {
			tb.InsertRune(e.Ch)
		}
	}
	return tb
}

func (tb *TextBox) InsertRune(r rune) {
	tb.rText = append(tb.rText, r)
	if tb.GetScreenSpace() <= len(tb.rText) {
		tb.showTextStartCursor++
	}
	tb.editTextCursor++
}

func (tb *TextBox) DeleteRune() {
	if len(tb.rText) == 0 || tb.editTextCursor == 0 {
		return
	}
	tb.rText = common.RemoveFromSlice(tb.rText, tb.editTextCursor-1)
	if tb.showTextStartCursor > 0 {
		tb.showTextStartCursor--
	}
	tb.editTextCursor--
}

// GetScreenSpace returns amount of runes that can be displayed at TextBox
func (tb *TextBox) GetScreenSpace() int {
	return tb.W * tb.H
}

func (tb *TextBox) drawBounds() {
	//  top
	termbox.SetCell(tb.X, tb.Y, tb.lu, tb.fg, tb.bg)
	common.FillArea(tb.X+1, tb.Y, tb.W+1, 1, tb.hs, tb.fg, tb.bg)
	termbox.SetCell(tb.X+tb.W+1, tb.Y, tb.ru, tb.fg, tb.bg)

	// sides
	for y := tb.Y + 1; y < tb.Y+tb.H+2; y++ {
		termbox.SetCell(tb.X, y, tb.vs, tb.fg, tb.bg)
		termbox.SetCell(tb.X+tb.W+1, y, tb.vs, tb.fg, tb.bg)
	}

	// bottom
	termbox.SetCell(tb.X, tb.Y+tb.H+1, tb.ld, tb.fg, tb.bg)
	common.FillArea(tb.X+1, tb.Y+tb.H+1, tb.W+1, 1, tb.hs, tb.fg, tb.bg)
	termbox.SetCell(tb.X+tb.W+1, tb.Y+tb.H+1, tb.rd, tb.fg, tb.bg)

}

func (tb *TextBox) drawContent() {
	cursorX, cursorY := tb.X, tb.Y+1
	text := tb.rText[tb.showTextStartCursor:]

	for len(text) > 0 {
		r := text[0]
		rLen := runewidth.RuneWidth(r)

		if cursorX+rLen > tb.X+tb.W {
			cursorX = tb.X + 1
			cursorY++
		} else {
			cursorX += rLen
		}

		if cursorY > tb.Y+tb.H {
			break
		}

		termbox.SetCell(cursorX, cursorY, r, tb.fg, tb.bg)

		text = text[1:]
	}

	filledCells := len(tb.rText)
	for filledCells < tb.GetScreenSpace() {
		if cursorX+1 > tb.X+tb.W {
			cursorX = tb.X + 1
			cursorY++
		} else {
			cursorX += 1
		}

		if cursorY > tb.Y+tb.H {
			break
		}
		termbox.SetCell(cursorX, cursorY, ' ', tb.fg, tb.bg)
	}
}

func (tb *TextBox) drawCursor() {
	termbox.SetCursor(
		tb.X+1+(tb.editTextCursor-tb.showTextStartCursor)%tb.W,
		tb.Y+1+(tb.editTextCursor-tb.showTextStartCursor)/tb.W,
	)
}
