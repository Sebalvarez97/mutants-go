package writer

import (
	"context"
	"fmt"
	"github.com/Sebalvarez97/mutants-go/errors"
	"github.com/Sebalvarez97/mutants-go/internal/domain/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MutantWriterHandler interface {
	IsMutantHandler(ctx *gin.Context)
}

type MutantService interface {
	IsMutant(ctx context.Context, dnaRequest model.IsMutantRequestBody) (bool, error)
}

type mutantWriterHandler struct {
	mutantSrv MutantService
}

func NewMutantWriterHandler(service MutantService) MutantWriterHandler {
	return &mutantWriterHandler{}
}

func (m *mutantWriterHandler) IsMutantHandler(ctx *gin.Context) {
	var json model.IsMutantRequestBody
	if err := ctx.BindJSON(&json); err != nil {
		apiErr := errors.BadRequestError(err)
		ctx.JSON(apiErr.Code, apiErr)
	}
	if valid, message := json.IsValid(); !valid {
		apiErr := errors.BadRequestError(fmt.Errorf(message))
		ctx.JSON(apiErr.Code, apiErr)
	}
	is, err := m.mutantSrv.IsMutant(ctx, json)
	if err != nil {
		if apiError, ok := err.(errors.ApiError); ok {
			ctx.JSON(apiError.Code, apiError)
		}
		ctx.JSON(http.StatusInternalServerError, err)
	}
	if is {
		ctx.Status(http.StatusOK)
	} else {
		ctx.Status(http.StatusForbidden)
	}
}
