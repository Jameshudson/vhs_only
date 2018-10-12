package routes

import (
	"github.com/gin-gonic/gin"
	// "github.com/joho/godotenv/cmd/godotenv"
	"github.com/jameshudson/vhs_only/http/controller"
	// "github.com/jameshudson/vhs_only/repositories"
	// "fmt"
)

func ApiRoutes(routes *gin.Engine, filmContoller *controller.FilmContoller) {

	filmController := new(controller.FilmContoller)

	routes.GET("/films", filmController.Index)
	routes.GET("/films/:title", filmController.Get)
}