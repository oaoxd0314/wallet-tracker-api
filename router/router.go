package router

import (
	"fmt"
	"os"
	"wallet-tracker-api/services"

	"github.com/gin-gonic/gin"

	_ "wallet-tracker-api/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	PORT := os.Getenv("PORT")                                                        //get env var
	url := ginSwagger.URL(fmt.Sprintf("http://localhost:%v/swagger/doc.json", PORT)) //set swagger default url

	router.LoadHTMLGlob("template/*") // load template

	router.GET("/", services.HomePage) // set homepage

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url)) // set swagger path

	router.Group("etherscan"). // set 3rd part api(etherscan) as a group
					POST("/getwalletfound", services.GetWalletFound).
					POST("/gettraderecord", services.GetTradeRecord)

	return router
}
