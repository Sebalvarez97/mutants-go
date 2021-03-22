package http

import (
	"github.com/Sebalvarez97/mutants-go/tools/web"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RouterHandler interface {
	RouteURLs(r *gin.Engine)
}

func Response(w http.ResponseWriter, method string, body interface{}, err error) error {
	if err != nil {
		return err
	}
	if method == http.MethodPost {
		return web.RespondJSON(w, body, http.StatusCreated)
	}
	return web.RespondJSON(w, body, http.StatusOK)
}
