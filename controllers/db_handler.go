package controllers

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB
var err error

func Connect() *sqlx.DB {
	db, err = sqlx.Connect("mysql", "root:@tcp(localhost:3306)/db_latihan_pbp?parseTime=true&loc=Asia%2FJakarta")

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return db
}