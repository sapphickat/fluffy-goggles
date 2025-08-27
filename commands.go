package main

import (
	"fmt"
	"time"
	"strconv"
)

func handleCommands(args []string, dataF dataFile) (dataFile, error) {
	if len(args) < 2 {
		return dataFile{}, fmt.Errorf("no command specified")
	}

	var err error

	switch args[1] {
	case "add":
		dataF, err = add(args, dataF)
	case "list":
		err = list(args, dataF.Tasks)
	case "update":
		dataF, err = update(args, dataF)
	case "delete":
		dataF, err = delete(args, dataF)
	case "mark-in-progress":
		dataF, err = markStatus(args, dataF, inProgress)
	case "mark-done":
		dataF, err = markStatus(args, dataF, done)
	case "mark-todo":
		dataF, err = markStatus(args, dataF, todo)
	}

	return dataF, err
}

func add(args []string, dataF dataFile) (dataFile, error) {
	if len(args) < 3  || len(args) > 4 {
		return dataFile{}, fmt.Errorf(`error! usage: add "[task name]"`)
	}

	newTask := task{
		Id: dataF.LastId + 1,
		Description: args[2],
		TaskStatus: todo,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	dataF.LastId++
	dataF.Tasks = append(dataF.Tasks, newTask)

	fmt.Printf("Task created sucessfully (Id: %v)\n", newTask.Id)

	return dataF, nil 
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
				v.UpdatedAt.Format("2006-01-02 15:04"))
		}
	} 
}

func update(args []string, dataF dataFile) (dataFile, error){
	if len(args) != 4 {
		return dataFile{}, fmt.Errorf("error! usage: update [id] [new description]")
	}

	id, err := strconv.Atoi(args[2])
	if err != nil {
		return dataFile{}, fmt.Errorf("error! invalid id")
	}

	for i, v := range dataF.Tasks {
		if v.Id == id {
			dataF.Tasks[i].Description = args[3] 
			dataF.Tasks[i].UpdatedAt = time.Now()
			fmt.Println("updated successfully")
			return dataF, nil
		}
	}

	return dataFile{}, fmt.Errorf("error! id not found") 
}


func delete(args []string, dataF dataFile) (dataFile, error) {
	if len(args) != 3 {
		return dataFile{}, fmt.Errorf("error! usage: delete [id]")
	}

	id, err := strconv.Atoi(args[2])
	if err != nil {
		return dataFile{}, fmt.Errorf("error, bad id")
	}

	for i, v := range dataF.Tasks {
		if v.Id == id {
			dataF.Tasks = append(dataF.Tasks[:i], dataF.Tasks[i+1:]...)
			println("deleted successfully")
			return dataF, nil
		}
	}
	return dataFile{}, fmt.Errorf("id not found")
}

func markStatus(args []string, dataF dataFile, s status) (dataFile, error) {
		if len(args) != 3 {
		return dataFile{}, fmt.Errorf("error! usage: mark-[todo|in-progress|done] [id]")
	}

	id, err := strconv.Atoi(args[2])
	if err != nil {
		return dataFile{}, fmt.Errorf("error, bad id")
	}

	for i, v := range dataF.Tasks {
		if v.Id == id {
			dataF.Tasks[i].TaskStatus = s
			println("status changed successfully")
			return dataF, nil
		}
	}
	return dataFile{}, fmt.Errorf("id not found")
}
