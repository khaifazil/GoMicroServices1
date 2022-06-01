package main

import "fmt"

func userInputYN(question string) bool {
	userInput := ""
	for {
		fmt.Printf("%s (y/n): ", question)
		fmt.Scanln(&userInput)
		if userInput == "y" {
			return true
		} else if userInput == "n" {
			return false
		} else {
			fmt.Println("Invalid input. Please reply with 'y' or 'n'.")
		}
	}
}

func userRawStringInput(question string) string {

	fmt.Printf(question)
	userInput := ""
	fmt.Scanln(&userInput)
	return userInput
}

func backToMain() {
	userInput := ""
	fmt.Println("\nPress enter to go back to main menu...")
	fmt.Scanln(&userInput)
	main()
}
