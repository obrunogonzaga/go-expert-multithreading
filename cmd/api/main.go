package main

import (
	"github.com/obrunogonzaga/go-expert-multithreading/internal/services"
	"log"
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

func main() {
	cep, err := services.GetCepInBrazilAPI()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(cep)
}
