package mutant

import (
	"context"
	"github.com/Sebalvarez97/mutants-go/errors"
	"github.com/Sebalvarez97/mutants-go/internal/domain/cerebro"
	"github.com/Sebalvarez97/mutants-go/internal/domain/model"
	"github.com/Sebalvarez97/mutants/api/util/matrix"
	"math"
)

const (
	ServiceName = "mutantService"
)

type Service interface {
	IsMutant(ctx context.Context, dnaRequest model.IsMutantRequestBody) (bool, error)
	GetMutantStats(ctx context.Context) (*model.Stats, error)
}

type service struct {
	container Container
}

func NewService(container Container) Service {
	return &service{container: container}
}

func (s *service) IsMutant(ctx context.Context, dnaRequest model.IsMutantRequestBody) (bool, error) {
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

		is, found := cerebro.IsMutant(dna)
		mutant <- is

		_ = s.container.DnaRepository.Upsert(ctx, model.NewDna(<-h, is, found))
	}(chain, mutant, h)

	mut := <-mutant

	return mut, nil
}

func (s *service) GetMutantStats(ctx context.Context) (*model.Stats, error) {
	mutants, humans, e := make(chan int), make(chan int), make(chan errors.ApiError)

	go func(chan int, chan errors.ApiError) {
		m, err := s.container.DnaRepository.FindNumberOfMutants(ctx)
		if err != nil {
			e <- err.(errors.ApiError)
		}
		mutants <- m
	}(mutants, e)

	go func(chan int, chan errors.ApiError) {
		h, err := s.container.DnaRepository.FindNumberOfHumans(ctx)
		if err != nil {
			e <- err.(errors.ApiError)
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
