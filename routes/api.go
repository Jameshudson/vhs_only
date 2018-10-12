package routes

import (
	"github.com/gin-gonic/gin"
	// "github.com/joho/godotenv/cmd/godotenv"
	"github.com/jameshudson/vhs_only/http/controller"
	// "fmt"
)

func ApiRoutes(routes *gin.Engine) {

	// filmController := new(controller.FilmContoller)

	routes.GET("/films", controller.Index)
	routes.GET("/films/:title", controller.Get)

}