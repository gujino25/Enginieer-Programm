package repository

import (
	"enginer/internal/domain"
	"maps"
	"sync"
)

type ProjectStore struct {
	projects map[string]domain.Project
	mtx      sync.RWMutex
}

func NewProjectSotre() *ProjectStore {
	return &ProjectStore{
		projects: make(map[string]domain.Project),
	}
}

func (p *ProjectStore) Create(project domain.Project) error {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	if _, ok := p.projects[project.Name]; ok {
		return domain.ErrProjectAlreadyExist
	}
	p.projects[project.Name] = project

	return nil
}

func (p *ProjectStore) GetByID(id string) (domain.Project, error) {
	p.mtx.RLock()
	defer p.mtx.RUnlock()

	project, ok := p.projects[id]

	if !ok {
		return domain.Project{}, domain.ErrProjectNotFound
	}
	return project, nil
}

func (p *ProjectStore) Lsit() map[string]domain.Project {
	p.mtx.RLock()
	defer p.mtx.RUnlock()

	tmp := make(map[string]domain.Project, len(p.projects))

	maps.Copy(tmp, p.projects)

	return tmp
}
