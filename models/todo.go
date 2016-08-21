package models

import "database/sql"

func ToDoInsert(value string, db *sql.DB) {
	query := `INSERT INTO to_do_list VALUES ($1)`
	db.Exec(query, value)
}
