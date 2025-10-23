package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"taskmanager/model"
	"taskmanager/repository"
	"taskmanager/utils"
)

func main() {
	var repo repository.TaskRepository
	repo = repository.GetTaskRepository("json")

	if len(os.Args) < 2 {
		fmt.Println("No command received")
		os.Exit(1)
	}

	args := os.Args[1:]

	switch args[0] {
	case "add":
		handleAdd(repo, args)
	case "list":
		handleList(repo)
	case "done":
		handleDone(repo, args)
	case "delete":
		handleDelete(repo, args)
	default:
		fmt.Println("Command not reckognized")
	}
}

func handleAdd(repo repository.TaskRepository, args []string) {
	if len(args) < 3 {
		fmt.Println("Usage: add <title> <duedate>")
		os.Exit(1)
	}
	id := repo.GenerateId()
	title := args[1]
	status := model.Status(model.StatusPending)
	ddate, err := time.Parse(utils.DateLayout, args[2])
	if err != nil {
		fmt.Printf("Error parsing date: %v\n", err)
		os.Exit(1)
	}
	task := model.Task{
		Id:      id,
		Title:   title,
		Status:  status,
		DueDate: ddate,
	}
	repo.Save(task)
}

func handleList(repo repository.TaskRepository) {
	tasks, err := repo.Load("DueDate")
	if err != nil {
		fmt.Printf("Error listing tasks: %v\n", err)
		os.Exit(1)
	}
	for _, v := range tasks {
		fmt.Printf("%d. %s\t[%s]\t%s\n", v.Id, v.Title, string(v.Status), v.DueDate.Format(utils.DateLayout))
	}
}

func handleDone(repo repository.TaskRepository, args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: done <id>")
		os.Exit(1)
	}
	id, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("Error parsing id: %v\n", err)
		os.Exit(1)
	}
	err = repo.Update(id, "Status", model.StatusCompleted)
	if err != nil {
		fmt.Printf("Error updating field: %v\n", err)
		os.Exit(1)
	}
}

func handleDelete(repo repository.TaskRepository, args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: delete <id>")
		os.Exit(1)
	}
	id, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("Error parsing id: %v\n", err)
		os.Exit(1)
	}
	err = repo.Delete(id)
	if err != nil {
		fmt.Printf("Error deleting task: %v\n", err)
		os.Exit(1)
	}
}
