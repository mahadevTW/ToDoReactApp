
.PHONY:build
build:
	 go build -o out/build/ToDoApp

.PHONY:db_migrate

db_migrate:
	goose up