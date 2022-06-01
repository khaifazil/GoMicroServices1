package main

import "fmt"

func main() {

	fmt.Println()
	fmt.Println("Courses Client App")
	fmt.Println("==================")
	fmt.Println("1. Create Course")
	fmt.Println("2. Retrieve All Courses")
	fmt.Println("3. Retrieve Course")
	fmt.Println("4. Update Course")
	fmt.Println("5. Delete Course")
	fmt.Println("6. Exit")
	fmt.Println()

	userInput := userRawStringInput("Pick a feature:")

	switch userInput {
	case "1":
		fmt.Println("\nCreate Course")
		fmt.Println("=============")
		userInput = userRawStringInput("Input course to create: ")
		fmt.Println()

		addCourse(userInput)

		fmt.Println()
		backToMain()
	case "2":
		fmt.Println("\nRetrieve All Courses")
		fmt.Println("====================")

		getAllCourses()

		fmt.Println()
		backToMain()
	case "3":
		fmt.Println("\nRetrieve Course")
		fmt.Println("===============")

		userInput = userRawStringInput("Input course to retrieve: ")
		fmt.Println()

		getCourse(userInput)

		fmt.Println()
		backToMain()
	case "4":
		fmt.Println("\nUpdate Course")
		fmt.Println("===============")

		courseToUpdate := userRawStringInput("Input course to update: ")
		toUpdate := userRawStringInput("Input update: ")

		var jsonData = map[string]string{"title": ""}
		jsonData["title"] = toUpdate
		updateCourse(courseToUpdate, jsonData)

		fmt.Println()
		backToMain()
	case "5":
		fmt.Println("\nDelete Course")
		fmt.Println("===============")

		fmt.Printf("Input course to delete: ")
		userInput = userRawStringInput("Input course to delete: ")
		fmt.Println()

		deleteCourse(userInput)

		fmt.Println()
		backToMain()
	case "6":
		if userInputYN("Are you sure you want to exit?") {
			fmt.Println("Goodbye!")
		} else {
			main()
		}
	default:
		fmt.Println("Not a valid input")
		main()
	}
}
