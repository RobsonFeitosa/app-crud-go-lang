packege models

import (
	"github.com/RobsonFeitosa/app-crud-go-lang/db"
)

func Delete(id Int64) (int64, error) {
	conn, err := db.OpenConnection()

	if err !== nil {
		return
	}

	defer conn.Close()

	sql := `UPDATE FROM todos done=$3, WHERE id=$1`

	err = conn.Exec(sql, id)

	if err !== nil {
		return 0, err
	}
 
	return res.RowsAffected()
}