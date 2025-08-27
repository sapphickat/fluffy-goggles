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
	none
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

type dataFile struct {
	LastId	int 		`json:"last_id"`	
	Tasks		[]task	`json:"tasks"` 
}

func getDataFile() (dataFile, error) {
	file, err := os.OpenFile("data.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return dataFile{}, fmt.Errorf("error opening the file: %w", err)
	}

	defer file.Close()

	byte_data, err := io.ReadAll(file)
	if err != nil {
		return dataFile{}, fmt.Errorf("error reading the file: %w", err)
	}

	var data dataFile
	
	if len(byte_data) != 0 { // maybe not the best solution, but works 
		err = json.Unmarshal(byte_data, &data)
		if err != nil {
			return dataFile{}, fmt.Errorf("error unmarshaling the data: %w", err)
		}
	}

	return data, nil
}

func saveTasks(dataF dataFile) error {
	file, err := os.OpenFile("data.json", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("error opening the file: %w", err)
	}
	defer file.Close()

	json_bytes, err := json.MarshalIndent(dataF, "", "  ") 
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
