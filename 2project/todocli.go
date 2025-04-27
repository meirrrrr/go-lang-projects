package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Todo struct {
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

var todoList []Todo
var fileName = "tasks.json"

// Function to load tasks from the JSON file
func loadTasks() error {
	file, err := os.Open(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			// If the file doesn't exist, return no error
			return nil
		}
		return err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&todoList)
	if err != nil {
		return err
	}
	return nil
}

// Function to save tasks to the JSON file
func saveTasks() error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(todoList)
	if err != nil {
		return err
	}
	return nil
}

// Function to add a new task
func addTask(description string) {
	todoList = append(todoList, Todo{Description: description, Done: false})
	saveTasks()
	fmt.Println("Task added:", description)
}

// Function to list all tasks
func listTasks() {
	if len(todoList) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	fmt.Println("Todo List:")
	for i, task := range todoList {
		doneStatus := "Not Done"
		if task.Done {
			doneStatus = "Done"
		}
		fmt.Printf("%d. %s [%s]\n", i+1, task.Description, doneStatus)
	}
}

// Function to mark a task as done
func markTaskAsDone(taskNumber int) {
	if taskNumber < 1 || taskNumber > len(todoList) {
		fmt.Println("Invalid task number.")
		return
	}

	todoList[taskNumber-1].Done = true
	saveTasks()
	fmt.Printf("Task %d marked as done: %s\n", taskNumber, todoList[taskNumber-1].Description)
}

func main() {
	// Load existing tasks from the JSON file
	err := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	// Command handling loop using bufio.NewReader to handle full input including spaces
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter command (add, list, done <task_number>, quit): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Process different commands
		switch {
		case input == "quit":
			fmt.Println("Exiting Todo app.")
			return
		case strings.HasPrefix(input, "add "):
			// Add task
			description := input[4:] // The part after "add "
			if description == "" {
				fmt.Println("Please provide a description for the task.")
				continue
			}
			addTask(description)
		case input == "list":
			// List tasks
			listTasks()
		case strings.HasPrefix(input, "done "):
			// Mark task as done
			var taskNumber int
			_, err := fmt.Sscanf(input, "done %d", &taskNumber)
			if err != nil {
				fmt.Println("Invalid command format. Use: done <task_number>")
			} else {
				markTaskAsDone(taskNumber)
			}
		default:
			fmt.Println("Unknown command. Available commands: add, list, done <task_number>, quit.")
		}
	}
}
