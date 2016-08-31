package testutils

import (
	"database/sql"
	"database/sql/driver"

	"github.com/DATA-DOG/go-sqlmock"
	"git.todo-app.com/ToDoReactApp/models"
	"regexp"
	"fmt"
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
	m.mock.ExpectExec(`INSERT INTO to_do_list`).WithArgs(value).WillReturnResult(sqlmock.NewResult(1, 1))
}

func (m *Mock) ExpectSelect(expectedRows [][]driver.Value) {
	columns := []string{"item"}
	rows := sqlmock.NewRows(columns)

	for _, value := range expectedRows {
		rows.AddRow(value...)
	}
	m.mock.ExpectQuery(sanitize(models.SelectQuery)).WillReturnRows(rows)
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
