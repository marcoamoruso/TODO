package main

import (
	"TODO/Models"
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"slices"
	"strconv"
	"strings"

	"time"
)

func main() {

	// Set server variables
	serverPort := os.Getenv("SERVER_PORT")
	serverHost := os.Getenv("SERVER_HOST")
	serverUrl := "http://" + serverHost + ":" + serverPort

	// User input loop
	for {

		// Welcome message
		fmt.Print("\nWelcome to your TODO list! Please select an option:\n[1] Check your TODO list\n[2] Add a TODO\n[3] Delete a TODO\n[4] Modify a TODO\n[5] Checkmark/Uncheckmark a TODO\n[6] Quit\nOption: ")

		// Read choice from input
		choice := ""
		fmt.Scanf("%s", &choice)

		switch choice {

		// [1] Check your TODO list
		case "1":

			// GET request
			res := httpRequest(http.MethodGet, serverUrl+"/todo", nil)

			// Check if there are elements in TODO list
			if len(res["data"].([]interface{})) == 0 {
				println("No elements in TODO list.")
			} else {
				fmt.Println("TODO list:")
			}

			// Parse response data array
			_, _ = parseArray(res["data"].([]interface{}))

		// [2] Add a TODO
		case "2":

			// Read Title and Deadline from user input
			var title string
			var deadline string
			fmt.Print("Please type your TODO: ")
			title = readLine()
			fmt.Print("Please type the deadline (YYYY-MM-DD): ")
			deadline = readLine()

			// Create TODO element
			todo := Models.TODO{Title: title, Deadline: deadline, Done: false}

			// JSON encoding of data
			jsonData, err := json.Marshal(todo)
			if err != nil {
				log.Fatal(err)
			}

			// POST request
			res := httpRequest(http.MethodPost, serverUrl+"/todo", jsonData)

			// Show response message
			if res["message"].(string) == "Creation: Success!" {
				fmt.Println("TODO created: Title: " + title + "; Deadline: " + deadline + "; Completed: No")
			} else {
				fmt.Println(res["message"])
			}

		// [3] Delete a TODO
		case "3":

			// GET request
			res := httpRequest(http.MethodGet, serverUrl+"/todo", nil)

			// Check if there are elements in TODO list
			if len(res["data"].([]interface{})) == 0 {
				println("No elements in TODO list.")
				break
			} else {
				fmt.Println("Which one would you like to delete? Please select an option")
			}

			// Parse response data array
			idMap, _ := parseArray(res["data"].([]interface{}))

			// Read choice from input and check if choice is valid
			fmt.Print("Option: ")
			choiceInt := readInt(idMap, nil)

			// Get correct TODO Id based on user choice
			id := strconv.Itoa((idMap[choiceInt]))

			// DELETE request
			res = httpRequest(http.MethodDelete, serverUrl+"/todo/"+id, nil)

			// Show response message
			fmt.Println(res["message"])

		// [4] Modify a TODO
		case "4":

			// GET request
			res := httpRequest(http.MethodGet, serverUrl+"/todo", nil)

			// Check if there are elements in TODO list
			if len(res["data"].([]interface{})) == 0 {
				println("No elements in TODO list.")
				break
			} else {
				fmt.Println("Which one would you like to modify?")
			}

			// Parse response data array
			idMap, todoList := parseArray(res["data"].([]interface{}))

			// Read choice from input and check if choice is valid
			fmt.Print("Option: ")
			choiceInt := readInt(idMap, nil)

			// Ask user which TODO field he/she wants to modify
			fmt.Print("What would you like to modify?\n[1] Title\n[2] Deadline\n[3] Title and Deadline\nOption: ")

			// Read choice from input and check if choice is valid
			validOptions := []int{1, 2, 3}
			option := readInt(nil, validOptions)
			todo := Models.TODO{}
			switch option {

			// Modify Title
			case 1:
				fmt.Print("Please type a new title: ")
				todo.Title = readLine()
				// Set Deadline and Done equal to the current ones
				for i := 0; i < len(todoList); i++ {
					if todoList[i].Id == idMap[choiceInt] {
						todo.Deadline = todoList[i].Deadline
						todo.Done = todoList[i].Done
					}
				}

			// Modify Deadline
			case 2:
				fmt.Print("Please type a new deadline (YYYY-MM-DD): ")
				todo.Deadline = readLine()
				// Set Title and Done equal to the current ones
				for i := 0; i < len(todoList); i++ {
					if todoList[i].Id == idMap[choiceInt] {
						todo.Title = todoList[i].Title
						todo.Done = todoList[i].Done
					}
				}

			// Modify Title and Deadline
			case 3:
				fmt.Print("Please type a new title: ")
				todo.Title = readLine()
				fmt.Print("Please type a new deadline (YYYY-MM-DD): ")
				todo.Deadline = readLine()
				// Set Done equal to the current one
				for i := 0; i < len(todoList); i++ {
					if todoList[i].Id == idMap[choiceInt] {
						todo.Done = todoList[i].Done
					}
				}
			default:
				fmt.Println("Please select a valid option.")
			}

			// JSON encoding of data
			jsonData, err := json.Marshal(todo)
			if err != nil {
				log.Fatal(err)
			}

			// Get correct TODO Id based on user choice
			id := strconv.Itoa((idMap[choiceInt]))

			// PUT request
			res = httpRequest(http.MethodPut, serverUrl+"/todo/"+id, jsonData)

			// Show response message
			fmt.Println(res["message"])

		// [5] Checkmark/Uncheckmark a TODO
		case "5":

			// GET request to show TODO elements
			res := httpRequest(http.MethodGet, serverUrl+"/todo", nil)

			// Check if there are elements in TODO list
			if len(res["data"].([]interface{})) == 0 {
				println("No elements in TODO list.")
				break
			} else {
				fmt.Println("Which one would you like to checkmark?")
			}

			// Parse response data array
			idMap, todoList := parseArray(res["data"].([]interface{}))

			// Read choice from input and check if choice is valid
			fmt.Print("Option: ")
			choiceInt := readInt(idMap, nil)

			// Create TODO element
			todo := Models.TODO{}

			// Set Title and Deadline equal to the current ones. Set Done as the opposite (true/false)
			for i := 0; i < len(todoList); i++ {
				if todoList[i].Id == idMap[choiceInt] {
					todo.Title = todoList[i].Title
					todo.Deadline = todoList[i].Deadline
					todo.Done = !todoList[i].Done
				}
			}

			// JSON encoding of data
			jsonData, err := json.Marshal(todo)
			if err != nil {
				log.Fatal(err)
			}

			// Get correct TODO Id based on user choice
			id := strconv.Itoa((idMap[choiceInt]))

			// PUT request
			res = httpRequest(http.MethodPut, serverUrl+"/todo/"+id, jsonData)

			// Print response message
			fmt.Println(res["message"])

			// If the TODO was checkmarked, ask the user if he/she wants to delete it from the list
			if todo.Done {
				fmt.Print("Would you like to delete it from the list?\n[1] Yes\n[2] No\nOption: ")

				// Read choice from input and check if choice is valid
				validOptions := []int{1, 2}
				option := readInt(nil, validOptions)

				if option == 1 {
					// DELETE Request
					res := httpRequest(http.MethodDelete, serverUrl+"/todo/"+id, nil)
					// Show response message
					fmt.Println(res["message"])
				}
			}

		// [6] Quit
		case "6":
			return

		// No valid option as input
		default:
			fmt.Println("Please select a valid option.")
		}

		// Show welcome message with a slight delay
		time.Sleep(time.Second / 2)
	}

}

