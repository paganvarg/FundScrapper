package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type yahooResponse struct {
	Quotes []struct {
		Symbol         string `json:"symbol"`
		Longname       string `json:"longname"`
		IsYahooFinance bool   `json:"isYahooFinance"`
	} `json:"quotes"`
}

func main() {
	var fundIsins = []string{"GB00B8L3WZ29", "GB00B7VHZX64", "GB00B8H99P30", "GB00BMMV5105", "GB00B28BBW75", "GB00B5LXGG05", "GB00BMMV5766", "GB00BVZ6VF19"}

	for _, isin := range fundIsins {
		var url = fmt.Sprintf("https://query2.finance.yahoo.com/v1/finance/search?q=%s", isin)
		resp, err := http.Get(url)
		if err != nil {
			log.Print(err)
			continue
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Print(err)
			continue
		}

		var yahooResponse yahooResponse
		err = json.Unmarshal(body, &yahooResponse)
		if err != nil {
			log.Print(err)
			continue
		}

		if len(yahooResponse.Quotes) == 0 {
			log.Print("Not valid response from Yahoo")
			continue
		}

		fmt.Printf("new Fund(\"%s\", \"%s\", %t, \"%s\"),\r\n", isin, yahooResponse.Quotes[0].Longname, yahooResponse.Quotes[0].IsYahooFinance, yahooResponse.Quotes[0].Symbol)
	}
}
