package model

type Stats struct {
	Mutants int     `json:"count_mutant_dna"`
	Humans  int     `json:"count_human_dna"`
	Ratio   float64 `json:"ratio"`
}

func NewStats(mutants int, humans int, ratio float64) *Stats {
	return &Stats{
		Mutants: mutants,
		Humans:  humans,
		Ratio:   ratio,
	}
}
