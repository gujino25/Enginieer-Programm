package repository

import (
	"enginer/internal/domain"
	"errors"
	"testing"
)

func TestSystemStore_CreateAndGetByID(t *testing.T) {
	store := NewSystemStore()

	system, err := domain.NewSystem("project-1", "В1", "air", "Дача")

	if err != nil {
		t.Fatalf("не удалось создать систему %v:", err)
	}
	if err := store.Create(system); err != nil {
		t.Fatalf("не удалось создать в сторе: %v", err)
	}
	got, err := store.GetByID(system.ID)
	if err != nil {
		t.Fatalf("не удалось получить систему %v:", err)
	}

	if got.Name != system.Name {
		t.Fatalf("Name =%q, ожидалось %q", got.Name, system.Name)
	}

	if got.Medium != system.Medium {
		t.Fatalf("Medium = %v, ожидалось %v", got.Medium, system.Medium)
	}

}

func TestSystemStore_GetByID_NotFound(t *testing.T) {

	store := NewSystemStore()

	_, err := store.GetByID("asdad")
	if !errors.Is(err, domain.ErrSystemNotFound) {
		t.Fatalf("err = %v, жидалась ErrSystemNotFound:", err)
	}

}

func TestSystemStore_List(t *testing.T) {
	store := NewSystemStore()

	system1, _ := domain.NewSystem("prokect-1", "П1", "air", "Дача приток")
	system2, _ := domain.NewSystem("project-2", "В1", "air", "Дача вытяжка")

	store.Create(system1)
	store.Create(system2)

	got := store.List()

	if len(got) != 2 {
		t.Fatalf("len(got) = %d, ожидалось 2", len(got))
	}

	if _, ok := got[system1.ID]; !ok {
		t.Errorf("система %s не найдена в рузльтате List", system1.ID)
	}

	if _, ok := got[system2.ID]; !ok {
		t.Errorf("система %s не найдена в рузльтате List", system2.ID)
	}

}
