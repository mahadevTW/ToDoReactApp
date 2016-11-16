package testutils

import (
	"database/sql"
	"database/sql/driver"

	"github.com/DATA-DOG/go-sqlmock"

	repo "git.todo-app.com/ToDoReactApp/repository"

	"fmt"
	"regexp"
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
func (m *Mock) ExpectInsertToDoItem(todoId int, todoItem string) {
	columns := []string{"id"}
	rows := sqlmock.NewRows(columns)

	rows.AddRow(todoId)

	m.mock.ExpectQuery(sanitize(repo.InsertQuery)).WithArgs(todoItem).WillReturnRows(rows)
}

func (m *Mock) ExpectSelect(expectedRows [][]driver.Value) {
	columns := []string{"id", "item"}
	rows := sqlmock.NewRows(columns)

	for _, value := range expectedRows {
		rows.AddRow(value...)
	}
	m.mock.ExpectQuery(sanitize(repo.SelectQuery)).WillReturnRows(rows)
}
func (m *Mock) ExpectSelectFails(err error) {
	m.mock.ExpectQuery(sanitize(repo.SelectQuery)).WillReturnError(err)
}
func (m *Mock) ExpectExecFails(query string, err error) {
	m.mock.ExpectExec(sanitize(query)).WillReturnError(err)
}
func (m *Mock) ExpectQueryFails(query string, err error) {
	m.mock.ExpectQuery(sanitize(query)).WillReturnError(err)
}
func (m *Mock) ExpectDeleteSuccess(id int) {
	m.mock.ExpectExec(sanitize(repo.DeleteQuery)).WillReturnResult(sqlmock.NewResult(1, 1))
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

func sanitize(query string) string {
	r, err := regexp.Compile("[\n]+")
	if err != nil {
		fmt.Print("problem")
	}
	query = r.ReplaceAllString(query, " ")
	return regexp.QuoteMeta(query)
}
