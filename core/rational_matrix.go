package core

/*
 * ------ Imported packages ------
 */

import (
	"math/big"
)

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