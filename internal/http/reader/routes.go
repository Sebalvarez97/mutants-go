package reader

import (
	"github.com/Sebalvarez97/mutants-go/internal/http"
	"github.com/gin-gonic/gin"
)

type routerHandler struct {
	handler MutantReaderHandler
}

func (r *routerHandler) RouteURLs(e *gin.Engine) {
	e.GET("/mutant/stats", r.handler.GetStatsHandler)
}

func NewRouterHandler(handler MutantReaderHandler) http.RouterHandler {
	return &routerHandler{
		handler: handler,
	}
}
