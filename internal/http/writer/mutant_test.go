package writer

import (
	"context"
	"github.com/Sebalvarez97/mutants-go/internal/domain/model"
	"github.com/stretchr/testify/mock"
)

type fakeMutantWriterHandler struct {
	MutantWriterHandler MutantWriterHandler
	MutantSrv           *mockMutantService
}

func newFakeMutantWriterHandler() *fakeMutantWriterHandler {
	frh := &fakeMutantWriterHandler{
		MutantSrv: &mockMutantService{},
	}
	frh.MutantWriterHandler = &mutantWriterHandler{
		mutantSrv: frh.MutantSrv,
	}
	return frh
}

type mockMutantService struct {
	mock.Mock
}

func (s *mockMutantService) IsMutant(ctx context.Context, dnaRequest model.IsMutantRequestBody) (bool, error) {
	args := s.Called(mock.Anything)
	is, _ := args.Get(0).(bool)
	return is, args.Error(1)
}
