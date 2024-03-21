package controller

import (
	"database/sql"
	"log"

	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() martini.Handler {
	db, err := sql.Open("mysql", "root:M4D3ENA@tcp(localhost:3306)/db_pbp_2")
	if err != nil {
		log.Fatal(err)
	}

	return func(c martini.Context) {
		c.Map(db)
		c.Next()
	}
}
