package common

import "github.com/nsf/termbox-go"

type Positioner interface {
	GetPosition() (int, int)
}

type AbsolutePositioning struct {
	X, Y int
}

func (a *AbsolutePositioning) GetPosition() (int, int) {
	return a.X, a.Y
}

type RelativePositioning struct {
	x, y float32
}

func NewRelativePositioning(x, y float32) *RelativePositioning {
	return &RelativePositioning{
		x: x,
		y: y,
	}
}

func (r *RelativePositioning) GetPosition() (int, int) {
	w, h := termbox.Size()

	return int(float32(w) * r.x), int(float32(h) * r.y)
}
