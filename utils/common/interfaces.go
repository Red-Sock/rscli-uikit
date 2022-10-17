package common

type Positioner interface {
	GetPosition() (int, int)
}

type Sizer interface {
	GetSize() (width int, height int)
}
