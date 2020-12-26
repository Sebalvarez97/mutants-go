package service

import (
	"fmt"
	errors "github.com/Sebalvarez97/mutants/src/app/common"
)

type CerebroServiceImpl struct{}

func (i CerebroServiceImpl) IsMutant(input []string) (bool, *errors.ApiErrorImpl) {

	dna := make([][]byte, len(input))
	//fmt.Printf("%T\n\n", dna)
	for i, v := range input {
		dna[i] = []byte(v)
	}
	and, err := transposeAndValidate(dna)

	if err != nil {
		return false, err
	}

	ms := 0
	for ri := 0; ri < len(dna); ri++ {
		if ms > 1 {
			break
		}
		row1 := dna[ri]
		row2 := and[ri]
		for ci := 0; ci < len(row1); ci++ {
			result1 := checkNext(ci, row1[ci], row1)
			if result1 >= 3 {
				ms++
			}
			if ms > 1 {
				break
			}
			result2 := checkNext(ci, row2[ci], row2)
			if result2 >= 3 {
				ms++
			}
			if ms > 1 {
				break
			}
			result3 := checkDiagonalUpRight(ri, ci, row1[ci], dna)
			if result3 >= 3 {
				ms++
			}
			if ms > 1 {
				break
			}
			result4 := checkDiagonalDownLeft(ri, ci, row1[ci], dna)
			if result4 >= 3 {
				ms++
			}
		}

	}
	return ms > 1, nil
}

func transposeAndValidate(a [][]byte) ([][]byte, *errors.ApiErrorImpl) {
	newArr := make([][]byte, len(a))
	for i := 0; i < len(a); i++ {
		if len(a[i]) != len(a) {
			err := errors.BadRequestError(fmt.Errorf("invalid input, it isn't a NxN matrix, this could cause an Internal Error"))
			return newArr, &err
		}
		for j := 0; j < len(a[0]); j++ {
			if !isValidDna(a[i][j]) {
				err := errors.BadRequestError(fmt.Errorf("invalid nitrogen base found: %q ", a[i][j]))
				return newArr, &err
			}
			newArr[j] = append(newArr[j], a[i][j])
		}
	}
	return newArr, nil
}

func checkNext(x int, v byte, row []byte) int {
	y := x + 1
	if y >= len(row) {
		return 0
	}
	if v == row[y] {
		cn := checkNext(y, row[y], row)
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
	if v == vi {
		cd := checkDiagonalUpRight(xi, yi, vi, a)
		if cd <= 3 {
			return cd + 1
		}
	}
	return 0
}

func checkDiagonalDownLeft(x int, y int, v byte, a [][]byte) int {
	xi := x + 1
	yi := y - 1
	if xi >= len(a) {
		return 0
	}
	row := a[xi]
	if yi < 0 {
		return 0
	}
	vi := row[yi]
	if v == vi {
		cd := checkDiagonalDownLeft(xi, yi, vi, a)
		if cd <= 3 {
			return cd + 1
		}
	}
	return 0
}

var validInputs = []byte{
	byte('A'),
	byte('T'),
	byte('G'),
	byte('C'),
}

func isValidDna(bn byte) bool {
	for _, v := range validInputs {
		if v == bn {
			return true
		}
	}
	return false
}
