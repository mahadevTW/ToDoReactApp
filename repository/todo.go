package repository

import (
	"database/sql"
	"fmt"

	"git.todo-app.com/ToDoReactApp/models"
)

const (
	DeleteQuery = `DELETE FROM to_do_list where id=$1`
	InsertQuery = `INSERT INTO to_do_list VALUES (nextval('todo_sequence'),$1) RETURNING id`
	SelectQuery = `SELECT id, text FROM to_do_list`
)

type ToDoRepository interface {
	Delete(*sql.DB, int) error
	Insert(string, *sql.DB) (string, error)
	Select(db *sql.DB) ([]models.ToDo, error)
}

type ToDo struct {
}

func (*ToDo) Delete(db *sql.DB, id int) error {
	_, err := db.Exec(DeleteQuery, id)
	return err
}

func (*ToDo) Insert(value string, db *sql.DB) (id string, err error) {
	err = db.QueryRow(InsertQuery,value).Scan(&id)
	return
}

func (*ToDo) Select(db *sql.DB) ([]models.ToDo, error) {

	var resultSet []models.ToDo
	rows, err := db.Query(SelectQuery)

	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	defer rows.Close()
	var todoElement string
	var todoId int
	for rows.Next() {
		err := rows.Scan(&todoId, &todoElement)
		if err != nil {
			fmt.Print("Skipping this row")
		}
		resultSet = append(resultSet, models.ToDo{
			Item: todoElement,
			Id:   todoId,
		})
	}
	return resultSet, nil
}
