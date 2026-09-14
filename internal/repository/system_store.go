package repository

import (
	"enginer/internal/domain"
	"maps"
	"sync"
)

type SystemStore struct {
	systems map[string]domain.System
	mtx     sync.RWMutex
}

func NewSystemStore() *SystemStore {

	return &SystemStore{
		systems: make(map[string]domain.System),
	}

}

func (s *SystemStore) Create(system domain.System) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	if _, ok := s.systems[system.ID]; ok {
		return domain.ErrSystemAlreadyExists
	}
	s.systems[system.ID] = system

	return nil
}

func (s *SystemStore) GetByID(id string) (domain.System, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	system, ok := s.systems[id]

	if !ok {
		return domain.System{}, domain.ErrSystemNotFound
	}

	return system, nil
}

func (s *SystemStore) ListByProject(id string) map[string]domain.System {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	tmp := make(map[string]domain.System)

	for k, v := range s.systems {
		if v.ProjectID == id {
			tmp[k] = v
		}
	}

	return tmp
}

func (s *SystemStore) List() map[string]domain.System {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	tmp := make(map[string]domain.System, len(s.systems))

	maps.Copy(tmp, s.systems)

	return tmp
}
