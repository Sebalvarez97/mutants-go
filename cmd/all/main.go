package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"github.com/Sebalvarez97/mutants-go/config"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	log.Printf("Setting up server...\n")
	conf := config.GetConfig()

	return nil
}

func ginRouter(config.Gin){
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
}

func setUpServer() {

	/*
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
	 */

}
/*

 */

