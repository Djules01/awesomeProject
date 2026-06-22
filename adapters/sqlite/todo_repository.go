package sqliteadapter

import (
	"awesomeProject/domain"
	"database/sql"

	_ "modernc.org/sqlite"
)

type Repository struct {
	DB *sql.DB
}

func InitDB(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	query := `
	CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		titre TEXT NOT NULL,
		date_creation TEXT NOT NULL,
		date_echeance TEXT NOT NULL,
		completer INTEGER NOT NULL DEFAULT 0
	);`

	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) Create(todo domain.TodoList) (domain.TodoList, error) {
	result, err := r.DB.Exec("INSERT INTO todos (titre, date_creation, date_echeance, completer) VALUES (?, ?, ?, ?)",
		todo.Titre,
		todo.DateCreation,
		todo.DateEcheance,
		0,
	)
	if err != nil {
		return domain.TodoList{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.TodoList{}, err
	}

	todo.ID = int(id)
	todo.Completer = false

	return todo, nil
}

func (r *Repository) DisplayByDate() ([]domain.TodoList, error) {
	rows, err := r.DB.Query(`SELECT id, titre, date_creation, date_echeance, completer
	FROM todos
	ORDER BY date_echeance ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := []domain.TodoList{}
	for rows.Next() {
		var todo domain.TodoList
		var completerInt int

		err := rows.Scan(
			&todo.ID,
			&todo.Titre,
			&todo.DateCreation,
			&todo.DateEcheance,
			&completerInt)
		if err != nil {
			return nil, err
		}

		todo.Completer = completerInt == 1
		list = append(list, todo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return list, nil
}

func (r *Repository) Update(id int, todo domain.TodoList) (domain.TodoList, bool, error) {
	completerInt := 0
	if todo.Completer {
		completerInt = 1
	}

	result, err := r.DB.Exec("UPDATE todos SET titre = ?, date_creation = ?, date_echeance = ?, completer = ? WHERE id = ?",
		todo.Titre,
		todo.DateCreation,
		todo.DateEcheance,
		completerInt,
		id)
	if err != nil {
		return domain.TodoList{}, false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return domain.TodoList{}, false, err
	}

	if rowsAffected == 0 {
		return domain.TodoList{}, false, nil
	}

	todo.ID = id

	return todo, true, nil
}

func (r *Repository) Delete(id int) (bool, error) {
	result, err := r.DB.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return rowsAffected > 0, nil
}
