package main

import (
	"khanek/exercise-generator/api"
	"khanek/exercise-generator/core/db"
	"log"

	"github.com/gin-gonic/gin"

	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func main() {
	database := db.DB()
	defer database.Close()

	g.Go(initializeAPI)

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}

func initializeAPI() error {
	r := gin.Default()
	api.AddUrls(r)
	return r.Run(":9000")
}
