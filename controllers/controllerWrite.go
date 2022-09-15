package controllers

import (
	"net/http"
	"trivy_v3/models"
	"trivy_v3/trivy"

	"github.com/gin-gonic/gin"
)

func FindCodes(c *gin.Context) {

}

func FindCode(c *gin.Context) {

}

func PostCode(c *gin.Context) {
	//db := models.DB
	var input models.Dockerfiles
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	pathFile := trivy.MkdirWriteFile()
	pathJson := trivy.MkdirWriteJson()
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
		Pathfile:  pathFile,
		PathJson:  pathJson,
		ProjectID: input.ProjectID,
	}
	models.DB.Create(&dockerfile)
	c.JSON(http.StatusOK, gin.H{
		"data": dockerfile,
	})
}

func UpdateCode(c *gin.Context) {

}

func DeleteCode(c *gin.Context) {

}
