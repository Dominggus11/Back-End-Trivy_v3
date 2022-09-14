package controllers

import (
	"net/http"
	"trivy_v3/models"
	"trivy_v3/trivy"

	"github.com/gin-gonic/gin"
)

func FindAll(c *gin.Context) {
	var dockerfiles []models.Dockerfiles
	models.DB.Find(&dockerfiles)

	c.JSON(http.StatusOK, gin.H{"data": dockerfiles})
}

func PostDockerfile(c *gin.Context) {
	//db := models.DB
	var input models.Dockerfiles
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	pathFile := trivy.MkdirUploadFile()
	pathJson := trivy.MkdirUploadJson()
	file, _ := c.FormFile("pathfile")
	input.Pathfile = file.Filename
	c.SaveUploadedFile(file, pathFile+"/"+file.Filename)
	trivy.TrivyScan(pathJson, pathFile, input.Pathfile)

	//create Dockerfile
	dockerfile := models.Dockerfiles{
		Pathfile: pathFile,
		PathJson: pathJson,
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
