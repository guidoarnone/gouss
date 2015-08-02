package parser

/*
 * ------ Imported packages ------
 */

import (
 	"errors"
	"math/big"
)

/*
 * ------ Variables ------
 */ 

var RationalMatrixParseError = errors.New("An error ocurred parsing information to coherent numeric values. Please check your input data.")
var InvalidMatrixDimentionsError = errors.New("Your input does not correspond to a matrix with vaild dimentions.")
var InvalidElementSyntaxError = errors.New("One of your numbers has an invalid format. Please check your input data.")

/*
 * ------ Types and Methods ------
 */

type RationalMatrix [][]big.Rat

/*
 * Parses a rational matrix into a string, for user interaction.
 */

func (matrix *RationalMatrix) String() string {
	strMatrix := ""
	for i := 0; i < len(*matrix); i++ {
		columnLength := len((*matrix)[i])
		for j:= 0; j < columnLength; j++ {
			value := (*matrix)[i][j]
			//if denominator is one, it's removed from the string
			if big.NewInt(1).Cmp(value.Denom()) == 0 {
				strMatrix += value.Num().String()
			} else {
				strMatrix += value.String()
			}
			if j < columnLength -1 {
				strMatrix += " "
			}
		}
		strMatrix += "\n"
	}
	return strMatrix
}

/*
 * Parses a rational matrix into a byte array, for user interaction.
 */

func (matrix *RationalMatrix) Bytes() []byte {
	return []byte(matrix.String())
}

type MatrixElement struct {
	rowNumber, columnNumber int
	value big.Rat
}

type MatrixElementMessage struct {
	element MatrixElement
	err error
}


/*
 * ------ Functions ------
 */ 

/*
 * Returns how many rows and columns the matrix has. An error is returned in case of invalid size.
 */

func validateAndRetrieveMatrixDimentions(matrix [][]string) (rows, columns int, err error) {
	rows = len(matrix)
	if rows <= 0 {
		return 0, 0, InvalidMatrixDimentionsError
	}

	columns = -1
	for _, row := range matrix {
		if len(row) != columns && columns != -1 {
			err = InvalidMatrixDimentionsError
		}
		columns = len(row)
	}

	return rows, columns, err
}

/*
 * Parses and validates an input matrix of strings into a matrix of rational numbers.
 */

func ParseToRationalMatrix(stringMatrix [][]string) (matrix RationalMatrix, err error) {
	rowLength, columnLength, err := validateAndRetrieveMatrixDimentions(stringMatrix)
	if err != nil {
		return matrix, err
	}
	
	//Initialize the matrix and the channel
	ch := make(chan MatrixElementMessage)
	ok := make(chan bool)
	initMatrix(rowLength, columnLength, &matrix)
	go matrixSaver(rowLength, columnLength, &matrix, ch, ok)	

	//Concurrent parsing
	for i := 0; i < rowLength; i++ {
		for j := 0; j < columnLength; j++ {
			go createAndSaveElement(i, j, stringMatrix[i][j], ch)
		}
	}

	if result := <-ok; !result {
		err = RationalMatrixParseError
	}
	return matrix, err
}

/*
 * Initializes the rational numbers matrix 
 */

func initMatrix(rows, columns int, matrix *RationalMatrix) {
	for i := 1; i <= rows; i++ {
		row := make([]big.Rat, columns)
		*matrix = append(*matrix, row)
	}
}

/*
 * Saves incoming MatrixElements in a matrix.
 */

func matrixSaver(rows, columns int, matrix *RationalMatrix, ch chan MatrixElementMessage, ok chan bool) {
	var err error
	for i := 0; i < rows*columns && err == nil; i++ {
		message := <-ch
		element := message.element
		err = message.err
		(*matrix)[element.rowNumber][element.columnNumber] = element.value 
	}

	ok <- (err == nil)
}


/*
 * Creates a MatrixElement and saves it.
 */

func createAndSaveElement(row, column int, value string, ch chan MatrixElementMessage) {
	rationalValue, err := stringToRational(value)
	me := MatrixElement{rowNumber: row, columnNumber: column, value: rationalValue}
	ch <- MatrixElementMessage{element: me, err: err}
}

/*
 * Self-explanatory :).
 */

func stringToRational(num string) (rational big.Rat, err error) {
	_, ok := rational.SetString(num)
	if !ok {
		return rational, InvalidElementSyntaxError
	}
	return rational, nil
}