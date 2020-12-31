package controller

import (
	. "github.com/Sebalvarez97/mutants/api/model"
	"github.com/Sebalvarez97/mutants/api/errors"
	. "github.com/Sebalvarez97/mutants/api/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IsMutantRequestBody struct {
	Dna []string `form:"dna" json:"dna" binding:"required"`
}

type MutantService interface {
	IsMutant(input []string) (bool, *errors.ApiErrorImpl)
	GetMutantStats() (*Stats, *errors.ApiErrorImpl)
}

type MutantServiceImpl struct{}

func (i MutantServiceImpl) IsMutant(input []string) (bool, *errors.ApiErrorImpl) {
	return IsMutant(input)
}

func (i MutantServiceImpl) GetMutantStats() (*Stats, *errors.ApiErrorImpl) {
	return GetMutantStats()
}

var Service MutantService

func init() {
	Service = MutantServiceImpl{}
}

func IsMutantHandler(ctx *gin.Context) {
	var json IsMutantRequestBody
	if err := ctx.BindJSON(&json); err != nil {
		apiErr := errors.BadRequestError(err)
		ctx.JSON(apiErr.Code, apiErr)
	}
	is, apiErr := Service.IsMutant(json.Dna)
	if apiErr != nil {
		ctx.JSON(apiErr.Code, apiErr)
	} else if is {
		ctx.Status(http.StatusOK)
	} else {
		ctx.Status(http.StatusForbidden)
	}
}

func GetStatsHandler(ctx *gin.Context) {
	stats, apiErr := Service.GetMutantStats()
	if apiErr != nil {
		ctx.JSON(apiErr.Code, apiErr)
	}
	ctx.JSON(http.StatusOK, stats)
}
