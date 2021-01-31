package model

import "fmt"

const InvalidNitrogenBaseFoundMessage = "invalid nitrogen base found: %v"
const InvalidInputMatrixToShortMessage = "invalid input, the matrix is to short, has to be 4x4 or bigger"
const InvalidInputNotAnNxNMatrixMessage = "invalid input, it isn't a NxN matrix, this could cause an Internal Error"

type IsMutantRequestBody struct {
	Dna []string `form:"dna" json:"dna" binding:"required"`
}

var validDna = map[string]bool{
	"A": true,
	"T": true,
	"C": true,
	"G": true,
}

func (i IsMutantRequestBody) IsValid() (bool, string) {
	input := i.Dna
	size := len(input)
	if size < 4 {
		return false, InvalidInputMatrixToShortMessage
	}
	for _, v := range input {
		if size != len(v) {
			return false, InvalidInputNotAnNxNMatrixMessage
		}
	}
	for _, v := range input {
		for _, w := range v {
			word := string(w)
			if !validDna[word] {
				return false, fmt.Sprintf(InvalidNitrogenBaseFoundMessage, word)
			}
		}
	}
	return true, ""
}
