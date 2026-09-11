package domain

import (
	"time"

	"github.com/google/uuid"
)

type Medium string

const (
	MediumAir   Medium = "air"
	MediumWater Medium = "water"
)

type System struct {
	ID        string
	ProjectID string
	Name      string
	Medium    Medium
	Purpose   string
	CreatedAt time.Time
}

func (m Medium) IsValid() bool {
	switch m {
	case MediumAir, MediumWater:
		return true
	}
	return false
}

func NewSystem(projectid string, name string, medium string, purpose string) (System, error) {
	m := Medium(medium)

	if !m.IsValid() {
		return System{}, ErrMediumInvalid
	}
	system := System{
		ID:        uuid.NewString(),
		ProjectID: projectid,
		Name:      name,
		Medium:    m,
		Purpose:   purpose,
		CreatedAt: time.Now(),
	}
	return system, nil
}
