package interfaces

import (
	"github.com/Sebalvarez97/mutants/api/errors"
	"github.com/Sebalvarez97/mutants/api/model"
)

type DnaRepository interface {
	Upsert(dna *model.Dna)
	FindAllMutants() ([]model.Dna, *errors.ApiErrorImpl)
	FindNumberOfHumans() (int, *errors.ApiErrorImpl)
	FindNumberOfMutants() (int, *errors.ApiErrorImpl)
	FindAllHumans() ([]model.Dna, *errors.ApiErrorImpl)
}
