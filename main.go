package main

import (
	"trivy_v3/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase() // new
	r.Run()
}
