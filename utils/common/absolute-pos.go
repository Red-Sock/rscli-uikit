package common

type AbsolutePositioning struct {
	X, Y int
}

func (a *AbsolutePositioning) GetPosition() (int, int) {
	return a.X, a.Y
}
