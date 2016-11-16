package testutils

import (
	"database/sql"
	"github.com/stretchr/testify/mock"

	"git.todo-app.com/ToDoReactApp/models"
)

type MockToDoRepository struct {
	mock.Mock
}

func (m *MockToDoRepository) Delete(db *sql.DB, id int) error {
	args := m.Called(db, id)
	return args.Error(0)
}

func (m *MockToDoRepository) Insert(value string, db *sql.DB) (string,error) {
	args := m.Called(value, db)
	return args.String(0),args.Error(1)
}

func (m *MockToDoRepository) Select(db *sql.DB) ([]models.ToDo, error) {
	args := m.Called(db)
	var todos []models.ToDo
	var err error
	if args[0] != nil {
		todos = args[0].([]models.ToDo)
	}
	if args[1] != nil {
		err = args[1].(error)
	}
	return todos, err
}
