package services

import (
	"net/http"
	"wallet-tracker-api/models"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	data := new(models.IndexData)
	data.Title = "首頁"
	data.Content = "我的第一個首頁"
	c.HTML(http.StatusOK, "index.html", data)
}
