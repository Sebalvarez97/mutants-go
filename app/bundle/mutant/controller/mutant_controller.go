package controller

import (
	. "github.com/Sebalvarez97/mutants/app/bundle/mutant/service"
	errors "github.com/Sebalvarez97/mutants/app/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IsMutantRequestBody struct {
	Dna []string `form:"dna" json:"dna" binding:"required"`
}

type MutantService interface {
	IsMutant(input []string) (bool, *errors.ApiErrorImpl)
}

var service MutantService = MutantServiceImpl{}

func IsMutantHandler(ctx *gin.Context) {
	var json IsMutantRequestBody
	if err := ctx.ShouldBind(&json); err != nil {
		apiErr := errors.BadRequestError(err)
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
