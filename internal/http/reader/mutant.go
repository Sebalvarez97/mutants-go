package reader

import (
	"context"
	"github.com/Sebalvarez97/mutants-go/internal/domain/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MutantReaderHandler interface {
	GetStatsHandler(ctx *gin.Context)
}

type MutantService interface {
	GetMutantStats(ctx context.Context) (*model.Stats, error)
}

type mutantReaderHandler struct {
	mutantSrv MutantService
}

func NewMutantReaderHandler(service MutantService) MutantReaderHandler {
	return &mutantReaderHandler{}
}

func (m *mutantReaderHandler) GetStatsHandler(ctx *gin.Context) {
	stats, err := m.mutantSrv.GetMutantStats(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, stats)
}
