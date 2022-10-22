package prices

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type RespAPI struct {
	Bitcoin struct {
		USD    float64 `json:"usd"`
		Latest int64 `json:"last_updated_at"`
	} `json:"bitcoin"`
}

var url string = "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd&include_last_updated_at=true"

func GetPrice() (price float64) {
	cli := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := cli.Do(req)
	if err != nil {
		log.Println(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
	}
	var priceResp RespAPI
	json.Unmarshal(body, &priceResp)
	return priceResp.Bitcoin.USD
}
