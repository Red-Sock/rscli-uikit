package endscreen

import (
	rscliuitkit "github.com/Red-Sock/rscli-uikit"
	"github.com/nsf/termbox-go"
)

// EndScreen Temporary decision for ending cli app
type EndScreen struct {
	rscliuitkit.UIElement
}

func (e *EndScreen) Render() {
	e.UIElement.Render()
}

func (e *EndScreen) Process(ev termbox.Event) rscliuitkit.UIElement {
	return nil
}

func (e *EndScreen) SetPreviousScreen(element rscliuitkit.UIElement) {}
