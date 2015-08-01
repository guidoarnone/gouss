package parser
/*
 * ------ Imported packages ------
 */

import (
	"os"
	"errors"
)

/*
 * ------ Main Variables, regarding OS and I/O interaction. ------
 */ 

var ArgumentParsingError = errors.New("Invalid arguments. See gouss -h for help.")
const FileModeArgument = "-f"
const HelpModeArgument = "-h"

/*
 * ------ Types and Methods ------
 */


/*
 * FileConfig: Handles the file I/O configuration
 */

type FileConfig struct {
	FileModeOn bool 
	InputPath, OutputPath string
}

/*
 * If the FileConfig has no input path, its value is set to path.
 * Otherwise, the value is either stored as the ouput path, if it's 
 * not set, or an error is returned. 
 */

func (f *FileConfig) PushFilePath(path string) (err error) {
	switch {
		case f.InputPath == "":
			 f.InputPath = path
		case f.OutputPath == "":
			 f.OutputPath = path
		default:
			err = errors.New("Both input and output path are set.")
	}
	return err
}

/*
 * ------ Functions ------
 */

/*
 * Parsing CLI input to determine the behaviour of Gouss regarding I/O.
 */

func ParseArguments() (helpMode bool, fileConf FileConfig, parseError error) {
	
	//Initialization, by convention the fist argument corresponds to the executable name.
	helpMode = false
	fileConf.FileModeOn = false
	args := os.Args[1:]
	errorChannel := make(chan error)

	//Parse arguments concurrently	
	for _, argument := range args {
		go checkArgument(&helpMode, &fileConf, argument, errorChannel)
	}

	//Check whether any of the arguments generates an error.
	var err error
	for i := 0; i < len(args); i++ {
		err = <- errorChannel
	}
	if err != nil {
		parseError = ArgumentParsingError
	}

	return helpMode, fileConf, parseError
}

func checkArgument(helpMode *bool, fileConf *FileConfig, argument string, errCh chan error) {
	var err error
	
	switch argument {
		case HelpModeArgument:
			*helpMode = true
		case FileModeArgument:
			fileConf.FileModeOn = true
		default:
			//Argument is presumably a file name.
			err = fileConf.PushFilePath(argument)
	}

	errCh <- err
}