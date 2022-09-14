package router

import (
	"trivy_v3/controllers"

	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()

	r.GET("/getAll", controllers.FindAll)
	r.POST("/upload", controllers.PostDockerfile)
}
