package http

import (
	"github.com/gin-gonic/gin"
)

type RouterHandler interface {
	RouteURLs(r *gin.Engine)
}
