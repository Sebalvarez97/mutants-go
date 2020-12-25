package main

import (
	"fmt"
)

func transpose(a [][]byte) [][]byte {
	newArr := make([][]byte, len(a))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			newArr[j] = append(newArr[j], a[i][j])
		}
	}

	return newArr
}

func checkNext(x int, v byte, row []byte) int {
	y := x + 1
	if y >= len(row) {
		return 0
	}
	//fmt.Printf("\nCheckRight in %v|%v: \n", x, y)
	//fmt.Printf("Comparing : %v to %v \n", v, (row)[y])
	if v == (row)[y] {
		cn := checkNext(y, (row)[y], row)
		if cn <= 3 {
			return cn + 1
		}
	}
	return 0
}

func checkDiagonalUpRight(x int, y int, v byte, a [][]byte) int {
	xi := x + 1
	yi := y + 1
	if xi >= len(a) {
		return 0
	}
	row := a[xi]
	if yi >= len(row) {
		return 0
	}
	vi := row[yi]
	//fmt.Printf("\nCheckDiagonal in %v|%v: \n", x, y)
	//fmt.Printf("Comparing : %v to %v \n", v, vi)
	if v == vi {
		cd := checkDiagonalUpRight(xi, yi, vi, a)
		if cd <= 3 {
			return cd + 1
		}
	}
	return 0
}

func checkDiagonalDownLeft(x int, y int, v byte, a [][]byte) int {
	xi := x - 1
	yi := y - 1
	if xi < 0 {
		return 0
	}
	row := a[xi]
	if yi < 0 {
		return 0
	}
	vi := row[yi]
	//fmt.Printf("\nCheckDiagonal in %v|%v: \n", x, y)
	//fmt.Printf("Comparing : %v to %v \n", v, vi)
	if v == vi {
		cd := checkDiagonalDownLeft(xi, yi, vi, a)
		if cd <= 3 {
			return cd + 1
		}
	}
	return 0
}

func isMutant(dna [][]byte) bool {
	//fmt.Printf("Printing nitrogen bases for dna %v \n", dna)
	and := transpose(dna)
	//fmt.Printf("Printing nitrogen bases for and %v \n", and)
	ms := 0
	if len(dna) != len(and) {
		return false
	}
	for ri := 0; ri < len(dna); ri++ {
		if ms > 1 {
			break
		}
		row1 := dna[ri]
		row2 := and[ri]
		//fmt.Printf("\nRows %v : %+q - %v | %+q - %v \n", ri, row1, row1, row2, row2)
		for ci := 0; ci < len(row1); ci++ {
			sr1 := row1[ci:]
			sr2 := row2[ci:]
			//fmt.Printf("Subrows: %v | %v \n", sr1, sr2)
			result1 := checkNext(ci, row1[ci], sr1)
			result2 := checkNext(ci, row2[ci], sr2)
			result3 := checkDiagonalUpRight(ri, ci, row1[ci], dna)
			result4 := checkDiagonalDownLeft(ri, ci, row1[ci], dna)
			if result1 >= 3 {
				ms++
			}
			if result2 >= 3 {
				ms++
			}
			if result3 >= 3 {
				ms++
			}
			if result4 >= 3 {
				ms++
			}
			//fmt.Printf("Results of checking: %v | %v | %v \n", result1, result2, result3)
		}

	}
	fmt.Printf("\nMutant sequences: %v \n", ms)
	return ms > 1
}

func main() {

	fmt.Println("Hello, I'm cerebro.")

	input := [...]string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}

	//fmt.Printf("Printing input: %+q \n", input)

	dna := make([][]byte, len(input))
	//fmt.Printf("%T\n\n", dna)
	for i, v := range input {
		dna[i] = []byte(v)
	}

	im := isMutant(dna)
	if im {
		fmt.Println("I found a mutant!!")
	} else {
		fmt.Println("No luck this time...")
	}

}
