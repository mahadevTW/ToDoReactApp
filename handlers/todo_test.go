package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	utils "git.todo-app.com/ToDoReactApp/testutils"
	"github.com/stretchr/testify/assert"
	"database/sql/driver"
	"encoding/json"
	"git.todo-app.com/ToDoReactApp/models"
)

func TestSuccessfulResponseFromStatusCheckHandler(t *testing.T) {
	r, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()
	MeAliveMethod(w, r)
	response := w.Body.Bytes()
	assert.Equal(t, "Alive", string(response))
}

func TestInsertToDoElement(t *testing.T) {
	const requestData = `{"Item": "hello"}`

	mock := utils.GenerateMock()
	mock.ExpectInsertToDoItem("hello")

	r, _ := http.NewRequest("GET", "", bytes.NewBufferString(requestData))
	w := httptest.NewRecorder()
	handler := AddToDo(mock.DB())
	handler(w, r)

	response := w.Body.Bytes()
	assert.Equal(t, "Success", string(response))

	err := mock.VerifyExpectations()
	assert.NoError(t, err, "Queries were not called")
}

func TestSelectToDoElements(t *testing.T) {
	mock := utils.GenerateMock()

	expectedRows := [][]driver.Value{
		{"item1"},
		{"item2"},
	}
	mock.ExpectSelect(expectedRows)

	r, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()
	handler := SelectToDos(mock.DB())
	handler(w, r)

	response := w.Body.Bytes()
	var todos []*models.ToDo
	json.Unmarshal(response,&todos)

	assert.Equal(t, models.ToDo{Item:"item1"}, *todos[0])
	assert.Equal(t, models.ToDo{Item:"item2"}, *todos[1])

	err := mock.VerifyExpectations()
	assert.NoError(t, err, "Queries were not called")
}
