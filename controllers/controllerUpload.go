package controllers

import (
	"net/http"
	"trivy_v3/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FindAll(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var dockerfiles []models.Dockerfiles
	db.Find(&dockerfiles)

	c.JSON(http.StatusOK, gin.H{"data": dockerfiles})
}

func PostDockerfile(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input models.Dockerfiles
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//create Dockerfile
	dockerfile := models.Dockerfiles{
		Pathfile: input.Pathfile,
		PathJson: input.PathJson,
	}
	db.Create(&dockerfile)
	c.JSON(http.StatusOK, gin.H{
		"data": dockerfile,
	})
}
