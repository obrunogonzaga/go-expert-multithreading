package main

import (
	"github.com/obrunogonzaga/go-expert-multithreading/internal/services"
	"log"
	"time"
)

type Address struct {
	cep         string
	logradouro  string
	complemento string
	bairro      string
	localidade  string
	uf          string
}

func main() {

	address := Address{}
	brazilApiChannel := make(chan Address)
	ViaCepChannel := make(chan Address)
	done := make(chan bool)
	//errChannel := make(chan error)

	go func() {
		cep, err := services.GetCepInBrazilAPI()
		if err != nil {
			log.Println(err)
			return
		}

		select {
		case <-done:
			return
		default:
			address.cep = cep.Cep
			address.logradouro = cep.Street
			address.bairro = cep.Neighborhood
			address.localidade = cep.City
			address.uf = cep.State

			brazilApiChannel <- address
		}
	}()

	go func() {
		viaCep, err := services.GetCepInViaCepAPI()
		if err != nil {
			log.Println(err)
			return
		}

		select {
		case <-done:
			return
		default:
			address.cep = viaCep.Cep
			address.logradouro = viaCep.Logradouro
			address.complemento = viaCep.Complemento
			address.bairro = viaCep.Bairro
			address.localidade = viaCep.Localidade
			address.uf = viaCep.Uf

			ViaCepChannel <- address
		}
	}()

	select {
	case address := <-brazilApiChannel:
		log.Printf("Received from BrazilAPI: %+v\n", address)
		close(done)

	case address := <-ViaCepChannel:
		log.Printf("Received from ViaCep: %+v\n", address)
		close(done)

	case <-time.After(1 * time.Second):
		println("Timeout BrazilAPI and ViaCep")
	}

}
