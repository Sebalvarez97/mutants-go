package reader

import (
	"context"
	"github.com/Sebalvarez97/mutants-go/internal/domain/model"
	"github.com/stretchr/testify/mock"
)

type fakeMutantReaderHandler struct {
	MutantReaderHandler MutantReaderHandler
	MutantSrv           *mockMutantService
}

func newFakeMutantReaderHandler() *fakeMutantReaderHandler {
	frh := &fakeMutantReaderHandler{
		MutantSrv: &mockMutantService{},
	}
	frh.MutantReaderHandler = &mutantReaderHandler{
		mutantSrv: frh.MutantSrv,
	}
	return frh
}

type mockMutantService struct {
	mock.Mock
}

func (s *mockMutantService) GetMutantStats(ctx context.Context) (*model.Stats, error) {
	args := s.Called(mock.Anything)
	stats, _ := args.Get(0).(*model.Stats)
	return stats, args.Error(1)
}
