package router

import (
	"trivy_v3/controllers"
	"trivy_v3/models"

	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()
	r.Use(CORS)
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
	// r.GET("/getCodes", controllers.FindCodes)
	r.POST("/postCode", controllers.PostCode)
	r.GET("/getCode/:id", controllers.FindCode)
	r.PUT("/updateCode/:id", controllers.UpdateCode)
	r.DELETE("/deleteCode/:id", controllers.DeleteCode)

	// untuk API return JSON
	r.GET("/jsonfile/:id", controllers.GetJson)

	r.Run()
}

func CORS(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, "+
		"Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
}
