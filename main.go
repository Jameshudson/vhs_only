package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jameshudson/vhs_only/routes"
	"github.com/jameshudson/vhs_only/http/controller"
	"github.com/jameshudson/vhs_only/repositories"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	filmController := new(controller.FilmContoller)
	filmRepositories := new(repositories.FilmRepositories)

	filmController.FilmsRepositories = filmRepositories

	r := gin.Default()

	routes.ApiRoutes(r, filmController)
	
	r.Run() // listen and serve on 0.0.0.0:8080
}

