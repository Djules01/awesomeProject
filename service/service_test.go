package service

import (
	"awesomeProject/domain"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTodoSuccess(t *testing.T) {
	mockRepo := new(MockTodoRepository)
	todoService := NewService(mockRepo)

	input := domain.TodoList{
		Titre:        "Tester service",
		DateCreation: "2026-06-18",
		DateEcheance: "2026-06-25",
	}

	expected := input
	expected.ID = 1

	mockRepo.On("Create", input).Return(expected, nil)

	result, err := todoService.CreateTodo(input)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestCreateTodoWithoutTitleReturnsError(t *testing.T) {
	mockRepo := new(MockTodoRepository)
	todoService := NewService(mockRepo)

	input := domain.TodoList{
		DateCreation: "2026-06-18",
		DateEcheance: "2026-06-25",
	}

	result, err := todoService.CreateTodo(input)

	assert.Error(t, err)
	assert.Equal(t, "pas de titre !", err.Error())
	assert.Equal(t, domain.TodoList{}, result)
	mockRepo.AssertNotCalled(t, "Create")
}

func TestCreateTodoWithoutCreationDateReturnsError(t *testing.T) {
	mockRepo := new(MockTodoRepository)
	todoService := NewService(mockRepo)

	input := domain.TodoList{
		Titre:        "Tester service",
		DateEcheance: "2026-06-25",
	}

	result, err := todoService.CreateTodo(input)

	assert.Error(t, err)
	assert.Equal(t, "pas de date !", err.Error())
	assert.Equal(t, domain.TodoList{}, result)
	mockRepo.AssertNotCalled(t, "Create")
}

func TestCreateTodoWithoutDueDateReturnsError(t *testing.T) {
	mockRepo := new(MockTodoRepository)
	todoService := NewService(mockRepo)

	input := domain.TodoList{
		Titre:        "Tester service",
		DateCreation: "2026-06-18",
	}

	result, err := todoService.CreateTodo(input)

	assert.Error(t, err)
	assert.Equal(t, "Veuillez rentrer une échéance", err.Error())
	assert.Equal(t, domain.TodoList{}, result)
	mockRepo.AssertNotCalled(t, "Create")
}

func TestCreateTodoRepositoryError(t *testing.T) {
	mockRepo := new(MockTodoRepository)
	todoService := NewService(mockRepo)

	input := domain.TodoList{
		Titre:        "Tester service",
		DateCreation: "2026-06-18",
		DateEcheance: "2026-06-25",
	}

	expectedErr := errors.New("erreur repository")
	mockRepo.On("Create", input).Return(domain.TodoList{}, expectedErr)

	result, err := todoService.CreateTodo(input)

	assert.ErrorIs(t, err, expectedErr)
	assert.Equal(t, domain.TodoList{}, result)
	mockRepo.AssertExpectations(t)
}

func TestGetTodoByDateSuccess(t *testing.T) {
	mockRepo := new(MockTodoRepository)
	todoService := NewService(mockRepo)

	expected := []domain.TodoList{
		{
			ID:           1,
			Titre:        "Tester service",
			DateCreation: "2026-06-18",
			DateEcheance: "2026-06-25",
		},
	}

	mockRepo.On("DisplayByDate").Return(expected, nil)

	result, err := todoService.GetTodoByDate()

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestUpdateTodoSuccess(t *testing.T) {
	mockRepo := new(MockTodoRepository)
	todoService := NewService(mockRepo)

	input := domain.TodoList{
		Titre:        "Todo modifiee",
		DateCreation: "2026-06-18",
		DateEcheance: "2026-06-30",
		Completer:    true,
	}

	expected := input
	expected.ID = 1

	mockRepo.On("Update", 1, input).Return(expected, true, nil)

	result, ok, err := todoService.UpdateTodo(1, input)

	assert.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestDeleteTodoSuccess(t *testing.T) {
	mockRepo := new(MockTodoRepository)
	todoService := NewService(mockRepo)

	mockRepo.On("Delete", 1).Return(true, nil)

	ok, err := todoService.DeleteTodo(1)

	assert.NoError(t, err)
	assert.True(t, ok)
	mockRepo.AssertExpectations(t)
}
