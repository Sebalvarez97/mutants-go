package service

import (
	"fmt"
	. "github.com/Sebalvarez97/mutants/src/main/bundle/mutant/domain"
	. "github.com/Sebalvarez97/mutants/src/main/bundle/mutant/service"
	"github.com/Sebalvarez97/mutants/src/main/common/errors"
	"testing"
)

const serviceError = "Service is not working fine: %s"

const invalidInputMatrixToShortMessage = "Invalid value entered: invalid input, the matrix is to short, has to be 4x4 or bigger"
const invalidInputNotAnNxNMatrixMessage = "Invalid value entered: invalid input, it isn't a NxN matrix, this could cause an Internal Error"
const invalidNitrogenBaseFoundMessage = "Invalid value entered: invalid nitrogen base found: %q"

func TestNxMMatrix(t *testing.T) {
	sa := []string{
		"ATGCGA",
		"CAGTGC",
		"TTATGT",
		"AGAAGG",
		"TCACTG"}
	_, err := IsMutant(sa)
	if err == nil || err.Message != invalidInputNotAnNxNMatrixMessage {
		t.Error(fmt.Sprintf(serviceError, "an invalid matrix is valid"))
	}
}

func TestToShortMatrix(t *testing.T) {
	sa := []string{
		"ATGCGA",
		"CAGTGC",
		"TTATGT"}
	_, err := IsMutant(sa)
	if err == nil || err.Message != invalidInputMatrixToShortMessage {
		t.Error(fmt.Sprintf(serviceError, "an invalid matrix is valid"))
	}
}

func TestInvalidCharacterMatrix(t *testing.T) {
	b := byte('B')
	sa := []string{
		"ATGCGA",
		"CAG" + string(b) + "GC",
		"TTATGT",
		"AGAAGG",
		"CCCCTA",
		"TCACTG"}
	_, err := IsMutant(sa)
	if err == nil || err.Message != fmt.Sprintf(invalidNitrogenBaseFoundMessage, b) {
		t.Error(fmt.Sprintf(serviceError, "an invalid matrix is valid"))
	}
}

var cerebroServiceIsMutantMock func(input [][]byte) (bool, int)

type CerebroServiceImplMock struct{}

func (i CerebroServiceImplMock) IsMutantDna(input [][]byte) (bool, int) {
	return cerebroServiceIsMutantMock(input)
}

var dnaRepositoryUpsertMock func(dna *Dna) *errors.ApiErrorImpl
var dnaRepositoryFindAllMutantsMock func() ([]Dna, *errors.ApiErrorImpl)
var dnaRepositoryFindAllHumansMock func() ([]Dna, *errors.ApiErrorImpl)

type DnaRepositoryImplMock struct{}

func (i DnaRepositoryImplMock) FindAllMutants() ([]Dna, *errors.ApiErrorImpl) {
	return dnaRepositoryFindAllMutantsMock()
}
func (i DnaRepositoryImplMock) FindAllHumans() ([]Dna, *errors.ApiErrorImpl) {
	return dnaRepositoryFindAllHumansMock()
}
func (i DnaRepositoryImplMock) Upsert(dna *Dna) *errors.ApiErrorImpl {
	return dnaRepositoryUpsertMock(dna)
}

func TestIsMutantOk(t *testing.T) {
	sa := []string{
		"ATGCGA",
		"CAGAGC",
		"TTATGT",
		"AGAAGG",
		"CCCCTA",
		"TCACTG"}

	Cerebro = CerebroServiceImplMock{}
	Repository = DnaRepositoryImplMock{}
	cerebroServiceIsMutantMock = func(input [][]byte) (bool, int) {
		return true, 2
	}
	dnaRepositoryUpsertMock = func(dna *Dna) *errors.ApiErrorImpl {
		return nil
	}
	im, err := IsMutant(sa)
	if im != true || err != nil {
		t.Error(fmt.Sprintf(serviceError, "this mocked dna is mutant"))
	}
}

func TestIsNotMutantOk(t *testing.T) {
	sa := []string{
		"ATGCGA",
		"CAGAGC",
		"TTATGT",
		"AGAAGG",
		"CCCCTA",
		"TCACTG"}

	Cerebro = CerebroServiceImplMock{}
	Repository = DnaRepositoryImplMock{}
	cerebroServiceIsMutantMock = func(input [][]byte) (bool, int) {
		return false, 0
	}
	dnaRepositoryUpsertMock = func(dna *Dna) *errors.ApiErrorImpl {
		return nil
	}
	im, err := IsMutant(sa)
	if im != false || err != nil {
		t.Error(fmt.Sprintf(serviceError, "this mocked dna is human"))
	}
}

