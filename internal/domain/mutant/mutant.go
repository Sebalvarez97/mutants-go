package mutant

/*
import (
	"github.com/Sebalvarez97/mutants/api/errors"
	"github.com/Sebalvarez97/mutants/api/interfaces"
	"github.com/Sebalvarez97/mutants/api/model"
	"github.com/Sebalvarez97/mutants/api/util/matrix"
	"math"
)

type DnaRepository interface {
	Upsert(dna *model.Dna)
	FindAllMutants() ([]model.Dna, *errors.ApiErrorImpl)
	FindNumberOfHumans() (int, *errors.ApiErrorImpl)
	FindNumberOfMutants() (int, *errors.ApiErrorImpl)
	FindAllHumans() ([]model.Dna, *errors.ApiErrorImpl)
}

type MutantService interface {
	IsMutant(dnaRequest model.IsMutantRequestBody) bool
	GetMutantStats() (*model.Stats, *errors.ApiErrorImpl)
}

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

	mutants, humans, e := make(chan int), make(chan int), make(chan errors.ApiErrorImpl)

	go func(chan int, chan errors.ApiErrorImpl) {
		m, err := i.repository.FindNumberOfMutants()
		if err != nil {
			e <- *err
		}
		mutants <- m
	}(mutants, e)

	go func(chan int, chan errors.ApiErrorImpl) {
		h, err := i.repository.FindNumberOfHumans()
		if err != nil {
			e <- *err
		}
		humans <- h
	}(humans, e)

	select {
	case err := <-e:
		return &model.Stats{}, &err
	default:
		m := <-mutants
		h := <-humans
		ratio := 1.0
		if h != 0 {
			ratio = math.Round((float64(m)/float64(h))*100) / 100
		}
		return model.NewStats(m, h, ratio), nil
	}
}
 */

