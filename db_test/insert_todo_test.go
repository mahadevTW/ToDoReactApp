package db_test

import (
	"testing"

	repo "git.todo-app.com/ToDoReactApp/repository"
	utils "git.todo-app.com/ToDoReactApp/testutils"
	"github.com/stretchr/testify/assert"
)

func TestInsertToDo(t *testing.T) {

	testDB := utils.OpenTestDB()
	defer testDB.Close()
	todoRepo := repo.ToDo{}
	exisitingToDos, err := todoRepo.Select(testDB)
	assert.Nil(t, err, "Did not expect error while fetching todos")

	todoRepo.Insert("hello", testDB)

	currentToDos, err := todoRepo.Select(testDB)
	assert.Nil(t, err, "Did not expect error while fetching todos")
	expectedCount := len(exisitingToDos) + 1
	assert.Equal(t, expectedCount, len(currentToDos))

}
