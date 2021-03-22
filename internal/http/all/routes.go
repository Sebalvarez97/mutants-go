package all

import (
	"github.com/Sebalvarez97/mutants-go/internal/http"
	"github.com/Sebalvarez97/mutants-go/internal/http/reader"
	"github.com/Sebalvarez97/mutants-go/internal/http/writer"
	"github.com/gin-gonic/gin"
)

type routerHandler struct {
	wh writer.MutantWriterHandler
	rh reader.MutantReaderHandler
}

func (r routerHandler) RouteURLs(e *gin.Engine) {
	e.GET("/mutant/stats", r.rh.GetStatsHandler)
	e.POST("/mutant", r.wh.IsMutantHandler)
}

func NewRouterHandler(reader reader.MutantReaderHandler, writer writer.MutantWriterHandler) http.RouterHandler {
	return &routerHandler{
		wh: writer,
		rh: reader,
	}
}
