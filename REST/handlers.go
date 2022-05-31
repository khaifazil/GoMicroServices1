package main

import (
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

	results, err := db.Query("SELECT * FROM courses")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var course courseInfo
		err := results.Scan(&course.Id, &course.Title)
		if err != nil {
			panic(err.Error())
		}

		courses[course.Title] = course
	}
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

	params := mux.Vars(r)

	results, err := db.Query("SELECT * FROM courses")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var course courseInfo
		err := results.Scan(&course.Id, &course.Title)
		if err != nil {
			panic(err.Error())
		}

		courses[course.Title] = course
	}

	if r.Method == "GET" {

		if _, ok := courses[params["courseTitle"]]; ok {
			json.NewEncoder(w).Encode(courses[params["courseTitle"]])
		} else {
			http.Error(w, "404 - No course found", http.StatusNotFound)
		}
	}

	if r.Method == "DELETE" {
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

	if r.Header.Get("Content-type") == "application/json" {

		//POST is for creating new course
		if r.Method == "POST" {

			//read the string sent to the service
			var newCourse courseInfo
			reqBody, err := ioutil.ReadAll(r.Body)
			defer r.Body.Close()

			if err == nil {
				//convert JSON to object
				json.Unmarshal(reqBody, &newCourse)

				if newCourse.Title == "" {
					http.Error(w, "422 - Please supply course information in JSON format", http.StatusNotFound)
					return
				}

				//check if course exists; add only if course does not exists
				if _, ok := courses[params["courseTitle"]]; !ok {
					courses[params["courseTitle"]] = newCourse
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - Course added: " + params["courseTitle"]))
				} else {
					http.Error(w, "409 - Duplicate course ID", http.StatusConflict)
				}
			} else {
				http.Error(w, "422 - Please supply course information in JSON format", http.StatusUnprocessableEntity)
			}
		} // end of post for new course

		//---PUT is for creating or updating existing course---
		if r.Method == "PUT" {
			var newCourse courseInfo
			reqBody, err := ioutil.ReadAll(r.Body)
			if err == nil {
				json.Unmarshal(reqBody, &newCourse)
				if newCourse.Title == "" {
					http.Error(w, "422 - Please supply course information in JSON format", http.StatusUnprocessableEntity)
					return
				}

				// check if course exists; add only if course does not exist
				if _, ok := courses[params["courseTitle"]]; !ok {
					courses[params["courseTitle"]] = newCourse
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - Course added: " + params["courseTitle"]))
				} else {
					// update course
					courses[params["courseTitle"]] = newCourse
					w.WriteHeader(http.StatusAccepted)
					w.Write([]byte("202 - Course updated: " + params["courseTitle"]))
				}
			} else {
				http.Error(w, "422 - Please supply course information in JSON format", http.StatusUnprocessableEntity)
			}
		} // end of put for update
	}
}
