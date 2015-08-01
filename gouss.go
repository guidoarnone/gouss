package main

import (
	"fmt"
	"os"
	"os/exec"
	"io"
	"io/ioutil"
	"bufio"
	"strings"
	"gouss/parser"
)

/*
 * ------ Variables ------
 */

const helpFilePath = "./.data/help"


/*
 * ------ Functions ------
 */

func main() {
	
	helpModeOn, fileConfiguration, err := parser.ParseArguments()
	exitProcessIfError(err, -1)

	if helpModeOn {
		showHelp()
		return
	} 

	input, output := resolveFileIO(fileConfiguration)
	
	stringMatrix := userInputToString(input)
	
	//test
	for index, _ := range stringMatrix {
		for _, subvalue := range stringMatrix[index] {
			output.Write([]byte(subvalue+" "))
		}
		output.Write([]byte("\n"))
	}

}

/*
 * Displays the help screen, reading its content from a local file.
 */

func showHelp() {
	err := clearScreen()
	exitProcessIfError(err, -1)

	helpBytes, err := ioutil.ReadFile(helpFilePath)
	exitProcessIfError(err, -1)

	help := string(helpBytes)
	fmt.Println(help)
	meh := "" //empty auxiliar
	fmt.Scanf("%s", &meh)
}

/*
 * Sets the Gouss input and output sources, depending on the file cofiguration.
 */

func resolveFileIO(fileConfiguration parser.FileConfig) (input io.Reader, output io.Writer) {
	
	//Default I/O is Stdin & Stdout
	input  = os.Stdin
	output = os.Stdout

	//If file mode is on, set up files as I/O
	if fileConfiguration.FileModeOn {
		fileInput, err := os.OpenFile(fileConfiguration.InputPath, os.O_RDONLY, os.ModePerm)
		exitProcessIfError(err, -1)
		input = fileInput
		//If output is not defined, it defaults to stdout
		if fileConfiguration.OutputPath != "" {
			fileOutput, err := os.OpenFile(fileConfiguration.OutputPath, os.O_WRONLY | os.O_CREATE, 0660)
			exitProcessIfError(err, -1)
			output = fileOutput
		}
	}

	return input, output
}

/*
 * Formats the user input, formed by several lines, as a rune matrix.
 */

func userInputToString(inputReader io.Reader) [][]string {
	
	//Read the input by lines
	var err error = nil
	var lines []string
	var line string
	
	reader := bufio.NewReader(inputReader)
	for err == nil {
		line, err = reader.ReadString(byte('\n'))
		switch {
			case line == "\n":
				err = io.EOF
			case  err == nil:
				//discards the delimiter (\n)
				line = line[: len(line) - 1]
				lines = append(lines, line)
		}
	}

	//Concurrently separate each line using a blank space as delimiter
	var matrix [][]string
	var ch = make(chan []string)

	for _, line := range lines {
		go splitLine(line, ch)
	}

	for i := 0; i < len(lines); i++ {
		matrix = append(matrix, <-ch)
	}

	return matrix
}

/*
 * Line splitter.
 */

func splitLine(l string, ch chan []string) {
	var split []string
	rawSplit := strings.Split(l, "")
	for _, s := range rawSplit {
		if s[0] != ' ' {
			split = append(split, s)
		}
	}
	ch <- split
}

/*
 * ------ OS Related Functions ------
 */

func clearScreen() (err error) {
	clearCmd := "clear" 
	command := exec.Command(clearCmd, "")
	command.Stdout = os.Stdout
	err = command.Run()
	return err
}

func exitProcessIfError(err error, errorCode int) {
	if err != nil {
		exitProcess(err, errorCode)
	}
}

func exitProcess(err error, errorCode int) {
	fmt.Println(err.Error())
	os.Exit(errorCode)
}
