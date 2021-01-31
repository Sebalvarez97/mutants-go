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

var dnaRepositoryFindNumberOfHumansMock func() (int, *errors.ApiErrorImpl)
var dnaRepositoryFindNumberOfMutantsMock func() (int, *errors.ApiErrorImpl)

type DnaRepositoryImplMock struct{}

func (i DnaRepositoryImplMock) FindAllMutants() ([]Dna, *errors.ApiErrorImpl) {
	panic("implement me")
}

func (i DnaRepositoryImplMock) FindAllHumans() ([]Dna, *errors.ApiErrorImpl) {
	panic("implement me")
}

func (i DnaRepositoryImplMock) FindNumberOfHumans() (int, *errors.ApiErrorImpl) {
	return dnaRepositoryFindNumberOfHumansMock()
}

func (i DnaRepositoryImplMock) FindNumberOfMutants() (int, *errors.ApiErrorImpl) {
	return dnaRepositoryFindNumberOfMutantsMock()
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
	dnaRepositoryFindNumberOfHumansMock = func() (int, *errors.ApiErrorImpl) {
		return h, nil
	}
	dnaRepositoryFindNumberOfMutantsMock = func() (int, *errors.ApiErrorImpl) {
		return m, nil
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
	dnaRepositoryFindNumberOfHumansMock = func() (int, *errors.ApiErrorImpl) {
		return h, nil
	}
	dnaRepositoryFindNumberOfMutantsMock = func() (int, *errors.ApiErrorImpl) {
		return m, nil
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
	dnaRepositoryFindNumberOfHumansMock = func() (int, *errors.ApiErrorImpl) {
		return h, nil
	}
	dnaRepositoryFindNumberOfMutantsMock = func() (int, *errors.ApiErrorImpl) {
		return m, nil
	}

	service := NewMutantService(repository, nil)
	s, err := service.GetMutantStats()
	if s == nil || err != nil || s.Humans != h || s.Mutants != m || s.Ratio != ratio {
		t.Error(fmt.Sprintf(serviceError, "fail to get stats"))
	}
}
