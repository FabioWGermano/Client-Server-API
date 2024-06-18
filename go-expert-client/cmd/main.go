package main

import (
	"log"
	"net/http"

	controller "github.com/FabioWGermano/go-expert-client/internal/controller/action"
)

func main() {
	http.HandleFunc("/", controller.Handle)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
