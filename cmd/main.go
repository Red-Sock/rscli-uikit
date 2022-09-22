package main

import (
	rscliuitkit "github.com/Red-Sock/rscli-uikit"
	"github.com/Red-Sock/rscli-uikit/textbox"
	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	sc := textbox.NewTextBox(
		textbox.NewAttributeSideSymbols('-', '-', '-', '-', '=', '-'),
	)

	sc.Y = 2
	sc.X = 3

	sc.W, sc.H = 20, 1
	q := make(chan struct{})
	rscliuitkit.NewHandler(sc).Start(q)
}

//func eventLoop() {
//	for {
//		switch ev := termbox.PollEvent(); ev.Type {
//		case termbox.EventKey:
//			switch ev.Key {
//			case termbox.KeyEsc:
//				return
//			case termbox.KeyArrowLeft, termbox.KeyCtrlB:
//				edit_box.MoveCursorOneRuneBackward()
//			case termbox.KeyArrowRight, termbox.KeyCtrlF:
//				edit_box.MoveCursorForward()
//			case termbox.KeyBackspace, termbox.KeyBackspace2:
//				edit_box.DeleteRuneBackward()
//			case termbox.KeyDelete, termbox.KeyCtrlD:
//				edit_box.DeleteRuneForward()
//			case termbox.KeyTab:
//				edit_box.InsertRune('\t')
//			case termbox.KeySpace:
//				edit_box.InsertRune(' ')
//			case termbox.KeyCtrlK:
//				edit_box.DeleteTheRestOfTheLine()
//			case termbox.KeyHome, termbox.KeyCtrlA:
//				edit_box.MoveCursorToBeginningOfTheLine()
//			case termbox.KeyEnd, termbox.KeyCtrlE:
//				edit_box.MoveCursorToEndOfTheLine()
//			default:
//				if ev.Ch != 0 {
//					edit_box.InsertRune(ev.Ch)
//				}
//			}
//		case termbox.EventError:
//			panic(ev.Err)
//		}
//		redraw_all()
//	}
//}
