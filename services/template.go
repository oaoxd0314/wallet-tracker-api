package services

import (
	"net/http"
	"wallet-tracker-api/models"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	data := new(models.IndexData)
	data.Title = "Wallet Tracker API"
	data.Content = ""
	c.HTML(http.StatusOK, "index.html", data)
}