func TestIsMutantFailToInsert(t *testing.T) {
	sa := []string{
		"ATGCGA",
		"CAGAGC",
		"TTATGT",
		"AGAAGG",
		"CCCCTA",
		"TCACTG"}

	Cerebro = CerebroServiceImplMock{}
	Repository = DnaRepositoryImplMock{}
	cerebroServiceIsMutantMock = func(input [][]byte) (bool, int) {
		return false, 0
	}
	dnaRepositoryUpsertMock = func(dna *Dna) *errors.ApiErrorImpl {
		apiErr := errors.GenericError(fmt.Errorf("error on insert opeartion: %s", "generic"))
		return &apiErr
	}
	im, err := IsMutant(sa)
	if im != false || err == nil || err.Message != fmt.Sprintf("Server failed to perform request because of %s", fmt.Sprintf("error on insert opeartion: %s", "generic")) {
		t.Error(fmt.Sprintf(serviceError, "this has to fail because an insertion error"))
	}
}

func TestGetMutantStatsOk(t *testing.T) {
	Repository = DnaRepositoryImplMock{}
	m := 40
	h := 100
	ratio := 0.4
	dnaRepositoryFindAllHumansMock = func() ([]Dna, *errors.ApiErrorImpl) {
		return make([]Dna, h), nil
	}
	dnaRepositoryFindAllMutantsMock = func() ([]Dna, *errors.ApiErrorImpl) {
		return make([]Dna, m), nil
	}
	s, err := GetMutantStats()
	if s == nil || err != nil || s.Humans != h || s.Mutants != m || s.Ratio != ratio {
		t.Error(fmt.Sprintf(serviceError, "fail to get stats"))
	}
}

func TestGetMutantStatsOkWithCeroHumans(t *testing.T) {
	Repository = DnaRepositoryImplMock{}
	m := 40
	h := 0
	ratio := 1.0
	dnaRepositoryFindAllHumansMock = func() ([]Dna, *errors.ApiErrorImpl) {
		return make([]Dna, h), nil
	}
	dnaRepositoryFindAllMutantsMock = func() ([]Dna, *errors.ApiErrorImpl) {
		return make([]Dna, m), nil
	}
	s, err := GetMutantStats()
	if s == nil || err != nil || s.Humans != h || s.Mutants != m || s.Ratio != ratio {
		t.Error(fmt.Sprintf(serviceError, "fail to get stats"))
	}
}

func TestGetMutantStatsOkWithCeroMutants(t *testing.T) {
	Repository = DnaRepositoryImplMock{}
	m := 0
	h := 100
	ratio := 0.0
	dnaRepositoryFindAllHumansMock = func() ([]Dna, *errors.ApiErrorImpl) {
		return make([]Dna, h), nil
	}
	dnaRepositoryFindAllMutantsMock = func() ([]Dna, *errors.ApiErrorImpl) {
		return make([]Dna, m), nil
	}
	s, err := GetMutantStats()
	if s == nil || err != nil || s.Humans != h || s.Mutants != m || s.Ratio != ratio {
		t.Error(fmt.Sprintf(serviceError, "fail to get stats"))
	}
}

func TestGenerateHashEqualDna(t *testing.T) {
	sa := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
	sa2 := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
	result := "6d32697470c08f971ea2f5a71113166f9abc2d7f"

	id := make(chan string)
	id2 := make(chan string)
	go GenerateHashDna(id, sa)
	go GenerateHashDna(id2, sa2)

	var r1 string
	r1 = <-id

	var r2 string
	r2 = <-id2

	if r1 != r2 || r1 != result {
		t.Error(fmt.Sprintf(serviceError, "fail to create equal hash"))
	}
}

func TestGenerateHashDiferentDna(t *testing.T) {
	sa := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
	sa2 := []string{"AAGAGA", "CAGGGC", "TCATGT", "AGACGG", "CCCCTA", "TCACTG"}
	result := "6d32697470c08f971ea2f5a71113166f9abc2d7f"
	result2 := "3b57ab803423e48543f27602a63618433cc23c37"

	id := make(chan string)
	id2 := make(chan string)
	go GenerateHashDna(id, sa)
	go GenerateHashDna(id2, sa2)

	var r1 string
	r1 = <-id

	var r2 string
	r2 = <-id2

	if r1 == r2 || r1 != result || r2 != result2 {
		t.Error(fmt.Sprintf(serviceError, "fail to create diferent hash"))
	}
}