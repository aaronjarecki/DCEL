package dcel

import (
	"fmt"
)

type DCEL struct {
	edges []*HEdge
	faces []*Face
}

func (d *DCEL) String() string {
	return fmt.Sprintf("DCEL")
}

func NewDCEL() *DCEL {
	d := DCEL{}
	d.edges = make([]*HEdge, 0)
	d.faces = make([]*Face, 0)
	return &d
}

func (d *DCEL) ReadEdges() []*HEdge {
	return d.edges
}

func (d *DCEL) BoundingBox(maxX, maxY int) error {
	if maxX == 0 || maxY == 0 {
		return fmt.Errorf("Bounding box has no area")
	}

	v1 := NewVertex(0,0)
	v2 := NewVertex(maxX,0)
	v3 := NewVertex(maxX,maxY)
	v4 := NewVertex(0,maxY)

	f1 := Face{}
	h1 := HEdge{start:v1,face:&f1}
	v1.edge = &h1
	h2 := HEdge{start:v2,face:&f1}
	v2.edge = &h2
	h3 := HEdge{start:v3,face:&f1}
	v3.edge = &h3
	h4 := HEdge{start:v4,face:&f1}
	v4.edge = &h4
	f1.edge = &h1

	h1.next = &h2
	h2.next = &h3
	h3.next = &h4
	h4.next = &h1

	h1.prev = &h4
	h2.prev = &h1
	h3.prev = &h2
	h4.prev = &h3

	d.edges = append(d.edges, []*HEdge{&h1,&h2,&h3,&h4}...)
	d.faces = append(d.faces, &f1)

	for _, edge := range d.edges {
		edge.calculateMetadata()
	}

	return nil
}
