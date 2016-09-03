package db_test

import (
	utils "git.todo-app.com/ToDoReactApp/testutils"
	"testing"
	"git.todo-app.com/ToDoReactApp/models"
	"github.com/stretchr/testify/assert"
)

func TestInsertToDo(t *testing.T) {

	testDB := utils.OpenTestDB()
	defer testDB.Close()
	exisitingToDos := models.ToDoSelectAll(testDB)
	models.ToDoInsert("hello", testDB)
	currentToDos := models.ToDoSelectAll(testDB)
	expectedCount := len(exisitingToDos) + 1
	assert.Equal(t,expectedCount,len(currentToDos))

}
