package common

// BoundedPositioning defines position of element inside some bounds (upper element)
type BoundedPositioning struct {
	p    Positioner
	s    Sizer
	x, y float32
}

func NewBoundedPositioning(p Positioner, s Sizer, x, y float32) *BoundedPositioning {
	return &BoundedPositioning{
		p: p,
		s: s,
		x: x,
		y: y,
	}
}

func (b *BoundedPositioning) GetPosition() (int, int) {
	sourceX, sourceY := b.p.GetPosition()
	sourceW, sourceH := b.s.GetSize()

	return sourceX + int(b.x*float32(sourceW)), sourceY + int(b.y*float32(sourceH))
}
