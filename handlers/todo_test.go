package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	utils "git.todo-app.com/ToDoReactApp/testutils"
	"github.com/stretchr/testify/assert"
)

func TestSuccessfulResponseFromStatusCheckHandler(t *testing.T) {
	r, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()
	MeAliveMethod(w, r)
	response := w.Body.Bytes()
	assert.Equal(t, "Alive", string(response))
}

func TestInsertToDoElement(t *testing.T) {
	mock := utils.GenerateMock()
	mock.ExpectInsertToDoItem("hello")

	r, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()
	handler := AddToDo("hello", mock.DB())
	handler(w, r)

	response := w.Body.Bytes()
	assert.Equal(t, "Success", string(response))

	err := mock.VerifyExpectations()
	assert.NoError(t, err, "Queries were not called")
}
