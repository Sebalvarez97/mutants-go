package main

import (
	"github.com/Sebalvarez97/mutants-go/config"
	"github.com/Sebalvarez97/mutants-go/internal/domain/mutant"
	"github.com/Sebalvarez97/mutants-go/internal/http/all"
	"github.com/Sebalvarez97/mutants-go/internal/http/reader"
	"github.com/Sebalvarez97/mutants-go/internal/http/writer"
	"github.com/Sebalvarez97/mutants-go/internal/middleware"
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
	services := make(map[string]interface{}, 0)
	conf := config.GetConfig()

	mutantService := mutant.NewService(mutant.InitializeContainer(conf))
	services[mutant.ServiceName] = mutantService

	readerHandler := reader.NewMutantReaderHandler(mutantService)
	writerHandler := writer.NewMutantWriterHandler(mutantService)

	routerHandler := all.NewRouterHandler(readerHandler, writerHandler)

	log.Printf("Setting up server...\n")
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	log.Printf("Setting up auth...\n")

	authMiddleWare := middleware.GetAuthMiddleWare(r)
	r.Use(authMiddleWare)

	errorHandler := middleware.GetCustomErrorHandler()
	r.Use(errorHandler)

	routerHandler.RouteURLs(r)

	log.Printf("Will run on port: %v\n", port)

	if err := http.ListenAndServe(":"+port, r); err != nil {
		return err
	}
	return nil
}
