package service

import (
	"fmt"
	. "github.com/Sebalvarez97/mutants/src/app/bundle/mutant/domain"
	. "github.com/Sebalvarez97/mutants/src/app/bundle/mutant/repository"
	"github.com/Sebalvarez97/mutants/src/app/common/errors"
)

type MutantServiceImpl struct{}

type CerebroService interface {
	IsMutant(input [][]byte) (bool, *errors.ApiErrorImpl)
}

var cerebro CerebroService = CerebroServiceImpl{}

type DnaRepository interface {
	Insert(dna *Dna) *errors.ApiErrorImpl
}

var repository = DnaRepositoryImpl{}

func (i MutantServiceImpl) IsMutant(input []string) (bool, *errors.ApiErrorImpl) {
	dna, err := validateInput(input)
	if err != nil {
		return false, err
	}
	isMutant, err := cerebro.IsMutant(dna)
	if err != nil {
		return false, err
	}
	if err := repository.Insert(NewDna([][]byte{}, isMutant, 0)); err != nil {
		return false, err
	}
	return isMutant, nil
}

var validInputs = map[byte]bool{
	byte('A'): true,
	byte('T'): true,
	byte('C'): true,
	byte('G'): true,
}

func validateInput(input []string) ([][]byte, *errors.ApiErrorImpl) {
	size := len(input)
	if size < 4 {
		err := errors.BadRequestError(fmt.Errorf("invalid input, the matrix is to short, has to be 4x4 or bigger"))
		return nil, &err
	}
	dna := make([][]byte, len(input))
	for i, v := range input {
		if size != len(v) {
			err := errors.BadRequestError(fmt.Errorf("invalid input, it isn't a NxN matrix, this could cause an Internal Error"))
			return nil, &err
		}
		dna[i] = []byte(v)
	}
	for _, v := range dna {
		for _, w := range v {
			if !validInputs[w] {
				err := errors.BadRequestError(fmt.Errorf("invalid nitrogen base found: %q ", w))
				return nil, &err
			}
		}
	}
	return dna, nil
}
