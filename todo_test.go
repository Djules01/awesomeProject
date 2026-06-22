package main

import (
	"awesomeProject/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTodoListModel(t *testing.T) {
	todo := domain.TodoList{
		ID:           "todo-id-1",
		Titre:        "Apprendre Go",
		DateCreation: "2026-06-17",
		DateEcheance: "2026-06-20",
		Completer:    false,
	}

	assert.Equal(t, "todo-id-1", todo.ID)
	assert.Equal(t, "Apprendre Go", todo.Titre)
	assert.Equal(t, "2026-06-17", todo.DateCreation)
	assert.Equal(t, "2026-06-20", todo.DateEcheance)
	assert.False(t, todo.Completer)
}
