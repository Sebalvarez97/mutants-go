package service

import errors "github.com/Sebalvarez97/mutants/app/common"

type MutantServiceImpl struct{}

type CerebroService interface {
	IsMutant(input []string) (bool, *errors.ApiErrorImpl)
}

var cerebro CerebroService = CerebroServiceImpl{}

func (i MutantServiceImpl) IsMutant(input []string) (bool, *errors.ApiErrorImpl) {
	return cerebro.IsMutant(input)
}
