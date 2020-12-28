package controller

import (
	. "github.com/Sebalvarez97/mutants/src/app/bundle/mutant/service"
	errors2 "github.com/Sebalvarez97/mutants/src/app/common/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IsMutantRequestBody struct {
	Dna []string `form:"dna" json:"dna" binding:"required"`
}

type MutantService interface {
	IsMutant(input []string) (bool, *errors2.ApiErrorImpl)
}

var service MutantService = MutantServiceImpl{}

func IsMutantHandler(ctx *gin.Context) {
	var json IsMutantRequestBody
	if err := ctx.ShouldBind(&json); err != nil {
		apiErr := errors2.BadRequestError(err)
		ctx.JSON(apiErr.Code, apiErr)
	}
	is, apiErr := service.IsMutant(json.Dna)
	if apiErr != nil {
		ctx.JSON(apiErr.Code, apiErr)
	} else if is {
		ctx.Status(http.StatusOK)
	} else {
		ctx.Status(http.StatusForbidden)
	}

}
