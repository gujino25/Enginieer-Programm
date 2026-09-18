package httptransport

import (
	"enginer/internal/domain"
	"errors"
)

type CreateProjectDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (p CreateProjectDTO) Validate() error {
	if p.Name == "" {
		return errors.New("name is required")
	}
	if p.Description == "" {
		return errors.New("description is required")
	}
	return nil
}

type ProjectResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

func toProjectResponse(p domain.Project) ProjectResponse {
	return ProjectResponse{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		CreatedAt:   p.CreatedAt.Format("02.01.2006"),
	}
}
