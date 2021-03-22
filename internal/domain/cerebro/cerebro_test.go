package cerebro

import (
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
