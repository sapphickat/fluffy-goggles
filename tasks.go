package main

import (
	"fmt"
	"time"
	"os"
	"encoding/json"
	"io"
)

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
	TaskStatus	status		`json:"task_status"`	
	CreatedAt		time.Time `json:"created_at"` 
	UpdatedAt		time.Time `json:"updated_at"`
}

func getTasks() ([]task, error) {
	file, err := os.OpenFile("data.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, fmt.Errorf("error opening the file: %w", err)
	}

	defer file.Close()

	byte_data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading the file: %w", err)
	}

	var tasks []task
	
	if len(byte_data) != 0 { // maybe not the best solution, but works 
		err = json.Unmarshal(byte_data, &tasks)
		if err != nil {
			return nil, fmt.Errorf("error unmarshaling the data: %w", err)
		}
	}

	return tasks, nil
}

func saveTasks(tasks []task) error {
	file, err := os.OpenFile("data.json", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("error opening the file: %w", err)
	}
	defer file.Close()

	json_bytes, err := json.MarshalIndent(tasks, "", "  ") 
	if err != nil {
		return fmt.Errorf("error marshaling data: %w", err)
	}

	_, err = file.Write(json_bytes)
	if err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}

	return nil
}

func criticalErrorCheck(err error, exitCode int) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(exitCode)
	} 
}
