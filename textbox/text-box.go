package textbox

import (
	"github.com/Red-Sock/rscli-uikit"
	"github.com/Red-Sock/rscli-uikit/common"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

type TextBoxAttribute interface {
	Apply(box *TextBox)
}

type TextBox struct {
	// new fields
	rText      []rune
	X, Y, H, W int

	showTextStartCursor int
	editTextCursor      int

	// TODO
	IsFullscreen bool
}

func NewTextBox() *TextBox {
	return &TextBox{}
}

func (tb *TextBox) Render() {
	defaultColor := termbox.ColorDefault
	// filling background of text box

	// filling top
	termbox.SetCell(tb.X, tb.Y, '┌', defaultColor, defaultColor)
	common.FillArea(tb.X+1, tb.Y, tb.W+1, 1, '─', defaultColor, defaultColor)
	termbox.SetCell(tb.X+tb.W+1, tb.Y, '┐', defaultColor, defaultColor)

	// filling left and right bounds
	for y := tb.Y + 1; y < tb.Y+tb.H+2; y++ {
		termbox.SetCell(tb.X, y, '│', defaultColor, defaultColor)        // tb.W+2)*y].char = '│'
		termbox.SetCell(tb.X+tb.W+1, y, '│', defaultColor, defaultColor) // pixels[(tb.W+2)*y+tb.W].char = '│'
	}

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

		termbox.SetCell(cursorX, cursorY, r, defaultColor, defaultColor)

		text = text[1:]
	}

	// draw bottom
	termbox.SetCell(tb.X, tb.Y+tb.H+1, '└', defaultColor, defaultColor)
	common.FillArea(tb.X+1, tb.Y+tb.H+1, tb.W+1, 1, '─', defaultColor, defaultColor)
	termbox.SetCell(tb.X+tb.W+1, tb.Y+tb.H+1, '┘', defaultColor, defaultColor)

	cursorX = tb.X + 1 + (tb.editTextCursor-tb.showTextStartCursor)%tb.W
	cursorY = tb.Y + 1 + (tb.editTextCursor-tb.showTextStartCursor)/tb.W

	termbox.SetCursor(cursorX, cursorY)

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
	//case termbox.KeyCtrlK:
	//	tb.DeleteTheRestOfTheLine()
	//case termbox.KeyHome, termbox.KeyCtrlA:
	//	tb.MoveCursorToBeginningOfTheLine()
	//case termbox.KeyEnd, termbox.KeyCtrlE:
	//	tb.MoveCursorToEndOfTheLine()
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
