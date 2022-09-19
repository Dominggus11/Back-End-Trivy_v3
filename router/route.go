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
	r.GET("/Projects", controllers.FindProjects)
	r.GET("/Project/:id", controllers.FindProject)
	r.POST("/Project", controllers.PostProject)
	r.PUT("/Project/:id", controllers.UpdateProject)
	r.DELETE("/Project/:id", controllers.DeleteProject)

	// untuk API Upload
	r.GET("/Uploads", controllers.FindDockers)
	r.POST("/Upload", controllers.PostDockerfile)
	r.GET("/Upload/:id", controllers.FindDocker)
	r.PUT("/Upload/:id", controllers.UpdateDocker)
	r.DELETE("/Upload/:id", controllers.DeleteDocker)

	// untuk API Code
	// r.GET("/getCodes", controllers.FindCodes)
	r.POST("/Code", controllers.PostCode)
	r.GET("/Code/:id", controllers.FindCode)
	r.PUT("/Code/:id", controllers.UpdateCode)
	r.DELETE("/Code/:id", controllers.DeleteCode)

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
