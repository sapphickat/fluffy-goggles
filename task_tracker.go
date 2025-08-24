package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type status int	

const (
	todo status = iota
	inProgress
	done
)

var statusName = map[status]string{
	todo: "todo",
	inProgress: "in-progress",
	done: "done",
}

func (s status) String() string {
	return statusName[s]
}

type task struct {
	Id 					int				`json:"id"`
	Description string		`json:"description"`
	TaskStatus	status			
	CreatedAt		time.Time 
	UpdatedAt		time.Time
}

func main() {
	// open/create file
	file, err := os.OpenFile("data.json", os.O_RDWR|os.O_CREATE, 0666)
	check(err)
	defer file.Close()

	byte_data, err := io.ReadAll(file)
	var tasks []task
	json.Unmarshal(byte_data, &tasks)

	// fmt.Printf("%+v", tasks)

	// logic
	tasks = handleCommands(os.Args, tasks)

	// saving to json
	json_bytes, err := json.MarshalIndent(tasks, "", "  ") // no caso, vou dar marshall em uma slice de tasks
	check(err)
	err = file.Truncate(0)
	check(err)
	_, err = file.Seek(0, 0)
	check(err)
	_, err = file.Write(json_bytes)
	check(err)
}

func handleCommands(args []string, tasks []task) []task {

	switch args[1] {
	case "add":
		tasks = add(args, tasks)
	case "list":
	case "update":
	case "delete":
	case "mark-in-progress":
	case "mark-done":
	}
	return tasks
}

func add(args []string, tasks []task) []task {
	if (len(args) -1 != 2) {
		fmt.Println("Error! Usage: add [task name]")
		return tasks
	}

	newTask := task{
		Id: len(tasks) + 1,
		Description: args[2],
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tasks = append(tasks, newTask)

	fmt.Printf("Task created sucessfully (Id: %v)\n", newTask.Id)

	return tasks
}