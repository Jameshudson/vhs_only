package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jameshudson/vhs_only/routes"
	// "net/http"
)

func main() {
	r := gin.Default()

	routes.ApiRoutes(r)
	
	r.Run() // listen and serve on 0.0.0.0:8080
}

