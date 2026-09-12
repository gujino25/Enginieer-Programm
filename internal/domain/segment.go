package domain

import (
	"time"

	"github.com/google/uuid"
)

type Shape string

const (
	ShapeRect  = "rect"
	ShapeRound = "round"
)

type RectGeometry struct {
	Width  int
	Height int
}

type RoundGeometry struct {
	Diameter int
}

type Segment struct {
	ID        string
	SystemID  string
	Name      string
	Shape     Shape
	Rect      *RectGeometry
	Round     *RoundGeometry
	Length    float64
	CreatedAt time.Time
}

func (s Shape) IsValid() bool {
	switch s {
	case ShapeRect, ShapeRound:
		return true
	}
	return false
}

func NewSegment(systemID string, name string, shape Shape, rect *RectGeometry, round *RoundGeometry, length float64) (Segment, error) {
	s := Shape(shape)

	if !s.IsValid() {
		return Segment{}, ErrShapeInvalid
	}
	switch shape {
	case ShapeRect:
		if round != nil {
			return Segment{}, ErrGeometryInvalid
		}
		if rect == nil {
			return Segment{}, ErrGeometryInvalid
		}
		if rect.Height <= 0 || rect.Width <= 0 {
			return Segment{}, ErrGeometryInvalid
		}
	case ShapeRound:
		if rect != nil {
			return Segment{}, ErrGeometryInvalid
		}
		if round == nil {
			return Segment{}, ErrGeometryInvalid
		}
		if round.Diameter <= 0 {
			return Segment{}, ErrGeometryInvalid
		}
	}

	if length <= 0 {
		return Segment{}, ErrLengthInvalid
	}

	return Segment{
		ID:        uuid.NewString(),
		SystemID:  systemID,
		Name:      name,
		Shape:     shape,
		Rect:      rect,
		Round:     round,
		Length:    length,
		CreatedAt: time.Now(),
	}, nil
}
