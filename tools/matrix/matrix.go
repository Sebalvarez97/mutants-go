package matrix

import "math/rand"

func Transpose(a [][]byte) [][]byte {
	newArr := make([][]byte, len(a))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			newArr[j] = append(newArr[j], a[i][j])
		}
	}
	return newArr
}

func Diagonals(a [][]byte) [][]byte {
	var returnable [][]byte
	x := 0
	for {
		if x >= len(a) {
			break
		}
		y := 0
		x1 := x
		var dr []byte
		for {
			if y >= len(a[x]) || x1 >= len(a) {
				break
			}
			dr = append(dr, a[x1][y])
			x1++
			y++
		}
		returnable = append(returnable, dr)
		x++
	}
	return returnable
}

func Reverse(a [][]byte) [][]byte {
	var returnable [][]byte
	x := len(a) - 1
	for {
		if x < 0 {
			break
		}
		returnable = append(returnable, a[x])
		x--
	}
	return returnable
}

func GenerateNxNMatrix(size int, values []byte) [][]byte {
	matrix := make([][]byte, size)
	for x, row := range matrix {
		row = make([]byte, size)
		for y, _ := range row {
			min := 0
			max := len(values)
			v := rand.Intn(max-min) - min
			row[y] = values[v]
		}
		matrix[x] = row
	}
	return matrix
}


