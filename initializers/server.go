package initializers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/iriojose/cada_test/handlers"
	"github.com/iriojose/cada_test/utils"
)

func NewServer(port string) {
	PORT := os.Getenv("PORT")

	r := mux.NewRouter()

	r.HandleFunc("/exchange-rate", utils.MakeHttpHandlerFunc(handlers.ExchangeRateHandler)).Methods("POST")
	r.HandleFunc("/persons", utils.MakeHttpHandlerFunc(handlers.GetPersonsHandler)).Methods("POST")

	http.ListenAndServe(PORT, r)
	fmt.Println("Server running on port", PORT)
}
