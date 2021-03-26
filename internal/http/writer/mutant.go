package writer

import (
	"context"
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
	if err := ctx.ShouldBindJSON(&json); err != nil {
		apiErr := errors.NewJsonError(err.Error())
		_ = ctx.Error(apiErr)
		return
	}
	if valid, message := json.IsValid(); !valid {
		apiErr := errors.NewValidationError(message)
		_ = ctx.Error(apiErr)
		return
	}
	is, err := m.mutantSrv.IsMutant(ctx, json)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	if is {
		ctx.Status(http.StatusOK)
	} else {
		ctx.Status(http.StatusForbidden)
	}
}
