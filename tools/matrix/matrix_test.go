package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTranspose(t *testing.T) {
	sa := []string{
		"ATGCGA",
		"CAGTGC",
		"TTATGT",
		"AGAAGG",
		"CCCCTA",
		"TCACTG"}
	dna := make([][]byte, len(sa))
	for i, v := range sa {
		dna[i] = []byte(v)
	}
	transpose := Transpose(dna)
	for x, row := range transpose {
		for y, cell := range row {
			assert.Equal(t, dna[y][x], cell)
		}
	}
}

func TestDiagonals(t *testing.T) {
	sa := []string{
		"ATGCGA",
		"CAGTGC",
		"TTATGT",
		"AGAAGG",
		"CCCCTA",
		"TCACTG"}
	dna := make([][]byte, len(sa))
	for i, v := range sa {
		dna[i] = []byte(v)
	}
	diagonals := Diagonals(dna)

	assert.Equal(t, 6, len(diagonals))
	assert.Equal(t, []byte("AAAATG"), diagonals[0])
	assert.Equal(t, []byte("CTACT"), diagonals[1])
	assert.Equal(t, []byte("TGCC"), diagonals[2])
	assert.Equal(t, []byte("ACA"), diagonals[3])
	assert.Equal(t, []byte("CC"), diagonals[4])
	assert.Equal(t, []byte("T"), diagonals[5])
}

func TestReverse(t *testing.T) {
	sa := []string{
		"ATGCGA",
		"CAGTGC",
		"TTATGT",
		"AGAAGG",
		"CCCCTA",
		"TCACTG"}
	dna := make([][]byte, len(sa))
	for i, v := range sa {
		dna[i] = []byte(v)
	}
	reverse := Reverse(dna)

	i := len(dna) - 1
	for _, row := range dna {
		assert.Equal(t, row, reverse[i])
		i--
	}
}

func TestGenerateNxNMatrix(t *testing.T) {
	size := 1000
	values := []byte("AGT561")

	posible := map[byte]bool{
		byte('A'): true,
		byte('G'): true,
		byte('T'): true,
		byte('5'): true,
		byte('6'): true,
		byte('1'): true,
	}

	generated := GenerateNxNMatrix(size, values)

	assert.Equal(t, size, len(generated))
	for _, row := range generated {
		assert.Equal(t, size, len(row))
		for _, cell := range row {
			assert.True(t, posible[cell])
		}
	}
}
