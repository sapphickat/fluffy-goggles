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
		// err = list(args, tasks)
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
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tasks = append(tasks, newTask)

	fmt.Printf("Task created sucessfully (Id: %v)\n", newTask.Id)

	return tasks, nil 
}

// func list(args []string, tasks []task) error {
// 	return 
// }

