package service

import (
	"enginer/internal/domain"
	"enginer/internal/repository"
)

type projectGetter interface {
	GetByID(id string) (domain.Project, error)
}

type SystemService struct {
	systemStore  *repository.SystemStore
	projectStore projectGetter
}

func NewSystemService(systemStore *repository.SystemStore, projectStore projectGetter) *SystemService {
	return &SystemService{
		systemStore:  systemStore,
		projectStore: projectStore,
	}
}

func (s *SystemService) CreateSystem(projectID, name string, medium domain.Medium, purpose string) (domain.System, error) {
	if _, err := s.projectStore.GetByID(projectID); err != nil {
		return domain.System{}, err
	}

	system, err := domain.NewSystem(projectID, name, string(medium), purpose)
	if err != nil {
		return domain.System{}, err
	}

	if err := s.systemStore.Create(system); err != nil {
		return domain.System{}, err
	}

	return system, nil
}

func (s *SystemService) ListByProject(projectID string) (map[string]domain.System, error) {

	if _, err := s.projectStore.GetByID(projectID); err != nil {
		return nil, err
	}

	return s.systemStore.ListByProject(projectID), nil
}
