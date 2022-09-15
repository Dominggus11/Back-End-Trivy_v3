package router

import (
	"trivy_v3/controllers"
	"trivy_v3/models"

	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()
	models.ConnectDatabase()
	// untuk API Upload
	r.GET("/getUploads", controllers.FindAllDocker)
	r.POST("/postUpload", controllers.PostDockerfile)
	r.GET("/getUpload/:id", controllers.FindDocker)
	r.PUT("/updateUpload/:id", controllers.UpdateDocker)
	r.DELETE("/deleteUpload/:id", controllers.DeleteDocker)

	// untuk API Project
	r.GET("/getProkjects", controllers.FindAllProject)
	r.GET("/getProject", controllers.FindProject)
	r.POST("/postProject", controllers.PostProject)
	r.PUT("/updateProject", controllers.UpdateProject)
	r.DELETE("/deleteProject", controllers.DeleteProject)
	r.Run()

	// untuk API return JSON
	r.GET("/jsonfile/:id", controllers.GetJson)
}
