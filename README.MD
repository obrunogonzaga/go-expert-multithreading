# Projeto Go Expert Multithreading

Este projeto é uma aplicação Go que utiliza multithreading para buscar informações de CEP de duas APIs diferentes: BrasilAPI e ViaCep.

## Descrição

O projeto consiste em duas partes principais:

1. `brazil_api.go`: Este arquivo contém a implementação para buscar informações de CEP da BrasilAPI.
2. `viacep_api.go`: Este arquivo contém a implementação para buscar informações de CEP da ViaCep.

O programa principal (`main.go`) inicia duas goroutines, cada uma fazendo uma solicitação a uma das APIs. A primeira resposta recebida é usada e a outra solicitação é cancelada.

## Instalação

Para instalar este projeto, siga estas etapas:

1. Clone o repositório para o seu ambiente local.
2. Navegue até o diretório do projeto.
3. Execute `go mod download` para baixar as dependências necessárias.

## Execução

Para executar este projeto, siga estas etapas:

1. Navegue até o diretório do projeto.
2. Execute `go run main.go`.

## Contribuição

Contribuições são bem-vindas! Por favor, leia as diretrizes de contribuição antes de enviar uma pull request.

## Licença

Este projeto está licenciado sob a licença MIT.