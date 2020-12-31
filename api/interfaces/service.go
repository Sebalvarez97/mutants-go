package interfaces

import (
	"github.com/Sebalvarez97/mutants/api/errors"
	"github.com/Sebalvarez97/mutants/api/model"
)

type MutantService interface {
	IsMutant(input []string) (bool, *errors.ApiErrorImpl)
	GetMutantStats() (*model.Stats, *errors.ApiErrorImpl)
}

type CerebroService interface {
	IsMutantDna(input [][]byte) (bool, int)
}
