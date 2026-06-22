package service

import (
	"awesomeProject/domain"
	"errors"
)

type TodoService interface {
	CreateTodo(todo domain.TodoList) (domain.TodoList, error)
	GetTodoByDate() ([]domain.TodoList, error)
	UpdateTodo(id int, todo domain.TodoList) (domain.TodoList, bool, error)
	DeleteTodo(id int) (bool, error)
}

type TodoRepository interface {
	Create(todo domain.TodoList) (domain.TodoList, error)
	DisplayByDate() ([]domain.TodoList, error)
	Update(id int, todo domain.TodoList) (domain.TodoList, bool, error)
	Delete(id int) (bool, error)
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

	return s.repo.Create(todo)

}

func (s *todoService) GetTodoByDate() ([]domain.TodoList, error) {
	return s.repo.DisplayByDate()
}

func (s *todoService) UpdateTodo(id int, todo domain.TodoList) (domain.TodoList, bool, error) {
	return s.repo.Update(id, todo)
}

func (s *todoService) DeleteTodo(id int) (bool, error) {
	return s.repo.Delete(id)
}
