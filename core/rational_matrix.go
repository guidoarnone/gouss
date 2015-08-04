package core

/*
 * ------ Imported packages ------
 */

import (
	"math/big"
)

/*
 * ------ Variables ------
 */

var zero = *big.NewRat(0,1)
var minusOne = *big.NewRat(-1,1)

/*
 * ------ Types and Methods ------
 */

type RationalMatrix [][]big.Rat

/*
 * Satisfying the sort.Interface interface, sorting by less zeroes left.
 */

func (r RationalMatrix) Len() int {
	return len(r)
}

func (r RationalMatrix) Less(i, j int) bool {
	for k := 0; k < len(r[0]); k++ {
		if !isZero(r[i][k]) && isZero(r[j][k]) {
			return true
		} else if isZero(r[i][k]) && !isZero(r[j][k]) {
			return false
		}
	}
	return true
}

func (r RationalMatrix) Swap(i, j int) {
	rowSwap := r[i]
	r[i] = r[j]
	r[j] = rowSwap
}

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

/*
 * Functions for rationals.
 */

func isZero(q big.Rat) bool {
	return q.Cmp(&zero) == 0
}

func Minus(q big.Rat) big.Rat {
	return *q.Mul(&q, &minusOne)
}
