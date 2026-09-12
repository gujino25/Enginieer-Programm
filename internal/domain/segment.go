package domain

import (
	"time"

	"github.com/google/uuid"
)

type Shape string

const (
	ShapeRect  = "rect"
	ShapeRound = "rond"
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

func NewSegemnt(systemID string, name string, shape Shape, rect *RectGeometry, round *RoundGeometry, length float64) (Segment, error) {
	s := Shape(shape)

	if !s.IsValid() {
		return Segment{}, ErrShapeInvalid
	}
	switch shape {
	case ShapeRect:
		if rect == nil {
			return Segment{}, ErrInvalidGeometry
		}
		if rect.Height <= 0 || rect.Width <= 0 {
			return Segment{}, ErrInvalidGeometry
		}
	case ShapeRound:
		if round == nil {
			return Segment{}, ErrInvalidGeometry
		}
		if round.Diameter <= 0 {
			return Segment{}, ErrInvalidGeometry
		}
	}

	if length <= 0 {
		return Segment{}, ErrInvalidLength
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
