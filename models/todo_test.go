package models_test

import (
	"database/sql/driver"
	"testing"

	utils "git.todo-app.com/ToDoReactApp/testutils"
	"github.com/stretchr/testify/assert"
	"git.todo-app.com/ToDoReactApp/models"
)

func TestSuccessfulInsertToDo(t *testing.T) {
	todoItem := "Shopping"
	mock := utils.GenerateMock()
	mock.ExpectInsertToDoItem(todoItem)
	models.ToDoInsert(todoItem, mock.DB())
	err := mock.VerifyExpectations()
	assert.NoError(t, err, "Queries were not called")
}

func TestSuccessfulSelectToDos(t *testing.T) {
	mock := utils.GenerateMock()

	expectedRows := [][]driver.Value{
		{"item1"},
		{"item2"},
	}
	mock.ExpectSelect(expectedRows)
	_ = models.ToDoSelectAll(mock.DB())

	err := mock.VerifyExpectations()
	assert.NoError(t, err, "Queries were not called")
}
