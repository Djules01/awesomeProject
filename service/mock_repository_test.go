package service

import (
	"awesomeProject/domain"

	"github.com/stretchr/testify/mock"
)

type MockTodoRepository struct {
	mock.Mock
}

func (m *MockTodoRepository) Create(todo domain.TodoList) (domain.TodoList, error) {
	args := m.Called(todo)
	return args.Get(0).(domain.TodoList), args.Error(1)
}

func (m *MockTodoRepository) DisplayByDate() ([]domain.TodoList, error) {
	args := m.Called()
	return args.Get(0).([]domain.TodoList), args.Error(1)
}

func (m *MockTodoRepository) Update(id int, todo domain.TodoList) (domain.TodoList, bool, error) {
	args := m.Called(id, todo)
	return args.Get(0).(domain.TodoList), args.Bool(1), args.Error(2)
}

func (m *MockTodoRepository) Delete(id int) (bool, error) {
	args := m.Called(id)
	return args.Bool(0), args.Error(1)
}
