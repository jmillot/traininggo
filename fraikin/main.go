package main

import (
	"fmt"
	"traininggo/fraikin/webservices"
)

func main() {

	// Production
	url := "https://api.eu-de.apiconnect.appdomain.cloud/fraikin-prd/panier/ligne-commande/addToPanier"
	// url := "https://api.eu-de.apiconnect.ibmcloud.com/g006api-connectfraikincom-dev/sb/ligne-commande/addToPanier"
	headers := make(map[string]string)
	headers["x-ibm-client-id"] = "800e7648-c291-4040-9be1-57ac22e02785" // ID IBM Fraikin
	headers["accept"] = "application/json"

	ws := webservices.New(url, headers)
	err := ws.CallWebService()
	if err != nil {
		fmt.Println(err)
	}
}
