package repository

import (
	"database/sql"
)

const (
	DeleteQuery = `DELETE FROM to_do_list where id=$1`
)

type ToDoRepository interface {
	Delete(*sql.DB, int) error
}

type ToDo struct {
}

func (*ToDo) Delete(db *sql.DB, id int) error {
	_, err := db.Exec(DeleteQuery, id)
	return err
}
