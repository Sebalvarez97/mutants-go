package main

import (
	"github.com/Sebalvarez97/mutants/api/auth"
	"github.com/Sebalvarez97/mutants/api/controller"
	dao2 "github.com/Sebalvarez97/mutants/api/dao"
	"github.com/Sebalvarez97/mutants/api/repository"
	"github.com/Sebalvarez97/mutants/api/service"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {

}

func setUpServer() {
	log.Printf("Setting up server...\n")

	port := os.Getenv("PORT")
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	if port == "" {
		port = "8080"
	}

	authMiddleWare := setUpAuth(r)

	log.Println("Setting up services")
	dao := dao2.NewMongoDao("mutants")

	dnaRepository := repository.NewDnaRepository(dao)
	cerebroService := service.NewCerebroService()

	mutantService := service.NewMutantService(dnaRepository, cerebroService)
	mutantController := controller.NewMutantController(mutantService)

	mutant := r.Group("/mutant")
	mutant.Use(authMiddleWare)
	{
		mutant.POST("", mutantController.IsMutantHandler)
		mutant.GET("/stats", mutantController.GetStatsHandler)
	}
	log.Printf("Will run on port: %v\n", port)

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}

func setUpAuth(r *gin.Engine) gin.HandlerFunc {
	log.Println("Setting up auth")
	authMiddleware, err := auth.GetAuthMiddleware()

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	authMiddlewareFunc := authMiddleware.MiddlewareFunc()

	r.NoRoute(authMiddlewareFunc, func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	r.Group("/auth").POST("/login", authMiddleware.LoginHandler).GET("/refresh_token", authMiddleware.RefreshHandler)

	return authMiddlewareFunc
}
