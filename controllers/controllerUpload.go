package controllers

import (
	"net/http"
	"trivy_v3/models"

	"github.com/gin-gonic/gin"
)

func FindAll(c *gin.Context) {
	var dockerfiles []models.Dockerfiles
	models.DB.Find(&dockerfiles)

	c.JSON(http.StatusOK, gin.H{"data": dockerfiles})
}

func PostDockerfile(c *gin.Context) {
	//db := c.MustGet("db").(*gorm.DB)
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
	models.DB.Create(&dockerfile)
	c.JSON(http.StatusOK, gin.H{
		"data": dockerfile,
	})
}

func Find(c *gin.Context) {
	db := models.DB
	// Get model if exist
	var dockerfile models.Dockerfiles
	if err := db.Where("id = ?", c.Param("id")).First(&dockerfile).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": dockerfile})
}
