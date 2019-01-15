package dcel

import (
	"math"
)

type Point struct {
	x int
	y int
}

func NewPoint(x,y int) Point {
	p := Point{x,y}
	return p
}

type Radians float64

type Vertex struct {
	id string
	point Point
	edge *HEdge
}

func NewVertex(x,y int) *Vertex {
	v := Vertex{point:Point{x,y}}
	return &v
}

type HEdge struct {
	id string
	start *Vertex
	next *HEdge
	prev *HEdge
	twin *HEdge
	face *Face

	end *Vertex
	heading Radians
}

func (h *HEdge) calculateMetadata() {
	h.end = h.next.start
	heading := math.Atan2(float64(h.end.point.y-h.start.point.y),float64(h.end.point.x-h.start.point.x))
	h.heading = Radians(heading)
}

func (h *HEdge) ReadHeading() Radians {
	return h.heading
}

type Face struct {
	id string
	edge *HEdge
}

func (f *Face) DivideAtPoints(start, end Vertex) error {
	return nil
}

func (f *Face) getEdges() []*HEdge {
	edges := make([]*HEdge, 0)
	startEdge := f.edge
	edge = startEdge
	edges = append(edges, edge)
	for ; edge.next != startEdge; {
		edges = append(edges, edge)
		edge = edge.next
	}
}

