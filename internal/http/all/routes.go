package all

import (
	//"log"
	//"net/http"
	//"github.com/gin-gonic/gin"
)

type routerHandler struct {
}
/*
func NewRouterHandler() http.RouterHandler {
	return &routerHandler{
	}
}
 */


func (r *routerHandler) API() error {
	return nil
}

func (r *routerHandler) RouteURLs() {
	//Mutant


	/*
	mutantController := controller.NewMutantController(mutantService)


		mutant := r.Group("/mutant")
		mutant.Use(authMiddleWare)
		{
			mutant.POST("", mutantController.IsMutantHandler)
			mutant.GET("/stats", mutantController.GetStatsHandler)
		}
		log.Printf("Will run on port: %v\n", port)

		app.Router.Method("GET", "/ceco", r.rh.SearchCeCo)
	 */

}
