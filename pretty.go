package dcel

import (
	"fmt"
)

func (p Point) String() string {
	return fmt.Sprintf("(%v,%v)", p.x, p.y)
}

func (r Radians) String() string {
	return fmt.Sprintf("θ(%f)", r)
}

func (v Vertex) String() string {
	return fmt.Sprintf("V%v%v",v.id, v.point)
}

func (h HEdge) String() string {
	return fmt.Sprintf("E%v - %v > %v", h.id, h.start, h.next.start)
}

func (f Face) String() string {
	return fmt.Sprintf("F%v", f.id)
}
