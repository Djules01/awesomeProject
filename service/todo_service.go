package service

import (
	"errors"
	"fmt"
	"awesomeProject/domain"
)

type TodoService interface {
	CreateTodo(todo domain.TodoList) (domain.TodoList, error)
	GetTodoByDate() ([]domain.TodoList, error)
	UpdateTodo(id string, todo domain.TodoList) (domain.TodoList, bool, error)
	DeleteTodo(id string) (bool, error)
}

type TodoRepository interface {
	Create(todo domain.TodoList) (domain.TodoList, error)
	DisplayByDate() ([]domain.TodoList, error)
	Update(id string, todo domain.TodoList) (domain.TodoList, bool, error)
	Delete(id string) (bool, error)
}

type todoService struct {
	repo TodoRepository
}

func NewService(repo TodoRepository) *todoService {
	return &todoService{
		repo: repo,
	}
}

func (s *todoService) CreateTodo(todo domain.TodoList) (domain.TodoList, error) {
	if todo.Titre == "" {
		return domain.TodoList{}, errors.New("pas de titre !")
	}
	if todo.DateCreation == "" {
		return domain.TodoList{}, errors.New("pas de date !")
	}
	if todo.DateEcheance == "" {
		return domain.TodoList{}, errors.New("Veuillez rentrer une échéance")
	}

	todo, err := s.repo.Create(todo)
	if err != nil {
		return domain.TodoList{}, fmt.Errorf("erreur lors de la création")
	}
	return todo, nil
}

func (s *todoService) GetTodoByDate() ([]domain.TodoList, error) {
	return s.repo.DisplayByDate()
}

func (s *todoService) UpdateTodo(id string, todo domain.TodoList) (domain.TodoList, bool, error) {
	if todo.Titre == "" {
		return domain.TodoList{}, false, errors.New("le titre est obligatoire")
	}
	if todo.DateEcheance == "" {
		return domain.TodoList{}, false, errors.New("Veuillez rentrer une échéance")
	}
	return s.repo.Update(id, todo)
}

	func (s *todoService) DeleteTodo(id string) (bool, error) {
		ok, err := s.repo.Delete(id)
		if err != nil {
			return false, fmt.Errorf("erreur lors de la suppression")
		}
		return ok, nil
}
