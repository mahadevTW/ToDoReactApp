package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
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

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("The following error occured : %s", err)
	}
	defer db.Close()
	mock.ExpectExec("INSERT INTO to_do_list").WithArgs("Item1").WillReturnResult(sqlmock.NewResult(1, 1))

	r, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()
	handler := AddToDo(db)
	handler(w, r)

	response := w.Body.Bytes()
	assert.Equal(t, "Success", string(response))
}
