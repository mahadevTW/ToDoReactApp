package testutils

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
)

type Mock struct {
	mock sqlmock.Sqlmock
	db   *sql.DB
}

func GenerateMock() *Mock {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	return &Mock{mock: mock, db: db}
}
func (m *Mock) ExpectInsertToDoItem(value string) {
	m.mock.ExpectExec("INSERT INTO to_do_list").WithArgs(value).WillReturnResult(sqlmock.NewResult(1, 1))
}
func (m *Mock) VerifyExpectations() error {
	error := m.mock.ExpectationsWereMet()
	if error != nil {
		return error
	}
	m.db.Close()
	return nil
}

func (m *Mock) DB() *sql.DB {
	return m.db
}
