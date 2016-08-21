package models

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestSuccessfulInsertToDo(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("The following error occured : %s", err)
	}
	defer db.Close()
	mock.ExpectExec("INSERT INTO to_do_list").WithArgs("Item1").WillReturnResult(sqlmock.NewResult(1, 1))
	
	ToDoInsert("Item1", db)
	
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "Queries were not called")
}
