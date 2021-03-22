package mutant

import (
	"context"
	"github.com/Sebalvarez97/mutants-go/internal/domain/model"
	"github.com/stretchr/testify/mock"
)

type fakeMutantContainer struct {
	Container     Container
	DnaRepository *mockDnaRepository
}

func newFakeContainer() *fakeMutantContainer {
	fc := &fakeMutantContainer{
		DnaRepository: &mockDnaRepository{},
	}
	fc.Container = Container{
		DnaRepository: fc.DnaRepository,
	}
	return fc
}

type mockDnaRepository struct {
	mock.Mock
}

func (r *mockDnaRepository) FindByDnaHash(ctx context.Context, hash string) (*model.Dna, error) {
	args := r.Called(mock.Anything)
	dna, _ := args.Get(0).(*model.Dna)
	return dna, args.Error(1)
}
func (r *mockDnaRepository) Upsert(ctx context.Context, dna *model.Dna) error {
	args := r.Called(mock.Anything)
	return args.Error(0)
}

func (r *mockDnaRepository) FindNumberOfHumans(ctx context.Context) (int, error) {
	args := r.Called(mock.Anything)
	humans, _ := args.Get(0).(int)
	return humans, args.Error(1)
}
func (r *mockDnaRepository) FindNumberOfMutants(ctx context.Context) (int, error) {
	args := r.Called(mock.Anything)
	mutants, _ := args.Get(0).(int)
	return mutants, args.Error(1)
}
