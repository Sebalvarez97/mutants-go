package interfaces

import (
	"github.com/Sebalvarez97/mutants/api/errors"
	"github.com/Sebalvarez97/mutants/api/model"
)

type MutantService interface {
	IsMutant(dnaRequest model.IsMutantRequestBody) bool
	GetMutantStats() (*model.Stats, *errors.ApiErrorImpl)
}

type CerebroService interface {
	IsMutantDna(input [][]byte) (bool, int)
}
