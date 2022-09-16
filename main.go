package main

import (
	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	redraw_all()
mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break mainloop
			case termbox.KeyArrowLeft, termbox.KeyCtrlB:
				edit_box.MoveCursorOneRuneBackward()
			case termbox.KeyArrowRight, termbox.KeyCtrlF:
				edit_box.MoveCursorOneRuneForward()
			case termbox.KeyBackspace, termbox.KeyBackspace2:
				edit_box.DeleteRuneBackward()
			case termbox.KeyDelete, termbox.KeyCtrlD:
				edit_box.DeleteRuneForward()
			case termbox.KeyTab:
				edit_box.InsertRune('\t')
			case termbox.KeySpace:
				edit_box.InsertRune(' ')
			case termbox.KeyCtrlK:
				edit_box.DeleteTheRestOfTheLine()
			case termbox.KeyHome, termbox.KeyCtrlA:
				edit_box.MoveCursorToBeginningOfTheLine()
			case termbox.KeyEnd, termbox.KeyCtrlE:
				edit_box.MoveCursorToEndOfTheLine()
			default:
				if ev.Ch != 0 {
					edit_box.InsertRune(ev.Ch)
				}
			}
		case termbox.EventError:
			panic(ev.Err)
		}
		redraw_all()
	}
}
