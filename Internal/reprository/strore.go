package reprository

import (
	"enginer/Internal/domain"
	"sync"
)

type ProjectStore struct {
	projects map[string]domain.Project
	mtx      sync.RWMutex
}

func Create() *ProjectStore {
	return &ProjectStore{
		projects: make(map[string]domain.Project),
	}
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

	for k, v := range p.projects {
		tmp[k] = v
	}

	return tmp
}
