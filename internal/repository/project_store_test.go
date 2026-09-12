package repository

import (
	"enginer/internal/domain"
	"errors"
	"testing"
)

func TestProjectCreateAndGetByID(t *testing.T) {
	store := NewProjectStore()

	project := domain.NewProject("Дача", "Вентиляция на даче")

	err := store.Create(project)

	if err != nil {
		t.Fatalf("Не удалось создать в сторе %v", err)
	}

	got, err := store.GetByID(project.ID)

	if err != nil {
		t.Fatalf("Не удалось получить систему %v", err)
	}

	if got.Name != project.Name {
		t.Fatalf("Name =%q, ожидалось %q", got.Name, project.Name)
	}
	if got.Description != project.Description {
		t.Fatalf("Description =%q, ожидалось %q", got.Description, project.Description)
	}

}

func TestGotByID_NotFound(t *testing.T) {
	store := NewProjectStore()

	_, err := store.GetByID("adadad")

	if !errors.Is(err, domain.ErrProjectNotFound) {
		t.Fatalf("err = %v, ожидалась ErrProjectNotFound:", err)
	}
}

func TestList(t *testing.T) {
	store := NewProjectStore()

	project1 := domain.NewProject("Дача", "Вентиляция на даче")
	project2 := domain.NewProject("Дача", "Водоснабжение на даче")

	store.Create(project1)
	store.Create(project2)

	got := store.List()

	if len(got) != 2 {
		t.Fatalf("len(got) = %d, ожидалось 2", len(got))
	}

	if _, ok := got[project1.ID]; !ok {
		t.Errorf("проект %s не найден в рузльтате List", project1.ID)
	}

	if _, ok := got[project2.ID]; !ok {
		t.Errorf("проект %s не найден в рузльтате List", project2.ID)
	}

}
