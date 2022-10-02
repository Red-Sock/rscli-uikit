package input

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

	textAboveBox        string
	textBelowBox        string
	showTextStartCursor int
	editTextCursor      int

	fgInput, bgInput,
	textAboveFg, textAboveBg,
	textBelowFg, textBelowBg termbox.Attribute

	lu, ld, ru, rd, vs, hs rune

	isExpandable  bool
	maxW, minW    int
	expandingStep int

	callback func(s string) rscliuitkit.UIElement
}

func New(callback func(s string) rscliuitkit.UIElement, atrs ...Attribute) *TextBox {
	tb := &TextBox{
		callback: callback,
		fgInput:  termbox.ColorDefault,
		bgInput:  termbox.ColorDefault,
		lu:       '┌',
		ld:       '└',
		ru:       '┐',
		rd:       '┘',
		vs:       '│',
		hs:       '─',
		W:        20,
		H:        1,
	}
	tb.minW, tb.minW = tb.W, tb.W

	for _, a := range atrs {
		a(tb)
	}

	return tb
}

func (tb *TextBox) Render() {
	tb.drawTextAbove()

	tb.drawBounds()
	tb.drawContent()
	tb.drawCursor()

	tb.drawTextBelow()
}

func (tb *TextBox) Process(e termbox.Event) rscliuitkit.UIElement {
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
		tb.DeleteRuneUnderCursor()
	case termbox.KeyDelete, termbox.KeyCtrlD:
		tb.DeleteRuneUnderCursor()
	case termbox.KeyTab:
		tb.InsertRune('\t')
	case termbox.KeySpace:
		tb.InsertRune(' ')
	case termbox.KeyEnter:
		return tb.callback(string(tb.rText))
	default:
		if e.Ch != 0 {
			tb.InsertRune(e.Ch)
		}
	}
	return tb
}

func (tb *TextBox) InsertRune(r rune) {
	if tb.isExpandable && tb.W <= len(tb.rText)+1 {
		tb.W = common.AddOrMax(tb.W, tb.expandingStep, tb.maxW)
	}

	tb.rText = append(tb.rText, r)
	if tb.GetScreenSpace() <= len(tb.rText) {
		tb.showTextStartCursor++
	}
	tb.editTextCursor++
}

func (tb *TextBox) DeleteRuneUnderCursor() {
	if len(tb.rText) == 0 || tb.editTextCursor == 0 {
		return
	}

	tb.rText = common.RemoveFromSlice(tb.rText, tb.editTextCursor-1)
	if tb.showTextStartCursor > 0 {
		tb.showTextStartCursor--
	}
	tb.editTextCursor--

	if tb.isExpandable && tb.W > tb.minW {
		tb.W = common.SubtractOrMin(tb.W, tb.expandingStep, tb.minW)
	}
}

// GetScreenSpace returns amount of runes that can be displayed at TextBox
func (tb *TextBox) GetScreenSpace() int {
	return tb.W * tb.H
}

func (tb *TextBox) drawTextAbove() {
	cursorX, cursorY := tb.X+tb.W/2, tb.Y-1

	cursorX -= len([]rune(tb.textAboveBox)) / 2
	for _, r := range []rune(tb.textAboveBox) {
		termbox.SetCell(cursorX, cursorY, r, tb.textAboveFg, tb.textAboveBg)
		cursorX += runewidth.RuneWidth(r)
	}
}
func (tb *TextBox) drawTextBelow() {
	cursorX, cursorY := tb.X+tb.W/2, tb.Y+tb.H+2

	cursorX -= len([]rune(tb.textBelowBox)) / 2
	for _, r := range []rune(tb.textBelowBox) {
		termbox.SetCell(cursorX, cursorY, r, tb.textBelowFg, tb.textBelowBg)
		cursorX += runewidth.RuneWidth(r)
	}
}

func (tb *TextBox) drawBounds() {
	//  top
	termbox.SetCell(tb.X, tb.Y, tb.lu, tb.fgInput, tb.bgInput)
	common.FillArea(tb.X+1, tb.Y, tb.W+1, 1, tb.hs, tb.fgInput, tb.bgInput)
	termbox.SetCell(tb.X+tb.W+1, tb.Y, tb.ru, tb.fgInput, tb.bgInput)

	// sides
	for y := tb.Y + 1; y < tb.Y+tb.H+2; y++ {
		termbox.SetCell(tb.X, y, tb.vs, tb.fgInput, tb.bgInput)
		termbox.SetCell(tb.X+tb.W+1, y, tb.vs, tb.fgInput, tb.bgInput)
	}

	// bottom
	termbox.SetCell(tb.X, tb.Y+tb.H+1, tb.ld, tb.fgInput, tb.bgInput)
	common.FillArea(tb.X+1, tb.Y+tb.H+1, tb.W+1, 1, tb.hs, tb.fgInput, tb.bgInput)
	termbox.SetCell(tb.X+tb.W+1, tb.Y+tb.H+1, tb.rd, tb.fgInput, tb.bgInput)

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

		termbox.SetCell(cursorX, cursorY, r, tb.fgInput, tb.bgInput)

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
		termbox.SetCell(cursorX, cursorY, ' ', tb.fgInput, tb.bgInput)
	}
}
func (tb *TextBox) drawCursor() {
	termbox.SetCursor(
		tb.X+1+(tb.editTextCursor-tb.showTextStartCursor)%tb.W,
		tb.Y+1+(tb.editTextCursor-tb.showTextStartCursor)/tb.W,
	)
}
