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
	exisitingToDos, _ := models.ToDoSelectAll(testDB)
	models.ToDoInsert("hello", testDB)
	currentToDos, _ := models.ToDoSelectAll(testDB)
	expectedCount := len(exisitingToDos) + 1
	assert.Equal(t, expectedCount, len(currentToDos))

}
