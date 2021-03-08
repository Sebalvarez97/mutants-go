package matrix

import (
	"fmt"
	"testing"
)

const hashingError = "Hash is not working fine: %s"

func TestGenerateHashEqualDna(t *testing.T) {
	sa := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
	sa2 := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
	result := "6d32697470c08f971ea2f5a71113166f9abc2d7f"

	id := make(chan string)
	id2 := make(chan string)
	go func([]string, chan string) {
		id <- GenerateHashForStringArray(sa)
	}(sa, id)

	go func([]string, chan string) {
		id2 <- GenerateHashForStringArray(sa2)
	}(sa2, id2)

	var r1 string
	r1 = <-id

	var r2 string
	r2 = <-id2

	if r1 != r2 || r1 != result {
		t.Error(fmt.Sprintf(hashingError, "fail to create equal hash"))
	}
}

func TestGenerateHashDiferentDna(t *testing.T) {
	sa := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
	sa2 := []string{"AAGAGA", "CAGGGC", "TCATGT", "AGACGG", "CCCCTA", "TCACTG"}
	result := "6d32697470c08f971ea2f5a71113166f9abc2d7f"
	result2 := "3b57ab803423e48543f27602a63618433cc23c37"

	id := make(chan string)
	id2 := make(chan string)
	go func([]string, chan string) {
		id <- GenerateHashForStringArray(sa)
	}(sa, id)

	go func([]string, chan string) {
		id2 <- GenerateHashForStringArray(sa2)
	}(sa2, id2)

	var r1 string
	r1 = <-id

	var r2 string
	r2 = <-id2

	if r1 == r2 || r1 != result || r2 != result2 {
		t.Error(fmt.Sprintf(hashingError, "fail to create diferent hash"))
	}
}
