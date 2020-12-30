package service

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	. "github.com/Sebalvarez97/mutants/src/main/bundle/mutant/domain"
	. "github.com/Sebalvarez97/mutants/src/main/bundle/mutant/repository"
	"github.com/Sebalvarez97/mutants/src/main/common/errors"
)

type CerebroService interface {
	IsMutantDna(input [][]byte) (bool, int)
}
type CerebroServiceImpl struct{}

func (i CerebroServiceImpl) IsMutantDna(input [][]byte) (bool, int) {
	return IsMutantDna(input)
}

type DnaRepository interface {
	Upsert(dna *Dna) *errors.ApiErrorImpl
	FindAllMutants() ([]Dna, *errors.ApiErrorImpl)
	FindAllHumans() ([]Dna, *errors.ApiErrorImpl)
}
type DnaRepositoryImpl struct{}

func (i DnaRepositoryImpl) Upsert(dna *Dna) *errors.ApiErrorImpl {
	return Upsert(dna)
}
func (i DnaRepositoryImpl) FindAllMutants() ([]Dna, *errors.ApiErrorImpl) {
	return FindAllMutants()
}
func (i DnaRepositoryImpl) FindAllHumans() ([]Dna, *errors.ApiErrorImpl) {
	return FindAllHumans()
}

var Cerebro CerebroService
var Repository DnaRepository

func init() {
	Cerebro = CerebroServiceImpl{}
	Repository = DnaRepositoryImpl{}
}

const invalidInputMatrixToShortMessage = "invalid input, the matrix is to short, has to be 4x4 or bigger"
const invalidInputNotAnNxNMatrixMessage = "invalid input, it isn't a NxN matrix, this could cause an Internal Error"

func IsMutant(input []string) (bool, *errors.ApiErrorImpl) {
	chain, hash, err := validateInput(input)
	if err != nil {
		return false, err
	}
	isMutant, sequences := Cerebro.IsMutantDna(chain)
	if err := Repository.Upsert(NewDna(chain, hash, isMutant, sequences)); err != nil {
		return false, err
	}
	return isMutant, nil
}

func GetMutantStats() (*Stats, *errors.ApiErrorImpl) {
	mutants, err := Repository.FindAllMutants()
	if err != nil {
		return &Stats{}, err
	}
	humans, err := Repository.FindAllHumans()
	if err != nil {
		return &Stats{}, err
	}

	m := len(mutants)
	h := len(humans)
	ratio := 1.0
	if h != 0 {
		ratio = float64(m) / float64(h)
	}
	return NewStats(len(mutants), len(humans), ratio), nil
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
				err := errors.BadRequestError(fmt.Errorf("invalid nitrogen base found: %q", w))
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
