package main

import (
	"fmt"
	"time"
)

func handleCommands(args []string, tasks []task) ([]task, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("no command specified")
	}

	var err error

	switch args[1] {
	case "add":
		tasks, err = add(args, tasks)
	case "list":
		err = list(args, tasks)
	case "update":
	case "delete":
	case "mark-in-progress":
	case "mark-done":
	}

	return tasks, err
}

func add(args []string, tasks []task) ([]task, error) {
	if len(args) < 3  || len(args) > 4 {
		return nil, fmt.Errorf(`error! usage: add "[task name]"`)
	}

	newTask := task{
		Id: len(tasks) + 1,
		Description: args[2],
		TaskStatus: todo,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tasks = append(tasks, newTask)

	fmt.Printf("Task created sucessfully (Id: %v)\n", newTask.Id)

	return tasks, nil 
}

func list(args []string, tasks []task) error {
	if len(args) < 2 || len(args) > 3 {
		return fmt.Errorf("error! usage: list [done|todo|in-progress]")
	}
	if len(tasks) < 1 {
		return fmt.Errorf("error! empty list, try 'add'")
	}

	filter := none
	if len(args) == 3 {
		switch args[2] {
		case "todo":
			filter = todo
		case "done":
			filter = done 
		case "in-progress":
			filter = inProgress 
		default:
		return fmt.Errorf("error! usage: list [done|todo|in-progress]")
		}
	}

	listHelper(filter, tasks)
	return nil  
}

func listHelper(filter status, tasks []task) {
	for _, v := range tasks {
		if filter == none || filter == v.TaskStatus {
			fmt.Printf("[%v] %q %v (created: %v updated: %v)\n",
				v.Id, v.Description, v.TaskStatus,
				v.CreatedAt.Format("2006-01-02 15:04"),
				v.CreatedAt.Format("2006-01-02 15:04"))
		}
	} 
}
