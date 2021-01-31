package service

import (
	"fmt"
	"github.com/Sebalvarez97/mutants/api/errors"
	. "github.com/Sebalvarez97/mutants/api/model"
	"testing"
)

const serviceError = "Service is not working fine: %s"

var cerebroServiceIsMutantMock func(input [][]byte) (bool, int)

type CerebroServiceImplMock struct{}

func (i CerebroServiceImplMock) IsMutantDna(input [][]byte) (bool, int) {
	return cerebroServiceIsMutantMock(input)
}

var dnaRepositoryFindAllMutantsMock func() ([]Dna, *errors.ApiErrorImpl)
var dnaRepositoryFindAllHumansMock func() ([]Dna, *errors.ApiErrorImpl)

type DnaRepositoryImplMock struct{}

func (i DnaRepositoryImplMock) FindAllMutants() ([]Dna, *errors.ApiErrorImpl) {
	return dnaRepositoryFindAllMutantsMock()
}
func (i DnaRepositoryImplMock) FindAllHumans() ([]Dna, *errors.ApiErrorImpl) {
	return dnaRepositoryFindAllHumansMock()
}
func (i DnaRepositoryImplMock) Upsert(*Dna) {}

func TestIsMutantOk(t *testing.T) {
	sa := []string{
		"ATGCGA",
		"CAGAGC",
		"TTATGT",
		"AGAAGG",
		"CCCCTA",
		"TCACTG"}

	cerebro := CerebroServiceImplMock{}
	repository := DnaRepositoryImplMock{}
	cerebroServiceIsMutantMock = func(input [][]byte) (bool, int) {
		return true, 2
	}

	request := IsMutantRequestBody{Dna: sa}
	service := NewMutantService(repository, cerebro)
	im := service.IsMutant(request)
	if im != true {
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

	cerebro := CerebroServiceImplMock{}
	repository := DnaRepositoryImplMock{}
	cerebroServiceIsMutantMock = func(input [][]byte) (bool, int) {
		return false, 0
	}

	request := IsMutantRequestBody{Dna: sa}
	service := NewMutantService(repository, cerebro)
	im := service.IsMutant(request)
	if im != false {
		t.Error(fmt.Sprintf(serviceError, "this mocked dna is human"))
	}
}

func TestGetMutantStatsOk(t *testing.T) {
	repository := DnaRepositoryImplMock{}
	m := 40
	h := 100
	ratio := 0.4
	dnaRepositoryFindAllHumansMock = func() ([]Dna, *errors.ApiErrorImpl) {
		return make([]Dna, h), nil
	}
	dnaRepositoryFindAllMutantsMock = func() ([]Dna, *errors.ApiErrorImpl) {
		return make([]Dna, m), nil
	}

	service := NewMutantService(repository, nil)
	s, err := service.GetMutantStats()
	if s == nil || err != nil || s.Humans != h || s.Mutants != m || s.Ratio != ratio {
		t.Error(fmt.Sprintf(serviceError, "fail to get stats"))
	}
}

func TestGetMutantStatsOkWithCeroHumans(t *testing.T) {
	repository := DnaRepositoryImplMock{}
	m := 40
	h := 0
	ratio := 1.0
	dnaRepositoryFindAllHumansMock = func() ([]Dna, *errors.ApiErrorImpl) {
		return make([]Dna, h), nil
	}
	dnaRepositoryFindAllMutantsMock = func() ([]Dna, *errors.ApiErrorImpl) {
		return make([]Dna, m), nil
	}

	service := NewMutantService(repository, nil)
	s, err := service.GetMutantStats()
	if s == nil || err != nil || s.Humans != h || s.Mutants != m || s.Ratio != ratio {
		t.Error(fmt.Sprintf(serviceError, "fail to get stats"))
	}
}

func TestGetMutantStatsOkWithCeroMutants(t *testing.T) {
	repository := DnaRepositoryImplMock{}
	m := 0
	h := 100
	ratio := 0.0
	dnaRepositoryFindAllHumansMock = func() ([]Dna, *errors.ApiErrorImpl) {
		return make([]Dna, h), nil
	}
	dnaRepositoryFindAllMutantsMock = func() ([]Dna, *errors.ApiErrorImpl) {
		return make([]Dna, m), nil
	}

	service := NewMutantService(repository, nil)
	s, err := service.GetMutantStats()
	if s == nil || err != nil || s.Humans != h || s.Mutants != m || s.Ratio != ratio {
		t.Error(fmt.Sprintf(serviceError, "fail to get stats"))
	}
}
