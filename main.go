package main

import (
	"os"

	"wallet-tracker-api/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	router := gin.Default()
	router.LoadHTMLGlob("template/*")
	router.GET("/", services.Test)
	router.POST("/getwalletfound", services.GetWalletFound)
	router.POST("/gettraderecord", services.GetTradeRecord)
	router.Run(":" + os.Getenv("PORT"))
}
