package main

import (
	"os"
)

func main() {

	// load file
	dataF, err := getDataFile()
	criticalErrorCheck(err , 1)	

	// logic
	dataF, err = handleCommands(os.Args, dataF)
	criticalErrorCheck(err , 1)	

	// save 
	err = saveTasks(dataF)
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