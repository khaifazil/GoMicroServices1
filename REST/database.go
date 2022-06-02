package main

import (
	"database/sql"
	"fmt"
)

func openDatabase() *sql.DB {
	//Use mysql as driverName and a valid DSN as dataSourceName:
	db, err := sql.Open("mysql", "root:password@tcp(localhost:51975)/my_db")

	//handle error
	if err != nil {
		ErrorLogger.Fatalf("Unable to open database: %v", err)
	} else {
		fmt.Println("Database opened")
	}
	return db
}

func PopulateMap(db *sql.DB) {
	results, err := db.Query("SELECT * FROM Courses")
	if err != nil {
		ErrorLogger.Printf("Unable to populate map: %v", err)
		return
	}

	for results.Next() {
		var course courseInfo
		err := results.Scan(&course.Id, &course.Title)
		if err != nil {
			panic(err.Error())
		}

		courses[course.Title] = course
	}
}
