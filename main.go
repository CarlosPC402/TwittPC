package main

import (
	"log"

	"github.com/CarlosPC402/TwittPC/bd"
	"github.com/CarlosPC402/TwittPC/handlers" //Lo toma de la ruta de directorios no de github como tal
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a BD")
		return
	}
	handlers.Manejadores()
}
