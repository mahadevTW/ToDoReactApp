package db_test

import (
	"testing"

	"git.todo-app.com/ToDoReactApp/models"
	utils "git.todo-app.com/ToDoReactApp/testutils"
	"github.com/stretchr/testify/assert"
)

func TestInsertToDo(t *testing.T) {

	testDB := utils.OpenTestDB()
	defer testDB.Close()
	exisitingToDos, err := models.ToDoSelectAll(testDB)
	assert.Nil(t, err, "Did not expect error while fetching todos")
	models.ToDoInsert("hello", testDB)
	currentToDos, err := models.ToDoSelectAll(testDB)
	assert.Nil(t, err, "Did not expect error while fetching todos")
	expectedCount := len(exisitingToDos) + 1
	assert.Equal(t, expectedCount, len(currentToDos))

}
