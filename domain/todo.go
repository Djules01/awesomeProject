package domain

type TodoList struct {
	ID           string `json:"id"`
	Titre        string `json:"titre"`
	DateCreation string `json:"datecreation"`
	DateEcheance string `json:"dateecheance"`
	Completer    bool   `json:"completer"`
}

type TodoService interface {
	CreateTodo(todo TodoList) (TodoList, error)
	GetTodoByDate() ([]TodoList, error)
	UpdateTodo(id string, todo TodoList) (TodoList, bool, error)
	DeleteTodo(id string) (bool, error)
}
