package service

import (
	"github.com/Sebalvarez97/mutants/api/errors"
	"github.com/Sebalvarez97/mutants/api/interfaces"
	"github.com/Sebalvarez97/mutants/api/model"
	"github.com/Sebalvarez97/mutants/api/util/matrix"
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

func (i MutantServiceImpl) IsMutant(dnaRequest model.IsMutantRequestBody) bool {
	chain := dnaRequest.Dna
	h, mutant := make(chan string), make(chan bool)

	go func([]string, chan string) {
		h <- matrix.GenerateHashForStringArray(chain)
	}(chain, h)

	go func([]string, chan bool, chan string) {
		dna := make([][]byte, len(chain))
		for i, v := range chain {
			dna[i] = []byte(v)
		}

		is, s := i.cerebro.IsMutantDna(dna)
		mutant <- is

		i.repository.Upsert(model.NewDna(<-h, is, s))
	}(chain, mutant, h)

	mut := <-mutant

	return mut
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
