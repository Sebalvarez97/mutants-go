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
	if err := ctx.BindJSON(&json); err != nil {
		apiErr := errors.NewJsonError(err.Error())
		ctx.Error(apiErr)
		return
	} else {
		if valid, message := json.IsValid(); !valid {
			apiErr := errors.NewValidationError(message)
			ctx.Error(apiErr)
			return
		} else {
			is, err := m.mutantSrv.IsMutant(ctx, json)
			if err != nil {
				ctx.Error(err)
				return
			} else {
				if is {
					ctx.Status(http.StatusOK)
				} else {
					ctx.Status(http.StatusForbidden)
				}
			}
		}
	}
}
