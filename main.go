package main

import (
	"os"
)

func main() {

	// load file
	tasks, err := getTasks()
	criticalErrorCheck(err , 1)	

	// logic
	tasks, err = handleCommands(os.Args, tasks)
	criticalErrorCheck(err , 1)	

	// save 
	err = saveTasks(tasks)
	criticalErrorCheck(err , 1)	

}
/* 
	# Refactor checklist
	- use struct field tags JSON x
	- use helper function to load/save on file x
	- use better erro handling x
	- use better ids 
	- improve savefile truncate x
*/