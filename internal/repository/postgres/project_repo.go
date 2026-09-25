package postgres

import (
	"context"
	"enginer/internal/domain"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProjectRepo struct {
	pool *pgxpool.Pool
}

func NewProjectRepo(pool *pgxpool.Pool) *ProjectRepo {
	return &ProjectRepo{
		pool: pool,
	}
}

func (r *ProjectRepo) Create(ctx context.Context, project domain.Project) error {
	const query = `INSERT INTO projects (id, name, description, created_at) VALUES ($1, $2, $3, $4)`
	_, err := r.pool.Exec(ctx, query, project.ID, project.Name, project.Description, project.CreatedAt)
	if err != nil {
		var pgrErr *pgconn.PgError
		if errors.As(err, &pgrErr) && pgrErr.Code == "23505" {
			return domain.ErrProjectAlreadyExists
		}
		return fmt.Errorf("create project: %w", err)
	}
	return nil
}

func (r *ProjectRepo) GetByID(ctx context.Context, id string) (domain.Project, error) {
	const query = `SELECT id, name, description, created_at FROM projects WHERE id = $1`

	var project domain.Project
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&project.ID,
		&project.Name,
		&project.Description,
		&project.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Project{}, domain.ErrProjectNotFound
		}
		return domain.Project{}, fmt.Errorf("get project: %w", err)
	}
	return project, nil
}

func (r *ProjectRepo) List(ctx context.Context) ([]domain.Project, error) {
	const query = `SELECT id, name, description, created_at FROM projects ORDER BY created_at, id`
	rows, err := r.pool.Query(ctx, query)

	if err != nil {
		return nil, fmt.Errorf("get all projects: %w", err)
	}
	defer rows.Close()

	var projects []domain.Project
	for rows.Next() {
		var p domain.Project
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan projects: %w", err)
		}
		projects = append(projects, p)
	}
	return projects, rows.Err()
}
