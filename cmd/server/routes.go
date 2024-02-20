package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/iriojose/cada_test/cmd/server/handlers"
	"github.com/iriojose/cada_test/pkg/utils"
)

func NewServer() {
	PORT := os.Getenv("PORT")

	r := mux.NewRouter()

	r.HandleFunc("/exchange-rate", utils.MakeHttpHandlerFunc(handlers.ExchangeRateHandler)).Methods(http.MethodPost)
	r.HandleFunc("/persons", utils.MakeHttpHandlerFunc(handlers.GetPersonsHandler)).Methods(http.MethodGet)

	http.ListenAndServe(PORT, r)
	fmt.Println("Server running on port", PORT)
}
