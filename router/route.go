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
	r.PUT("/update/:id", controllers.Update)
	r.DELETE("/delete/:id", controllers.Delete)
	r.GET("/jsonfile/:id", controllers.GetJson)
	r.POST("/create", controllers.PostProject)
	r.Run()
}
