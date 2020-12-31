package service

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/Sebalvarez97/mutants/api/errors"
	"github.com/Sebalvarez97/mutants/api/interfaces"
	"github.com/Sebalvarez97/mutants/api/model"
)

type MutantServiceImpl struct {
	repository interfaces.DnaRepository
	cerebro    interfaces.CerebroService
}

func NewMutantService(repository interfaces.DnaRepository, cerebro interfaces.CerebroService) interfaces.MutantService {
	return MutantServiceImpl{
		repository: repository,
		cerebro:    cerebro,
	}
}

const invalidNitrogenBaseFoundMessage = "invalid nitrogen base found: %q"
const invalidInputMatrixToShortMessage = "invalid input, the matrix is to short, has to be 4x4 or bigger"
const invalidInputNotAnNxNMatrixMessage = "invalid input, it isn't a NxN matrix, this could cause an Internal Error"

func (i MutantServiceImpl) IsMutant(input []string) (bool, *errors.ApiErrorImpl) {
	chain, hash, err := validateInput(input)
	if err != nil {
		return false, err
	}
	isMutant, sequences := i.cerebro.IsMutantDna(chain)
	if err := i.repository.Upsert(model.NewDna(chain, hash, isMutant, sequences)); err != nil {
		return false, err
	}
	return isMutant, nil
}

func (i MutantServiceImpl) GetMutantStats() (*model.Stats, *errors.ApiErrorImpl) {
	mutants, err := i.repository.FindAllMutants()
	if err != nil {
		return &model.Stats{}, err
	}
	humans, err := i.repository.FindAllHumans()
	if err != nil {
		return &model.Stats{}, err
	}

	m := len(mutants)
	h := len(humans)
	ratio := 1.0
	if h != 0 {
		ratio = float64(m) / float64(h)
	}
	return model.NewStats(len(mutants), len(humans), ratio), nil
}

var validInputs = map[byte]bool{
	byte('A'): true,
	byte('T'): true,
	byte('C'): true,
	byte('G'): true,
}

func validateInput(input []string) ([][]byte, string, *errors.ApiErrorImpl) {
	id := make(chan string)
	go GenerateHashDna(id, input)
	size := len(input)
	if size < 4 {
		err := errors.BadRequestError(fmt.Errorf(invalidInputMatrixToShortMessage))
		return nil, "", &err
	}
	dna := make([][]byte, len(input))
	for i, v := range input {
		if size != len(v) {
			err := errors.BadRequestError(fmt.Errorf(invalidInputNotAnNxNMatrixMessage))
			return nil, "", &err
		}
		dna[i] = []byte(v)
	}
	for _, v := range dna {
		for _, w := range v {
			if !validInputs[w] {
				err := errors.BadRequestError(fmt.Errorf(invalidNitrogenBaseFoundMessage, w))
				return nil, "", &err
			}
		}
	}
	return dna, <-id, nil
}

func GenerateHashDna(id chan<- string, input []string) {
	var sa string
	for _, v := range input {
		sa += v
	}
	h := sha1.New()
	h.Write([]byte(sa))
	bs := h.Sum(nil)
	id <- hex.EncodeToString(bs)
}
