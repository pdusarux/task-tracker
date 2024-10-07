package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Task struct {
	ID          uint
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

const (
	noTasksFoundMessage = "No tasks found"
	taskFormatString    = "ID: %d, Description: %s, Status: %s\n"
)

func main() {
	var tasks []Task

	for {
		fmt.Println("--------------------------------")
		fmt.Println("Welcome to the Task Manager")
		fmt.Println("1. Add Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Update Task")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")

		fmt.Println("Enter your choice: ")
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			addTask(&tasks)
		case 2:
			viewTasks(&tasks)
		case 3:
			updateTask(&tasks)
		case 4:
			deleteTask(&tasks)
		case 5:
			fmt.Println("Exit. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice")
		}

	}

}

func addTask(tasks *[]Task) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter task description: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)

	status := selectStatus("add")

	newTask := Task{
		ID:          uint(len(*tasks) + 1),
		Description: description,
		Status:      status,
	}

	*tasks = append(*tasks, newTask)
	fmt.Println("Task added successfully!")
}

func viewTasks(tasks *[]Task) {
	if len(*tasks) == 0 {
		fmt.Println(noTasksFoundMessage)
		return
	}

	fmt.Println("View tasks by:")
	fmt.Println("1. All tasks")
	fmt.Println("2. Todo")
	fmt.Println("3. In-progress")
	fmt.Println("4. Done")
	fmt.Print("Choose an option (1-4): ")

	var choice int
	fmt.Scan(&choice)

	fmt.Println("Tasks:")
	switch choice {
	case 1:
		for _, task := range *tasks {
			fmt.Printf(taskFormatString, task.ID, task.Description, task.Status)
		}
	case 2, 3, 4:
		status := []string{"todo", "in-progress", "done"}[choice-2]
		for _, task := range *tasks {
			if task.Status == status {
				fmt.Printf(taskFormatString, task.ID, task.Description, task.Status)
			}
		}
	default:
		fmt.Println("Invalid choice. Showing all tasks.")
		for _, task := range *tasks {
			fmt.Printf(taskFormatString, task.ID, task.Description, task.Status)
		}
	}
}

func updateTask(tasks *[]Task) {
	if len(*tasks) == 0 {
		fmt.Println(noTasksFoundMessage)
		return
	}
	fmt.Println("Enter the ID of the task to update: ")
	var id int
	fmt.Scan(&id)
	for _, task := range *tasks {
		if task.ID == uint(id) {
			fmt.Println("Enter new task description: ")
			var description string
			fmt.Scan(&description)
			status := selectStatus("update")
			task.Description = description
			if status != "" {
				task.Status = status
			}
			task.UpdatedAt = time.Now()
			(*tasks)[id-1] = task
			fmt.Println("Task updated successfully!")
			return
		}
	}
}

func deleteTask(tasks *[]Task) {
	if len(*tasks) == 0 {
		fmt.Println(noTasksFoundMessage)
		return
	}
	fmt.Println("Enter the ID of the task to delete: ")
	var id int
	fmt.Scan(&id)
	*tasks = append((*tasks)[:id-1], (*tasks)[id:]...)
	fmt.Println("Task deleted successfully!")
}

func selectStatus(stage string) string {
	fmt.Println("Enter new task status:")
	fmt.Println("1. Todo")
	fmt.Println("2. In-progress")
	fmt.Println("3. Done")
	fmt.Print("Choose a status (1-3): ")
	var statusChoice int
	fmt.Scan(&statusChoice)
	var status string

	switch statusChoice {
	case 1:
		status = "todo"
	case 2:
		status = "in-progress"
	case 3:
		status = "done"
	default:
		if stage == "add" {
			fmt.Println("Invalid choice. Setting status to 'todo'.")
			status = "todo"
		} else {
			fmt.Println("Invalid choice.")
			return ""
		}
	}
	return status
}
