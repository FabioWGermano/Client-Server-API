# go-expert-server

Projeto inicial contendo o primeiro exercicio de Go.
Nesse projeto buscamos informações da cotação atual e persistimos num banco SQlite.

## Tecnologias Utilizadas

- Go
- go-sqlite3

## Instalação

Baixar o projeto https://github.com/FabioWGermano/go-expert-server.git
Executar com o comando "go run cmd/main.go"
Apos inicio da execução o database será criado automaticamente com o inicio da aplicação
A busca deve ser executada através do seguinte cURL "curl --location 'http://localhost:8080/cotacao'"