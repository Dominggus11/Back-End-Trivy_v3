package router

import (
	"trivy_v3/controllers"
	"trivy_v3/models"

	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()
	models.ConnectDatabase()
	r.GET("/", controllers.HelloUser)
	// untuk API Project
	r.GET("/getProjects", controllers.FindProjects)
	r.GET("/getProject/:id", controllers.FindProject)
	r.POST("/postProject", controllers.PostProject)
	r.PUT("/updateProject/:id", controllers.UpdateProject)
	r.DELETE("/deleteProject/:id", controllers.DeleteProject)

	// untuk API Upload
	r.GET("/getUploads", controllers.FindDockers)
	r.POST("/postUpload", controllers.PostDockerfile)
	r.GET("/getUpload/:id", controllers.FindDocker)
	r.PUT("/updateUpload/:id", controllers.UpdateDocker)
	r.DELETE("/deleteUpload/:id", controllers.DeleteDocker)

	// untuk API Code
	r.GET("/getCodes", controllers.FindCodes)
	r.POST("/postCode", controllers.PostCode)
	r.GET("/getCode/:id", controllers.FindCode)
	r.PUT("/updateCode/:id", controllers.UpdateCode)
	r.DELETE("/deleteCode/:id", controllers.DeleteCode)

	// untuk API return JSON
	r.GET("/jsonfile/:id", controllers.GetJson)

	r.Run()
}
