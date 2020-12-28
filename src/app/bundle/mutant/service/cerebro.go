package service

import (
	"github.com/Sebalvarez97/mutants/src/app/common/errors"
)

type CerebroServiceImpl struct{}

func (i CerebroServiceImpl) IsMutant(dna [][]byte) (bool, *errors.ApiErrorImpl) {
	and := transpose(dna)
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
