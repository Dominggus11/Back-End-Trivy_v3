package router

import (
	"trivy_v3/controllers"
	"trivy_v3/models"

	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/getAll", controllers.FindAll)
	r.POST("/upload", controllers.PostDockerfile)
	r.GET("/get/:id", controllers.Find)
	r.Run()
}
