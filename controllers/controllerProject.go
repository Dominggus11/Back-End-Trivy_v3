package controllers

import (
	"net/http"
	"trivy_v3/models"

	"github.com/gin-gonic/gin"
)

func FindProjects(c *gin.Context) {
	var projects []models.Projects
	models.DB.Find(&projects)

	c.JSON(http.StatusOK, gin.H{"data": projects})
}

func PostProject(c *gin.Context) {
	//db := models.DB
	var input models.Projects
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	//create Dockerfile
	project := models.Projects{
		ProjectName: input.ProjectName,
	}
	models.DB.Create(&project)
	c.JSON(http.StatusOK, gin.H{
		"data": project,
	})
}

func FindProject(c *gin.Context) {

}

func UpdateProject(c *gin.Context) {

}

func DeleteProject(c *gin.Context) {

}
