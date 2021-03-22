package mutant

import (
	"context"
	"github.com/Sebalvarez97/mutants-go/internal/domain/cerebro"
	"github.com/Sebalvarez97/mutants-go/internal/domain/model"
	"github.com/Sebalvarez97/mutants/api/util/matrix"
	"golang.org/x/sync/errgroup"
	"math"
	"sync"
)

const (
	ServiceName = "mutantService"
)

type Service interface {
	GetMutantStats(ctx context.Context) (*model.Stats, error)
	IsMutant(ctx context.Context, dnaRequest model.IsMutantRequestBody) (bool, error)
}

type service struct {
	container Container
}

func NewService(container Container) Service {
	return &service{container: container}
}

func (s *service) GetMutantStats(ctx context.Context) (*model.Stats, error) {
	var mutants, humans int
	var g errgroup.Group
	g.Go(func() error {
		m, err := s.container.DnaRepository.FindNumberOfMutants(ctx)
		mutants = m
		return err
	})
	g.Go(func() error {
		h, err := s.container.DnaRepository.FindNumberOfHumans(ctx)
		humans = h
		return err
	})
	if err := g.Wait(); err != nil {
		return &model.Stats{}, err
	}
	ratio := 1.0
	if humans != 0 {
		ratio = math.Round((float64(mutants)/float64(humans))*100) / 100
	}
	return model.NewStats(mutants, humans, ratio), nil
}

func (s *service) IsMutant(ctx context.Context, dnaRequest model.IsMutantRequestBody) (bool, error) {
	var hash string
	var mutant bool
	var found int
	var wg sync.WaitGroup
	chain := dnaRequest.Dna
	wg.Add(1)
	go func() {
		defer wg.Done()
		hash = matrix.GenerateHashForStringArray(chain)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		dna := make([][]byte, len(chain))
		for i, v := range chain {
			dna[i] = []byte(v)
		}
		mutant, found = cerebro.IsMutant(dna)
	}()
	wg.Wait()
	go s.container.DnaRepository.Upsert(ctx, model.NewDna(hash, mutant, found))
	return mutant, nil
}
