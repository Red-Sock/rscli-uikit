package rscliuitkit

import (
	"github.com/Red-Sock/rscli-uikit/utils/common"
	"github.com/nsf/termbox-go"
	"log"
)

func init() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
}

type Handler interface {
	Start(q <-chan struct{})
}

type UIElement interface {
	Render()
	Process(e termbox.Event) UIElement
	SetPreviousScreen(element UIElement)
}

type Labeler interface {
	UIElement
	GetSize() (w, h int)
	SetPosition(p common.Positioner)
}

type Pixel struct {
	x, y   int
	char   rune
	fg, bg termbox.Attribute
}

type handler struct {
	activeScreen UIElement
}

func NewHandler(screen UIElement) Handler {
	return &handler{
		activeScreen: screen,
	}
}

func (h *handler) Start(q <-chan struct{}) {
	defer termbox.Close()

	event := make(chan termbox.Event)
	go func() {
		for {
			event <- termbox.PollEvent()
		}
	}()
	h.draw()
	for {
		select {
		case e := <-event:
			oldScreen := h.activeScreen
			h.activeScreen = h.activeScreen.Process(e)
			if h.activeScreen == nil {
				return
			}
			h.activeScreen.SetPreviousScreen(oldScreen)
			h.draw()
		case <-q:
			return
		}
	}
}

func (h *handler) draw() {
	err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if err != nil {
		log.Fatal()
	}

	h.activeScreen.Render()

	err = termbox.Flush()
	if err != nil {
		log.Fatal(err)
	}
}
