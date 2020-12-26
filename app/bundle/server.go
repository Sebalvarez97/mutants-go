package main

import (
	"github.com/Sebalvarez97/mutants/app/auth"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	bundle "github.com/Sebalvarez97/mutants/app/bundle/mutant/controller"
)

func main() {
	port := os.Getenv("PORT")
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	if port == "" {
		port = "8080"
	}

	authMiddleware, err := auth.GetAuthMiddleware()

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth := r.Group("/auth")
	auth.POST("/login", authMiddleware.LoginHandler)
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)

	mutant := r.Group("/mutant")
	mutant.Use(authMiddleware.MiddlewareFunc())
	{
		mutant.POST("", bundle.IsMutantHandler)
	}

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
