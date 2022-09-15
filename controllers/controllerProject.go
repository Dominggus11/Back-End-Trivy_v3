package controllers

import (
	"net/http"
	"trivy_v3/models"

	"github.com/gin-gonic/gin"
)

func FindAllProject(c *gin.Context) {

}

func PostProject(c *gin.Context) {
	//db := models.DB
	var input models.Project
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	//create Dockerfile
	project := models.Project{
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
