package handlers

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

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
	fakeTodoRepository := &utils.MockToDoRepository{}
	fakeTodoRepository.On("Insert", "hello", mock.DB()).Return(nil)

	r, _ := http.NewRequest("GET", "", bytes.NewBufferString(requestData))
	w := httptest.NewRecorder()
	handler := AddToDo(mock.DB(), fakeTodoRepository)
	handler(w, r)

	response := w.Body.Bytes()
	assert.Equal(t, "Success", string(response))

	fakeTodoRepository.AssertExpectations(t)
}

func TestSelectToDoElements(t *testing.T) {
	mock := utils.GenerateMock()

	item1 := models.ToDo{Id:1, Item:"item1"}
	item2 := models.ToDo{Id:2, Item:"item2"}

	expectedItems := []models.ToDo{item1,item2}

	r, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()

	fakeTodoRepository := &utils.MockToDoRepository{}
	fakeTodoRepository.On("Select", mock.DB()).Return(expectedItems ,nil)

	handler := SelectToDos(mock.DB(), fakeTodoRepository)
	handler(w, r)

	response := w.Body.Bytes()
	var todos []*models.ToDo
	json.Unmarshal(response, &todos)

	assert.Equal(t, models.ToDo{Item: "item1", Id: 1}, *todos[0])
	assert.Equal(t, models.ToDo{Item: "item2", Id: 2}, *todos[1])

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
