package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func saveTasks(tasks []string) {
	data := strings.Join(tasks, "\n")
	os.WriteFile("tasks.txt", []byte(data), 0644)
}

func loadTasks() []string {
	data, err := os.ReadFile("tasks.txt")
	if err != nil {
		return []string{}
	}
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}

func main() {
	tasks := loadTasks()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Choose an option: add, list, exit")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		switch input {
		case "add":
			fmt.Println("Enter a task:")
			scanner.Scan()
			task := strings.TrimSpace(scanner.Text())
			tasks = append(tasks, task)
			saveTasks(tasks)
			fmt.Println("Task added.")
		case "list":
			fmt.Println("Tasks:")
			for i, task := range tasks {
				fmt.Printf("%d: %s\n", i+1, task)
			}
		case "exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please choose add, list, or exit.")
		}
	}
}
