package testutils

import (
	"database/sql"
	_ "github.com/lib/pq" //Postgresql driver
)

func OpenTestDB() *sql.DB {

	db, _:= sql.Open("postgres", "user=todo_user password=todo dbname=todo_app sslmode=disable")
	db.Exec("DROP DATABASE IF EXISTS todo_app_test");
	db.Exec("CREATE DATABASE todo_app_test WITH TEMPLATE todo_app OWNER todo_user_test")
	return db
}
