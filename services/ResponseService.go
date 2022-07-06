package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"wallet-tracker-api/controller"
	"wallet-tracker-api/models"
)

func fetchEtherScan(endpoint string) (respData models.EtherScanResp) {
	apikey := os.Getenv("API_KEY")
	path := fmt.Sprintf("https://api.etherscan.io/%v&apikey=%v", endpoint, apikey)

	response, err := http.Get(path)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Can not Read reponse body")
	}

	if err := json.Unmarshal(body, &respData); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	fmt.Println(controller.PrettyPrint(respData))

	return respData
}
