package httptransport

import (
	"enginer/internal/domain"
	"errors"
)

type RectGeometryDTO struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type RoundGeometryDTO struct {
	Diameter int `json:"diameter"`
}

type CreateSegmentDTO struct {
	Name   string            `json:"name"`
	Shape  string            `json:"shape"`
	Rect   *RectGeometryDTO  `json:"rect"`
	Round  *RoundGeometryDTO `json:"round"`
	Length float64           `json:"length"`
}

func (s CreateSegmentDTO) Validate() error {
	if s.Name == "" {
		return errors.New("name is empty")
	}

	return nil
}

type SegmentResponse struct {
	ID        string            `json:"id"`
	SystemID  string            `json:"system_id"`
	Name      string            `json:"name"`
	Shape     string            `json:"shape"`
	Rect      *RectGeometryDTO  `json:"rect"`
	Round     *RoundGeometryDTO `json:"round"`
	Length    float64           `json:"length"`
	CreatedAt string            `json:"created_at"`
}

func toSegmentResponse(d domain.Segment) SegmentResponse {
	var rect *RectGeometryDTO
	if d.Rect != nil {
		rect = &RectGeometryDTO{
			Width:  d.Rect.Width,
			Height: d.Rect.Height,
		}
	}

	var round *RoundGeometryDTO

	if d.Round != nil {
		round = &RoundGeometryDTO{
			Diameter: d.Round.Diameter,
		}
	}

	return SegmentResponse{
		ID:        d.ID,
		SystemID:  d.SystemID,
		Name:      d.Name,
		Shape:     string(d.Shape),
		Rect:      rect,
		Round:     round,
		Length:    d.Length,
		CreatedAt: d.CreatedAt.Format("02.01.2006"),
	}
}
