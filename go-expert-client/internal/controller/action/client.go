package controller

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/FabioWGermano/go-expert-client/internal/controller/response"
	"github.com/FabioWGermano/go-expert-client/internal/model" // Add import for io package
)

var (
	timeout = 300 * time.Millisecond
)

func Handle(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), timeout)
	defer cancel()

	var c model.Cambio
	if err := c.NewCambioServer(ctx); err != nil {
		log.Printf("Error getting exchange rate: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := c.SaveCotacao(); err != nil {
		log.Printf("Error setting exchange rate in file: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.NewSucess(c, http.StatusOK).Send(w)
}
