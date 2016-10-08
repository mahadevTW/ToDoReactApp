package handlers

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"database/sql/driver"
	"encoding/json"

	"git.todo-app.com/ToDoReactApp/models"
	utils "git.todo-app.com/ToDoReactApp/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
		{1, "item1"},
		{2, "item2"},
	}
	mock.ExpectSelect(expectedRows)

	r, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()
	handler := SelectToDos(mock.DB())
	handler(w, r)

	response := w.Body.Bytes()
	var todos []*models.ToDo
	json.Unmarshal(response, &todos)

	assert.Equal(t, models.ToDo{Item: "item1"}, *todos[0])
	assert.Equal(t, models.ToDo{Item: "item2"}, *todos[1])

	err := mock.VerifyExpectations()
	assert.NoError(t, err, "Queries were not called")
}

func TestDeleteToDoSucess(t *testing.T) {
	const requestData = `{"Id": 1}`
	mock := utils.GenerateMock()
	fakeTodoRepository := &utils.MockToDoRepository{}
	fakeTodoRepository.On("Delete", mock.DB(), 1).Return(nil)

	handler := DeleteToDoHandler(mock.DB(), fakeTodoRepository)
	r, _ := http.NewRequest("DELETE", "", bytes.NewBufferString(requestData))
	w := httptest.NewRecorder()
	handler(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteToDoFailure(t *testing.T) {
	const requestData = `{"Id": 1}`
	mock := utils.GenerateMock()
	fakeTodoRepository := &utils.MockToDoRepository{}
	fakeTodoRepository.On("Delete", mock.DB(), 1).Return(errors.New("Something went wrong"))

	handler := DeleteToDoHandler(mock.DB(), fakeTodoRepository)
	r, err := http.NewRequest("DELETE", "", bytes.NewBufferString(requestData))
	require.NoError(t, err, "Expected no error")
	w := httptest.NewRecorder()
	handler(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
