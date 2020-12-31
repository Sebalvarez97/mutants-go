package controller

import (
	"github.com/Sebalvarez97/mutants/api/errors"
	"github.com/Sebalvarez97/mutants/api/interfaces"
	. "github.com/Sebalvarez97/mutants/api/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MutantController struct {
	service interfaces.MutantService
}

func NewMutantController(service interfaces.MutantService) MutantController {
	return MutantController{service: service}
}

func (i MutantController) IsMutantHandler(ctx *gin.Context) {
	var json IsMutantRequestBody
	if err := ctx.BindJSON(&json); err != nil {
		apiErr := errors.BadRequestError(err)
		ctx.JSON(apiErr.Code, apiErr)
	}
	is, apiErr := i.service.IsMutant(json.Dna)
	if apiErr != nil {
		ctx.JSON(apiErr.Code, apiErr)
	} else if is {
		ctx.Status(http.StatusOK)
	} else {
		ctx.Status(http.StatusForbidden)
	}
}

func (i MutantController) GetStatsHandler(ctx *gin.Context) {
	stats, apiErr := i.service.GetMutantStats()
	if apiErr != nil {
		ctx.JSON(apiErr.Code, apiErr)
	}
	ctx.JSON(http.StatusOK, stats)
}
