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

type MatrixElement struct {
	rowNumber, columnNumber int
	value big.Rat
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

func ParseToRationalMatrix(stringMatrix [][]string) (matrix [][]big.Rat, err error) {
	rowLength, columnLength, err := validateAndRetrieveMatrixDimentions(stringMatrix)
	if err != nil {
		return nil, err
	}

	//TODO

	return matrix, err
}

/*
 * Saves incoming MatrixElements in a matrix.
 */

func matrixSaver(matrix *[][]string, ch chan MatrixElement) {
	//TODO
}


/*
 * Creates a MatrixElement and saves it.
 */

func createAndSaveElement(row, column int, value string, ch chan MatrixElement) {
	//TODO
}

/*
 * Self-explanatory :).
 */

func stringToRational(num string) (big.Rat, err error) {
	var rational big.Rat
	_, ok := rational.SetString(num)
	if !ok {
		return rational, InvalidElementSyntaxError
	}
	return rational
}