package models_test

import (
	"database/sql/driver"
	"errors"
	"testing"

	"git.todo-app.com/ToDoReactApp/models"
	utils "git.todo-app.com/ToDoReactApp/testutils"
	"github.com/stretchr/testify/assert"
)

func TestSuccessfulInsertToDo(t *testing.T) {
	todoItem := "Shopping"
	mock := utils.GenerateMock()
	mock.ExpectInsertToDoItem(todoItem)
	err := models.ToDoInsert(todoItem, mock.DB())
	assert.NoError(t, err, "Unexpected Error thrown while trying to insert a new todo")
	err = mock.VerifyExpectations()
	assert.NoError(t, err, "Queries were not called")
}

func TestInsertToDoFails(t *testing.T) {
	mock := utils.GenerateMock()
	newToDoItem := "Flour"

	mock.ExpectExecFails(models.InsertQuery, errors.New("bombed"))

	err := models.ToDoInsert(newToDoItem, mock.DB())

	assert.Equal(t, "bombed", err.Error(), "Unexpected Error thrown while trying to insert a new todo")
	err = mock.VerifyExpectations()
	assert.NoError(t, err, "Queries were not called")
}

func TestSuccessfulSelectToDos(t *testing.T) {
	mock := utils.GenerateMock()

	expectedRows := [][]driver.Value{
		{"item1"},
		{"item2"},
	}
	mock.ExpectSelect(expectedRows)
	todoList, err := models.ToDoSelectAll(mock.DB())
	assert.NoError(t, err, "Queries were not called")
	assert.Equal(t, 2, len(todoList), "Should have fetched two todo items")
	err = mock.VerifyExpectations()
	assert.NoError(t, err, "Queries were not called")
}

func TestSelectToDosFailure(t *testing.T) {
	mock := utils.GenerateMock()

	mock.ExpectSelectFails(errors.New("bombed"))
	todos, err := models.ToDoSelectAll(mock.DB())
	assert.NotNil(t, err, "Queries were not called")
	assert.Nil(t, todos, "Returned ToDo list should have been empty")

	err = mock.VerifyExpectations()
	assert.NoError(t, err, "Queries were not called")
}

func TestToDoDeleteSuccess(t *testing.T) {
	mock := utils.GenerateMock()
	todoID := 1

	mock.ExpectDeleteSuccess(todoID)

	err := models.DeleteToDo(mock.DB(), todoID)
	assert.NoError(t, err, "Queries were not called")

	err = mock.VerifyExpectations()
	assert.NoError(t, err, "Queries were not called")
}

func TestToDoDeleteFailure(t *testing.T) {
	mock := utils.GenerateMock()
	todoID := 1

	mock.ExpectExecFails(models.DeleteQuery, errors.New("DB operation failed"))

	err := models.DeleteToDo(mock.DB(), todoID)
	assert.Equal(t, err.Error(), "DB operation failed")

	err = mock.VerifyExpectations()
	assert.NoError(t, err, "Queries were not called")
}
