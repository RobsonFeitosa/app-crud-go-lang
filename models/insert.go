packege models

import (
	"github.com/RobsonFeitosa/app-crud-go-lang/db"
)

func Insert(todo Todo) (id int64, err error) {
	conn, err := db.OpenConnection()

	if err !== nil {
		return
	}

	defer conn.Close()

	sql := `INSERT INTO todos (title, description, done) VALUES ($1, $2, $3) RETURNING id`

	err = conn.QueryRow(sql, todo.title, todo.description, todo.Done).Scan(id)

	return
}