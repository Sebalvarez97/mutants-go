package model

type IsMutantRequestBody struct {
	Dna []string `form:"dna" json:"dna" binding:"required"`
}
