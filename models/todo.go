package models

import (
	"database/sql"
	"fmt"
)

type ToDo struct {
	Item string `json:"Item"`
}

const (
	InsertQuery = `INSERT INTO to_do_list VALUES (nextval('todo_sequence'),$1)`
	SelectQuery = `SELECT text FROM to_do_list`
	DeleteQuery = `DELETE FROM to_do_list where id=$1`
)

func ToDoInsert(value string, db *sql.DB) (err error) {
	_, err = db.Exec(InsertQuery, value)
	return
}

func ToDoSelectAll(db *sql.DB) ([]ToDo, error) {

	var resultSet []ToDo
	rows, err := db.Query(SelectQuery)

	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	defer rows.Close()
	var todoElement string
	for rows.Next() {
		err := rows.Scan(&todoElement)
		if err != nil {
			fmt.Print("Skipping this row")
		}
		resultSet = append(resultSet, ToDo{
			Item: todoElement,
		})
	}
	return resultSet, nil
}

func DeleteToDo(db *sql.DB, item int) error {
	_, err := db.Exec(DeleteQuery, item)
	return err
}
