package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Db_conect() *sql.DB {
	conexao := "user = postgres dbname = livraria_segabinassi password = 08012000 host = localhost sslmode = disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
