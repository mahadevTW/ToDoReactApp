package repository_test

import (
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
