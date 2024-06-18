package controller

import (
	"log"
	"net/http"

	"github.com/FabioWGermano/go-expert-server/internal/controller/response"
	"github.com/FabioWGermano/go-expert-server/internal/model"

	_ "github.com/mattn/go-sqlite3"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	var c model.Cambio
	if err := c.NewTaxaCambio(r.Context()); err != nil {
		log.Printf("Error getting exchange rate on server: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := c.SalvarCotacao(r.Context()); err != nil {
		log.Printf("Error setting exchange rate on db: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.NewSucess(c, http.StatusAccepted).Send(w)
}
