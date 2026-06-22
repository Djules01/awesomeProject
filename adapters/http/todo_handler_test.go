package httpadapter

import (
	"awesomeProject/domain"
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTodoService struct {
	mock.Mock
}

func (m *MockTodoService) CreateTodo(todo domain.TodoList) (domain.TodoList, error) {
	args := m.Called(todo)
	return args.Get(0).(domain.TodoList), args.Error(1)
}

func (m *MockTodoService) GetTodoByDate() ([]domain.TodoList, error) {
	args := m.Called()
	return args.Get(0).([]domain.TodoList), args.Error(1)
}

func (m *MockTodoService) UpdateTodo(id int, todo domain.TodoList) (domain.TodoList, bool, error) {
	args := m.Called(id, todo)
	return args.Get(0).(domain.TodoList), args.Bool(1), args.Error(2)
}

func (m *MockTodoService) DeleteTodo(id int) (bool, error) {
	args := m.Called(id)
	return args.Bool(0), args.Error(1)
}

func TestCreateTodoHandlerSuccess(t *testing.T) {
	mockService := new(MockTodoService)
	todoHandler := NewHandler(mockService)

	body := bytes.NewBufferString(`{
		"titre":"Tester handler",
		"datecreation":"2026-06-18",
		"dateecheance":"2026-06-25"
	}`)

	input := domain.TodoList{
		Titre:        "Tester handler",
		DateCreation: "2026-06-18",
		DateEcheance: "2026-06-25",
	}

	expected := domain.TodoList{
		ID:           1,
		Titre:        "Tester handler",
		DateCreation: "2026-06-18",
		DateEcheance: "2026-06-25",
		Completer:    false,
	}

	mockService.On("CreateTodo", input).Return(expected, nil)

	req := httptest.NewRequest(http.MethodPost, "/Creation", body)
	w := httptest.NewRecorder()

	todoHandler.CreateTodo(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), `"id":1`)
	assert.Contains(t, w.Body.String(), `"titre":"Tester handler"`)

	mockService.AssertExpectations(t)
}

func TestGetTodosByDateHandlerSuccess(t *testing.T) {
	mockService := new(MockTodoService)
	todoHandler := NewHandler(mockService)

	expected := []domain.TodoList{
		{
			ID:           1,
			Titre:        "Tester handler",
			DateCreation: "2026-06-18",
			DateEcheance: "2026-06-25",
			Completer:    false,
		},
	}

	mockService.On("GetTodoByDate").Return(expected, nil)

	req := httptest.NewRequest(http.MethodGet, "/AfficherParDate", nil)
	w := httptest.NewRecorder()

	todoHandler.GetTodosByDate(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"id":1`)
	assert.Contains(t, w.Body.String(), `"titre":"Tester handler"`)

	mockService.AssertExpectations(t)
}

func TestUpdateTodoHandlerSuccess(t *testing.T) {
	mockService := new(MockTodoService)
	todoHandler := NewHandler(mockService)

	body := bytes.NewBufferString(`{
		"titre":"Todo modifiee",
		"datecreation":"2026-06-18",
		"dateecheance":"2026-06-30",
		"completer":true
	}`)

	input := domain.TodoList{
		Titre:        "Todo modifiee",
		DateCreation: "2026-06-18",
		DateEcheance: "2026-06-30",
		Completer:    true,
	}

	expected := input
	expected.ID = 1

	mockService.On("UpdateTodo", 1, input).Return(expected, true, nil)

	req := httptest.NewRequest(http.MethodPut, "/Modifier/1", body)
	req = withURLParam(req, "id", "1")
	w := httptest.NewRecorder()

	todoHandler.UpdateTodo(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"id":1`)
	assert.Contains(t, w.Body.String(), `"titre":"Todo modifiee"`)
	assert.Contains(t, w.Body.String(), `"completer":true`)

	mockService.AssertExpectations(t)
}

func TestUpdateTodoHandlerNotFound(t *testing.T) {
	mockService := new(MockTodoService)
	todoHandler := NewHandler(mockService)

	body := bytes.NewBufferString(`{
		"titre":"Todo absente",
		"datecreation":"2026-06-18",
		"dateecheance":"2026-06-30"
	}`)

	input := domain.TodoList{
		Titre:        "Todo absente",
		DateCreation: "2026-06-18",
		DateEcheance: "2026-06-30",
	}

	mockService.On("UpdateTodo", 99, input).Return(domain.TodoList{}, false, nil)

	req := httptest.NewRequest(http.MethodPut, "/Modifier/99", body)
	req = withURLParam(req, "id", "99")
	w := httptest.NewRecorder()

	todoHandler.UpdateTodo(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "Todo introuvable")

	mockService.AssertExpectations(t)
}

func TestDeleteTodoHandlerSuccess(t *testing.T) {
	mockService := new(MockTodoService)
	todoHandler := NewHandler(mockService)

	mockService.On("DeleteTodo", 1).Return(true, nil)

	req := httptest.NewRequest(http.MethodDelete, "/Delete/1", nil)
	req = withURLParam(req, "id", "1")
	w := httptest.NewRecorder()

	todoHandler.DeleteTodo(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Empty(t, w.Body.String())

	mockService.AssertExpectations(t)
}

func TestDeleteTodoHandlerNotFound(t *testing.T) {
	mockService := new(MockTodoService)
	todoHandler := NewHandler(mockService)

	mockService.On("DeleteTodo", 99).Return(false, nil)

	req := httptest.NewRequest(http.MethodDelete, "/Delete/99", nil)
	req = withURLParam(req, "id", "99")
	w := httptest.NewRecorder()

	todoHandler.DeleteTodo(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "Todo introuvable")

	mockService.AssertExpectations(t)
}

func withURLParam(req *http.Request, key string, value string) *http.Request {
	routeContext := chi.NewRouteContext()
	routeContext.URLParams.Add(key, value)

	ctx := context.WithValue(req.Context(), chi.RouteCtxKey, routeContext)
	return req.WithContext(ctx)
}
