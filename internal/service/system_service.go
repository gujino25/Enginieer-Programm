package service

import (
	"enginer/internal/domain"
	"enginer/internal/repository"
)

type SystemService struct {
	systemStore  *repository.SystemStore
	projectStore *repository.ProjectStore
}

func (s *SystemService) CreateSystemService(projectID, name string, medium domain.Medium, pupose string) (domain.System, error) {
	if _, err := s.projectStore.GetByID(projectID); err != nil {
		return domain.System{}, err
	}

	system, err := domain.NewSystem(projectID, name, string(medium), pupose)
	if err != nil {
		return domain.System{}, err
	}

	if err := s.systemStore.CreateSystemStore(system); err != nil {
		return domain.System{}, err
	}

	return system, nil
}
