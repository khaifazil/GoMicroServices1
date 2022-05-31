package main

import (
	"database/sql"
	"fmt"
)

func openDatabase() *sql.DB {
	//Use mysql as driverName and a valid DSN as dataSourceName:
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/goms1_db")

	//handle error
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Database opened")
	}
	return db
}