// Parse response data array
func parseArray(anArray []interface{}) (map[int]int, []Models.TODO) {
	var title string
	var deadline string
	var done bool
	idMap := make(map[int]int)
	count := 1
	todoList := []Models.TODO{}
	// Iterate over TODO list
	for i, val := range anArray {
		fmt.Printf("[%d] ", i+1)
		todo := Models.TODO{}
		// Iterate over TODO element fields
		for key, concreteVal := range val.(map[string]interface{}) {
			switch key {
			case "id":
				id := int(concreteVal.(float64))
				idMap[count] = id
				count += 1
				todo.Id = id
			case "done":
				done = concreteVal.(bool)
				todo.Done = done
			case "title":
				title = concreteVal.(string)
				todo.Title = title
			case "deadline":
				deadline = concreteVal.(string)[0:10]
				todo.Deadline = deadline
			default:
				continue
			}
		}
		// Show TODO element to user
		if done {
			fmt.Println("Title: " + title + "; Deadline: " + deadline + "; Completed: Yes")
		} else {
			fmt.Println("Title: " + title + "; Deadline: " + deadline + "; Completed: No")
		}
		todoList = append(todoList, todo)
	}
	return idMap, todoList
}

// Read int from input
func readInt(idMap map[int]int, validOptions []int) int {
	// User input loop
	for {
		// Read choice from input
		var choice string
		fmt.Scanf("%s", &choice)
		choiceInt, err := strconv.Atoi(choice)
		if err != nil {
			fmt.Print("Please select a valid option.\nOption: ")
			continue
		}
		// Determine valid options if no valid options were passed as argument
		if idMap != nil {
			validOptions = []int{}
			for k := range idMap {
				validOptions = append(validOptions, k)
			}
		}
		// Check if option is valid
		if idMap != nil || validOptions != nil {
			if slices.Contains(validOptions, choiceInt) {
				return choiceInt
			} else {
				fmt.Print("Please select a valid option.\nOption: ")
			}
		}
	}
}

// Read line from input
func readLine() string {
	// User input loop
	for {
		// Read line from input
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again.", err)
			continue
		}
		// Remove \n from string
		line = strings.TrimSuffix(line, "\n")
		// Check if user inputs empty string
		if len(line) == 0 {
			fmt.Println("An error occured while reading input. Please try again.")
			continue
		}
		return line
	}
}

// Make a http request
func httpRequest(method string, requestUrl string, jsonData []byte) map[string]interface{} {
	req, err := http.NewRequest(method, requestUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
	return res
}
