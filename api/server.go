package main

import (
	"github.com/Sebalvarez97/mutants/api/auth"
	bundle "github.com/Sebalvarez97/mutants/api/controller"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Printf("Setting up server...\n")

	port := os.Getenv("PORT")
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	if port == "" {
		port = "8080"
	}

	log.Printf("Running on port: %v\n", port)

	authMiddleware, err := auth.GetAuthMiddleware()

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	r.Group("/auth").POST("/login", authMiddleware.LoginHandler).GET("/refresh_token", authMiddleware.RefreshHandler)

	mutant := r.Group("/mutant")
	mutant.Use(authMiddleware.MiddlewareFunc())
	{
		mutant.POST("", bundle.IsMutantHandler)
		mutant.GET("/stats", bundle.GetStatsHandler)
	}

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
