package domain

import (
	"errors"
	"testing"
)

func TestNewSegment(t *testing.T) {
	tests := []struct {
		name    string
		shape   Shape
		rect    *RectGeometry
		round   *RoundGeometry
		length  float64
		wantErr error
	}{
		{
			name:    "валидный прямоугольный",
			shape:   ShapeRect,
			rect:    &RectGeometry{Width: 400, Height: 250},
			round:   nil,
			length:  5.2,
			wantErr: nil,
		},
		{
			name:    "валидный круглый",
			shape:   ShapeRound,
			rect:    nil,
			round:   &RoundGeometry{Diameter: 200},
			length:  3.1,
			wantErr: nil,
		},
		{
			name:    "прямоугольный без геометрии",
			shape:   ShapeRect,
			rect:    nil,
			round:   nil,
			length:  5.2,
			wantErr: ErrGeometryInvalid,
		},
		{
			name:    "противоречие: rect с round одновременно",
			shape:   ShapeRect,
			rect:    &RectGeometry{Width: 400, Height: 250},
			round:   &RoundGeometry{Diameter: 200},
			length:  5.2,
			wantErr: ErrGeometryInvalid,
		},
		{
			name:    "нулевая длина",
			shape:   ShapeRect,
			rect:    &RectGeometry{Width: 400, Height: 250},
			round:   nil,
			length:  0,
			wantErr: ErrLengthInvalid,
		},
		{
			name:    "невалидная форма",
			shape:   "жопа",
			rect:    nil,
			round:   nil,
			length:  5.2,
			wantErr: ErrShapeInvalid,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewSegment("system-1", "тест", tt.shape, tt.rect, tt.round, tt.length)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("err = %v, ожидалось %v", err, tt.wantErr)
			}
		})
	}
}
