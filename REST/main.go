package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type courseInfo struct {
	Id    int
	Title string `json:"Title"`
}

// used for storing courses on the REST API
var courses map[string]courseInfo

func validKey(r *http.Request) bool {
	v := r.URL.Query()
	if key, ok := v["key"]; ok {
		if key[0] == "2c78afaf-97da-4816-bbee-9ad239abb296" {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func main() {
	// instantiate courses
	courses = make(map[string]courseInfo)

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/", home)
	router.HandleFunc("/api/v1/courses", allCourses)
	router.HandleFunc("/api/v1/courses/{courseTitle}", course).Methods("GET", "PUT", "POST", "DELETE")

	fmt.Println("Listening at port 5001")
	log.Fatal(http.ListenAndServe(":5001", router))
}
