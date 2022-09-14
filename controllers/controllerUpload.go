package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
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
	file, err := c.FormFile("pathfile")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}
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

func Update(c *gin.Context) {
	db := models.DB
	// Get model if exist
	var input models.Dockerfiles
	if err := db.Where("id = ?", c.Param("id")).First(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	pathFile := input.Pathfile
	pathJson := input.PathJson
	file, err := c.FormFile("pathfile")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}
	filename := file.Filename
	c.SaveUploadedFile(file, pathFile+"/"+file.Filename)
	trivy.TrivyScan(pathJson, pathFile, filename)
	dockerfile := models.Dockerfiles{
		Pathfile: pathFile,
		PathJson: pathJson,
	}
	//db.Updates(&dockerfile)
	db.Model(&input).Updates(dockerfile)
	c.JSON(http.StatusOK, gin.H{
		"data": dockerfile,
	})

}

func Delete(c *gin.Context) {
	db := models.DB
	// Get model if exist
	var input models.Dockerfiles
	if err := db.Where("id = ?", c.Param("id")).First(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	pathFile := input.Pathfile
	pathJson := input.PathJson

	os.RemoveAll(pathFile)
	fmt.Print(pathFile)
	os.RemoveAll(pathJson)
	db.Delete(&input)

	c.JSON(http.StatusOK, gin.H{
		"data": "Deleted",
	})
}

func GetJson(c *gin.Context) {
	db := models.DB
	// Get model if exist
	var input models.Dockerfiles
	if err := db.Where("id = ?", c.Param("id")).First(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	pathJson := input.PathJson + "/resultsImage.json"
	file, _ := os.Open(pathJson)
	defer file.Close()

	fileContent, _ := io.ReadAll(file)

	c.String(http.StatusOK, string(fileContent))
}
