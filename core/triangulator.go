package core

/*
 * ------ Imported packages ------
 */

import (
	"math"
	"math/big"
	"sort"
)

/*
 * ------ Functions ------
 */

/*
 * Concurrent Gaussian elimination.
 */

func (matrix *RationalMatrix) Triangulate() {
	colLength := len(*matrix)
	rowLength := len((*matrix)[0])
	columnsToOperate :=  int(math.Min(float64(rowLength), float64(colLength)))
	//For each column
	for colIndex := 0; colIndex < columnsToOperate; colIndex++ {
		//Select pivot
		pivotX := colIndex
		pivotY := 0
		for pivotY < colLength && !validPivot(pivotX, pivotY, matrix) {
			pivotY++
		}

		if pivotY < colLength {
			//Operate on each row below the column's pivot.
			pivot := (*matrix)[pivotY][pivotX]
			for y := pivotY + 1; y < colLength; y++ {
				columnNum := (*matrix)[y][pivotX]
				if !isZero(columnNum) {
					quotient:= Quotient(&pivot, &columnNum)
					MultiplyRow(y, quotient, matrix)
					SumRows(y, pivotY, matrix)
				}
			}
		}
	}
	//Rearrange rows to make the matrix upper triangular
	//Rows with less zeroes left go up, see the implementation of Interface
	sort.Sort(*matrix)
	//prettyFormat(matrix) e.g (1/8 2/8 2/8 9/8) row -> (1 2 2 9)
}

func validPivot(pivotX, pivotY int, matrix *RationalMatrix) bool {
	pivot := (*matrix)[pivotY][pivotX]
	switch {
		case isZero(pivot):
			return false
		case pivotX == 0:
			return true
		case hasZerosLeft(pivotX, pivotY, matrix):
			return true
	}
	return false
}

func Quotient(a, b *big.Rat) big.Rat {
	var q big.Rat
	return Minus(*q.Quo(a,b))
}

/*
 * Check if a row has zeroes until a certain column
 */

func hasZerosLeft(x, y int, matrix *RationalMatrix) bool {
	ch := make(chan bool)
	result := true
	for i := 0; i < x; i++ {
		go checkZero((*matrix)[y][i], ch)
	}
	for i := 0; i < x; i++ {
		result = result && <-ch
	}
	return result
}

func checkZero (r big.Rat, ch chan bool) {
	ch <- isZero(r)
}

func MultiplyRow(rowIndex int, scalar big.Rat, matrix *RationalMatrix) {
	done := make(chan bool)
	rowLength := len((*matrix)[rowIndex])
	for index, _ := range (*matrix)[rowIndex] {
		go multiplyRational(rowIndex, index, scalar, matrix, done)
	}
	for i := 0; i < rowLength; i++ {
		<- done
	}
}

func multiplyRational(rowIndex, columnIndex int, rational big.Rat, matrix *RationalMatrix, ch chan bool) {
	(*matrix)[rowIndex][columnIndex].Mul(&rational, &(*matrix)[rowIndex][columnIndex])
	ch <- true
}

/*
 * Sum is stored in the fist row.
 */

func SumRows(fstRowIndex, sndRowIndex int, matrix *RationalMatrix) {
	done := make(chan bool)
	rowLength := len((*matrix)[fstRowIndex])
	for i := 0; i < rowLength; i++ {
		 go SumRationals(fstRowIndex, sndRowIndex, i, matrix, done)
	}
	for i := 0; i < rowLength; i++ {
		<- done
	}
}

func SumRationals(fstRowIndex, sndRowIndex, columnIndex int, matrix *RationalMatrix, ch chan bool) {
	(*matrix)[fstRowIndex][columnIndex].Add(&(*matrix)[fstRowIndex][columnIndex], &(*matrix)[sndRowIndex][columnIndex])
	ch <- true
}
