package services

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"wallet-tracker-api/models"

	"github.com/gin-gonic/gin"
)

func GetWalletFound(c *gin.Context) {
	fmt.Println(os.Getenv("PORT"))
	var requestBody models.Wallet

	if err := c.BindJSON(&requestBody); err != nil {
		return
	}

	endpoint := fmt.Sprintf("api?module=account&action=balance&address=%v&tag=latest", requestBody.Wallet)

	var respData models.EtherScanResp = fetchEtherScan(endpoint)
	c.IndentedJSON(http.StatusCreated, respData)
}

func GetTradeRecord(c *gin.Context) {
	var requestBody models.Wallet

	if err := c.BindJSON(&requestBody); err != nil {
		return
	}
	endpoint := fmt.Sprintf("api?module=account&action=txlist&address=%v&startblock=0&endblock=99999999&page=1&offset=10&sort=asc", requestBody.Wallet)

	var respData models.EtherScanResp = fetchEtherScan(endpoint)
	c.IndentedJSON(http.StatusCreated, respData)
}

func GetMultiWalletFound(c *gin.Context) {

	var requestBody models.WalletArray

	if err := c.BindJSON(&requestBody); err != nil {
		log.Fatal(err)
		fmt.Println("Bind Json ERR")
		return
	}

	// walletStr := strings.Join(requestBody[:], ", ")

	// fmt.Println(walletStr)

	// endpoint := fmt.Sprintf("api?module=account&action=balancemulti&address=%v&tag=latest", walletStr)

	// var respData EtherScanResp = fetchEtherScan(endpoint)
	// c.IndentedJSON(http.StatusCreated, respData)
}
