package writer

import (
	"github.com/Sebalvarez97/mutants-go/internal/http"
	"github.com/gin-gonic/gin"
)

type routerHandler struct {
	handler MutantWriterHandler
}

func (r *routerHandler) RouteURLs(e *gin.Engine) {
	e.POST("/mutant", r.handler.IsMutantHandler)
}

func NewRouterHandler(handler MutantWriterHandler) http.RouterHandler {
	return &routerHandler{
		handler: handler,
	}
}
