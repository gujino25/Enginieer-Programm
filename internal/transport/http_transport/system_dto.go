package httptransport

import (
	"enginer/internal/domain"
	"errors"
)

type CreateSystemDTO struct {
	Name    string `json:"name"`
	Medium  string `json:"medium"`
	Purpose string `json:"purpose"`
}

func (s CreateSystemDTO) Validate() error {
	if s.Name == "" {
		return errors.New("name is empty")
	}
	if s.Purpose == "" {
		return errors.New("purpose is empty")
	}
	return nil
}

type SystemResponse struct {
	ID        string `json:"id"`
	ProjectID string `json:"project_id"`
	Name      string `json:"name"`
	Medium    string `json:"medium"`
	Purpose   string `json:"purpose"`
	CreatedAt string `json:"created_at"`
}

func toSystemResponse(d domain.System) SystemResponse {
	return SystemResponse{
		ID:        d.ID,
		ProjectID: d.ProjectID,
		Name:      d.Name,
		Medium:    string(d.Medium),
		Purpose:   d.Purpose,
		CreatedAt: d.CreatedAt.Format("02.01.2006"),
	}
}
