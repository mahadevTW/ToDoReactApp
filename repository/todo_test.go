package repository_test

import (
	"database/sql/driver"
	"errors"
	"testing"

	repo "git.todo-app.com/ToDoReactApp/repository"
	utils "git.todo-app.com/ToDoReactApp/testutils"
	"github.com/stretchr/testify/assert"
)

func TestToDoDeleteSuccess(t *testing.T) {
	mock := utils.GenerateMock()
	todoID := 1

	mock.ExpectDeleteSuccess(todoID)

	todoRepo := repo.ToDo{}

	err := todoRepo.Delete(mock.DB(), todoID)
	assert.NoError(t, err, "Queries were not called")

	err = mock.VerifyExpectations()
	assert.NoError(t, err, "Queries were not called")
}

func TestToDoDeleteFailure(t *testing.T) {
	mock := utils.GenerateMock()
	todoID := 1

	mock.ExpectExecFails(repo.DeleteQuery, errors.New("DB operation failed"))

	todoRepo := repo.ToDo{}
	err := todoRepo.Delete(mock.DB(), todoID)
	assert.Equal(t, err.Error(), "DB operation failed")

	err = mock.VerifyExpectations()
	assert.NoError(t, err, "Queries were not called")
}
func TestSuccessfulInsertToDo(t *testing.T) {
	todoItem := "Shopping"
	mock := utils.GenerateMock()
	mock.ExpectInsertToDoItem(1, todoItem)
	todoRepo := repo.ToDo{}

	id, err := todoRepo.Insert(todoItem, mock.DB())

	assert.NoError(t, err, "Unexpected Error thrown while trying to insert a new todo")
	assert.Equal(t, "1", id, "Unexpected Error thrown while trying to insert a new todo")
	err = mock.VerifyExpectations()
	assert.NoError(t, err, "Queries were not called")
}

func TestInsertToDoFails(t *testing.T) {
	mock := utils.GenerateMock()
	newToDoItem := "Flour"

	mock.ExpectQueryFails(repo.InsertQuery, errors.New("bombed"))

	todoRepo := repo.ToDo{}

	_,err := todoRepo.Insert(newToDoItem, mock.DB())

	assert.Equal(t, "bombed", err.Error(), "Unexpected Error thrown while trying to insert a new todo")
	err = mock.VerifyExpectations()
	assert.NoError(t, err, "Queries were not called")
}

func TestSuccessfulSelectToDos(t *testing.T) {
	mock := utils.GenerateMock()

	expectedRows := [][]driver.Value{
		{"1", "item1"},
		{"2", "item2"},
	}
	mock.ExpectSelect(expectedRows)

	todoRepo := repo.ToDo{}
	todoList, err := todoRepo.Select(mock.DB())

	assert.NoError(t, err, "Queries were not called")
	assert.Equal(t, 2, len(todoList), "Should have fetched two todo items")
	err = mock.VerifyExpectations()
	assert.NoError(t, err, "Queries were not called")
}

func TestSelectToDosFailure(t *testing.T) {
	mock := utils.GenerateMock()

	mock.ExpectSelectFails(errors.New("bombed"))

	todoRepo := repo.ToDo{}
	todoList, err := todoRepo.Select(mock.DB())

	assert.NotNil(t, err, "Queries were not called")
	assert.Nil(t, todoList, "Returned ToDo list should have been empty")

	err = mock.VerifyExpectations()
	assert.NoError(t, err, "Queries were not called")
}
