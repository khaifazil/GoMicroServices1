package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type CourseInfo struct {
	Id    int
	Title string
}

const baseURL = "http://localhost:5000/api/v1/courses"
const key = "?key=2c78afaf-97da-4816-bbee-9ad239abb296"

func getCourse(code string) {

	url := baseURL + "/" + code + key

	response, err := http.Get(url)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		var course CourseInfo
		data, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal(data, &course)

		fmt.Println("Status:", response.StatusCode)
		fmt.Println("ID:", course.Id)
		fmt.Println("Course Title:", course.Title)

		response.Body.Close()
	}
}

func getAllCourses() {
	response, err := http.Get(baseURL + key)
	var results map[string]interface{}
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {

		if data, err := ioutil.ReadAll(response.Body); err != nil {
			log.Fatal(err)
		} else {
			json.Unmarshal(data, &results)
			fmt.Println()
			fmt.Println("Status Code:", response.StatusCode)
			fmt.Println()
			for k, course := range results {
				fmt.Println("Course:", k)
				fmt.Println("ID:", course.(map[string]interface{})["Id"])
				fmt.Println("=========")
			}
		}
		response.Body.Close()
	}
}

func addCourse(code string) {

	response, err := http.Post(baseURL+"/"+code+key, "text/plain", nil)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println("Status Code:", response.StatusCode)
		fmt.Println(string(data))
		response.Body.Close()
	}
}

func updateCourse(code string, jsonData map[string]string) {
	jsonValue, _ := json.Marshal(jsonData)

	request, err := http.NewRequest(http.MethodPut, baseURL+"/"+code+key, bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println("Status Code:", response.StatusCode)
		fmt.Println(string(data))
		response.Body.Close()
	}
}

func deleteCourse(code string) {
	request, err := http.NewRequest(http.MethodDelete, baseURL+"/"+code+key, nil)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println("Status Code:", response.StatusCode)
		fmt.Println(string(data))
		response.Body.Close()
	}
}
