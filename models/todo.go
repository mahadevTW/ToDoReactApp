package models

import (
	"database/sql"
	"fmt"

)

type ToDo struct {
	Item string `json:"Item"`
}



const(
	InsertQuery=`INSERT INTO to_do_list VALUES (nextval('todo_sequence'),$1)`
	SelectQuery = `SELECT text FROM to_do_list`
)

func ToDoInsert(value string, db *sql.DB) {
	query := InsertQuery
	_,err := db.Exec(query, value)
	fmt.Print("ERR = %#V",err)
}

func ToDoSelectAll(db *sql.DB) []ToDo {


	var resultSet []ToDo
	rows, err := db.Query(SelectQuery)

	if err != nil {
		fmt.Print("ERROR>>%v<<",err)
	}

	defer rows.Close()
	var todoElement string
	for rows.Next(){
		err := rows.Scan(&todoElement)
		if err != nil{
			fmt.Print("Skipping this row")
		}
		resultSet = append(resultSet, ToDo{
			Item: todoElement,
		})
	}
	return resultSet
}
