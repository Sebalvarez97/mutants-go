package cerebro

/*
import (
	"fmt"
	"testing"
)

const cerebroError = "cerebro is not working fine: %s"

func TestCerebroMutant(t *testing.T) {
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

	service := NewCerebroService()

	result, sequences := service.IsMutantDna(dna)
	if !result && sequences != 2 {
		t.Error(fmt.Sprintf(cerebroError, "do not detecte mutants"))
	}
}

func TestCerebroHuman(t *testing.T) {
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

	service := NewCerebroService()

	result, sequences := service.IsMutantDna(dna)
	if result && sequences != 0 {
		t.Error(fmt.Sprintf(cerebroError, "it detected a human as a mutant"))
	}
}
 */

