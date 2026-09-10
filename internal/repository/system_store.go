package repository

import (
	"enginer/internal/domain"
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

func (s *SystemStore) CreateSystemStore(system domain.System) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	if _, ok := s.systems[system.ID]; ok {
		return domain.ErrSystemAlreadyExists
	}
	s.systems[system.ID] = system

	return nil
}
