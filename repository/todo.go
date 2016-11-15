package repository

import (
	"database/sql"
)

const (
	DeleteQuery = `DELETE FROM to_do_list where id=$1`
	InsertQuery = `INSERT INTO to_do_list VALUES (nextval('todo_sequence'),$1)`

)

type ToDoRepository interface {
	Delete(*sql.DB, int) error
	Insert(string, *sql.DB) error
}

type ToDo struct {
}

func (*ToDo) Delete(db *sql.DB, id int) error {
	_, err := db.Exec(DeleteQuery, id)
	return err
}

func (*ToDo) Insert(value string, db *sql.DB) (err error) {
	_, err = db.Exec(InsertQuery, value)
	return
}
