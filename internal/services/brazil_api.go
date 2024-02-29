package services

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type BrazilApiCEP struct {
	Cep      string `json:"cep"`
	City     string `json:"city"`
	Location struct {
		Coordinates struct {
			Latitude  string `json:"latitude"`
			Longitude string `json:"longitude"`
		} `json:"coordinates"`
		Type string `json:"type"`
	} `json:"location"`
	Neighborhood string `json:"neighborhood"`
	Service      string `json:"service"`
	State        string `json:"state"`
	Street       string `json:"street"`
}

func GetCepInBrazilAPI() (*BrazilApiCEP, error) {
	res, err := http.Get("https://brasilapi.com.br/api/cep/v1/01153000")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(string(body))

	var cep BrazilApiCEP
	err = json.Unmarshal(body, &cep)
	return &cep, err

}
