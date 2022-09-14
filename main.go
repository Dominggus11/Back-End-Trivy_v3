package main

import (
	"trivy_v3/controllers"
	"trivy_v3/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase() // new
	// router.Router()
	r.GET("/getAll", controllers.FindAll)
	r.Run()
}
