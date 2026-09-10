package domain

import (
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ID          string
	Name        string
	Description string
	CreatedAt   time.Time
}

func NewProject(name string, description string) Project {

	project := Project{
		ID:          uuid.NewString(),
		Name:        name,
		Description: description,
		CreatedAt:   time.Now(),
	}
	return project
}
