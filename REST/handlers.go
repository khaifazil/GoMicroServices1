package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {

	if !validKey(r) {
		http.Error(w, "401 - Invalid key", http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "Welcome to the REST API!")
}

func allCourses(w http.ResponseWriter, r *http.Request) {

	if !validKey(r) {
		http.Error(w, "401 - Invalid key", http.StatusUnauthorized)
		return
	}

	db := openDatabase()
	defer db.Close()
	defer fmt.Println("Database closed")

	PopulateMap(db)

	// returns all the courses in JSON
	json.NewEncoder(w).Encode(courses)
}

func course(w http.ResponseWriter, r *http.Request) {

	if !validKey(r) {
		http.Error(w, "401 - Invalid key", http.StatusUnauthorized)
		return
	}

	db := openDatabase()
	defer db.Close()
	defer fmt.Println("Database closed")

	PopulateMap(db)

	//create
	if r.Method == "POST" {
		CreateCourse(w, r, db)
		return
	}
	//retrieve
	if r.Method == "GET" {
		RetrieveCourse(w, r)
	}
	//update
	if r.Method == "PUT" {
		UpdateCourse(w, r, db)
		return
	}
	//delete
	if r.Method == "DELETE" {
		DeleteCourse(w, r, db)
	}
}

func CreateCourse(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var newCourse courseInfo
	params := mux.Vars(r)
	if _, ok := courses[params["courseTitle"]]; ok { //check for duplicate
		http.Error(w, "409 - Duplicate course ID", http.StatusConflict)
		return
	} else {
		query := fmt.Sprintf("INSERT INTO courses (title) VALUES ('%s')", params["courseTitle"])
		_, err := db.Query(query)
		if err != nil {
			panic(err.Error())
		}
		courses[params["courseTitle"]] = newCourse
		http.Error(w, "201 - Course added: "+params["courseTitle"], http.StatusCreated)
	}
}

func RetrieveCourse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if _, ok := courses[params["courseTitle"]]; ok {
		json.NewEncoder(w).Encode(courses[params["courseTitle"]])
	} else {
		http.Error(w, "404 - No course found", http.StatusNotFound)
	}
}

func UpdateCourse(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var newCourse courseInfo
	params := mux.Vars(r)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "422 - Please supply course information in JSON format", http.StatusUnprocessableEntity)
		return
	} else {
		json.Unmarshal(reqBody, &newCourse)
		if newCourse.Title == "" {
			http.Error(w, "422 - Please supply course information in JSON format", http.StatusUnprocessableEntity)
			return
		}
		// check if course exists; add only if course does not exist
		if k, ok := courses[params["courseTitle"]]; !ok {
			CreateCourse(w, r, db)
			return
		} else { // update course
			query := fmt.Sprintf("UPDATE courses SET title='%s' WHERE id=%d", newCourse.Title, k.Id)
			_, err := db.Query(query)
			if err != nil {
				panic(err.Error())
			}
			courses[newCourse.Title] = newCourse
			delete(courses, params["courseTitle"])
			http.Error(w, "201 - Course updated: "+params["courseTitle"], http.StatusAccepted)
		}
	}
}

func DeleteCourse(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	params := mux.Vars(r)
	if k, ok := courses[params["courseTitle"]]; ok {
		query := fmt.Sprintf("DELETE FROM courses WHERE title='%v'", k.Title)
		_, err := db.Query(query)
		if err != nil {
			panic(err.Error())
		}
		delete(courses, params["courseTitle"])
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("202 - Course deleted: " + k.Title))
	} else {
		http.Error(w, "404 - No course found", http.StatusNotFound)
	}
}
