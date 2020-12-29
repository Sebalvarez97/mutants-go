package service

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	. "github.com/Sebalvarez97/mutants/src/app/bundle/mutant/domain"
	. "github.com/Sebalvarez97/mutants/src/app/bundle/mutant/repository"
	"github.com/Sebalvarez97/mutants/src/app/common/errors"
)

type MutantServiceImpl struct{}

type CerebroService interface {
	IsMutant(input [][]byte) (bool, int, *errors.ApiErrorImpl)
}

var cerebro CerebroService = CerebroServiceImpl{}

type DnaRepository interface {
	Upsert(dna *Dna) *errors.ApiErrorImpl
}

var repository = DnaRepositoryImpl{}

func (i MutantServiceImpl) IsMutant(input []string) (bool, *errors.ApiErrorImpl) {
	chain, hash, err := validateInput(input)
	if err != nil {
		return false, err
	}
	isMutant, sequences, err := cerebro.IsMutant(chain)
	if err != nil {
		return false, err
	}
	if err := repository.Upsert(NewDna(chain, hash, isMutant, sequences)); err != nil {
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

func validateInput(input []string) ([][]byte, string, *errors.ApiErrorImpl) {
	id := make(chan string)
	go generateHashDna(id, input)
	size := len(input)
	if size < 4 {
		err := errors.BadRequestError(fmt.Errorf("invalid input, the matrix is to short, has to be 4x4 or bigger"))
		return nil, "", &err
	}
	dna := make([][]byte, len(input))
	for i, v := range input {
		if size != len(v) {
			err := errors.BadRequestError(fmt.Errorf("invalid input, it isn't a NxN matrix, this could cause an Internal Error"))
			return nil, "", &err
		}
		dna[i] = []byte(v)
	}
	for _, v := range dna {
		for _, w := range v {
			if !validInputs[w] {
				err := errors.BadRequestError(fmt.Errorf("invalid nitrogen base found: %q ", w))
				return nil, "", &err
			}
		}
	}
	return dna, <-id, nil
}

func generateHashDna(id chan<- string, input []string) {
	var sa string
	for _, v := range input {
		sa += v
	}
	h := sha1.New()
	h.Write([]byte(sa))
	bs := h.Sum(nil)
	id <- hex.EncodeToString(bs)
}
