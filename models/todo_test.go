package models

import (
	"testing"

	utils "git.todo-app.com/ToDoReactApp/testutils"
	"github.com/stretchr/testify/assert"
)

func TestSuccessfulInsertToDo(t *testing.T) {
	todoItem := "Shopping"
	mock := utils.GenerateMock()
	mock.ExpectInsertToDoItem(todoItem)
	ToDoInsert(todoItem, mock.DB())
	err := mock.VerifyExpectations()
	assert.NoError(t, err, "Queries were not called")
}
