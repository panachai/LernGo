package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type BxModel struct {
	PairingID         int     `json:"pairing_id"`
	PrimaryCurrency   string  `json:"primary_currency"`
	SecondaryCurrency string  `json:"secondary_currency"`
	Change            float64 `json:"change"`
	LastPrice         int     `json:"last_price"`
	Volume24Hours     float64 `json:"volume_24hours"`
	Orderbook         struct {
		Bids struct {
			Total   int     `json:"total"`
			Volume  float64 `json:"volume"`
			Highbid int     `json:"highbid"`
		} `json:"bids"`
		Asks struct {
			Total   int     `json:"total"`
			Volume  float64 `json:"volume"`
			Highbid int     `json:"highbid"`
		} `json:"asks"`
	} `json:"orderbook"`
}

func main() {
	// token := "Pd8OH3yI8quQyZyc4Xpe5wu3Kr1UOlz3MM4t32HwEeK"

	// c := linenotify.New()
	// c.Notify(token, "ทดสอบ Realize on Golang 8", "", "", nil)

	GetBxData()
	// CallnormalWay()
}

// func toJson(){
// 	jsonString, err := json.Marshal(datas)
// }

func CallnormalWay() {
	urlTest := "https://bx.in.th/api/"

	response, err := http.Get(urlTest)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		fmt.Println(string(data))
	}
}

func GetBxData() {
	url := "https://bx.in.th/api/"

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	// fmt.Println(string(data))

	var bx BxModel

	err := json.Unmarshal(data, &bx)
	// defer resp.Body.Close()

	var record BxModel

	if err := json.NewDecoder(response.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	fmt.Println("Phone No. = ", record.PairingID)
	fmt.Println("Country   = ", record.PrimaryCurrency)
	fmt.Println("Location  = ", record.SecondaryCurrency)
	fmt.Println("Carrier   = ", record.Change)
	fmt.Println("LineType  = ", record.LastPrice)
}
