package services

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type ViaCepApi struct {
	Bairro      string `json:"bairro"`
	Cep         string `json:"cep"`
	Complemento string `json:"complemento"`
	Ddd         string `json:"ddd"`
	Gia         string `json:"gia"`
	Ibge        string `json:"ibge"`
	Localidade  string `json:"localidade"`
	Logradouro  string `json:"logradouro"`
	Siafi       string `json:"siafi"`
	Uf          string `json:"uf"`
}

func GetCepInViaCepAPI() (*ViaCepApi, error) {
	res, err := http.Get("https://viacep.com.br/ws/80320310/json/")
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

	var cep ViaCepApi
	err = json.Unmarshal(body, &cep)
	return &cep, err
}
