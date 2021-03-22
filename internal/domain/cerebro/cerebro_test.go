package cerebro

import (
	"fmt"
	utils "github.com/Sebalvarez97/mutants-go/tools/matrix"
	"testing"
)

var validDna = []byte("AGCT")

func BenchmarkIsMutant(b *testing.B) {
	matrix := utils.GenerateNxNMatrix(10000, validDna)
	b.ResetTimer()
	b.Run("IsMutantNew", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			IsMutant(matrix)
		}
	})
}

func TestIsMutant(t *testing.T) {
	t.Run("IsMutantWithMutantMatrix", func(t *testing.T) {
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
		result, sequences := IsMutant(dna)
		if !result && sequences != 2 {
			t.Error(fmt.Sprintf("cerebro is failing to detect mutants"))
		}
	})
	t.Run("IsMutantWithNonMutantMatrix", func(t *testing.T) {
		sa := []string{
			"ATGCGA",
			"CAGTGC",
			"TTATGT",
			"AGAGTG",
			"CCCGTA",
			"TCACTG"}
		dna := make([][]byte, len(sa))
		for i, v := range sa {
			dna[i] = []byte(v)
		}
		result, sequences := IsMutant(dna)
		if result && sequences != 0 {
			t.Error(fmt.Sprintf("cerebro is detecting a human as a mutant"))
		}
	})
}
