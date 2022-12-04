package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/CarlosPC402/TwittPC/middlew"
	"github.com/CarlosPC402/TwittPC/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* Manejadores seteo mi puerto, handler y pongo a escuchar el servidor */
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routes.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
