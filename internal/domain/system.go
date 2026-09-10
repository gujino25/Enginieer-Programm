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

func (m Medium) MediumIsValid() bool {
	switch m {
	case MediumAir, MediumWater:
		return true
	}
	return false
}

func NewSystem(id string, name string, medium string, purpose string) (System, error) {
	if !MediumAir.MediumIsValid() && !MediumWater.MediumIsValid() {
		return System{}, ErrMediumInvailid
	}
	system := System{
		ID:        uuid.NewString(),
		ProjectID: id,
		Name:      name,
		Medium:    Medium(medium),
		Purpose:   purpose,
		CreatedAt: time.Now(),
	}
	return system, nil
}
