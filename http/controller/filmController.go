package controller

import(
	"github.com/gin-gonic/gin"
	"github.com/jameshudson/vhs_only/repositories"
)

type FilmContoller struct {
	FilmsRepositories *repositories.FilmRepositories
}

func (fc *FilmContoller) Index(c *gin.Context) {

	c.JSON(200, gin.H{
		"films": "hello",
	})
}

func (fc *FilmContoller) Get(c *gin.Context) {

	
}
