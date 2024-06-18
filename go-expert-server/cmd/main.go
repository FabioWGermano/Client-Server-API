package main

import (
	"log"
	"net/http"

	controller "github.com/FabioWGermano/go-expert-server/internal/controller/action"
)

func main() {
	controller.CriarTabela()
	http.HandleFunc("/cotacao", controller.Handle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
