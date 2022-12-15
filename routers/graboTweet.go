package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/CarlosPC402/TwittPC/bd"
	"github.com/CarlosPC402/TwittPC/models"
)

/* GraboTweet permite grabar el tweet en la bd */
func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet

	err := json.NewDecoder(r.Body).Decode(&mensaje)

	/* Para que no de error de que no se está utilizando
	Quitar si da error porque en el video esto no lo hace */
	if err != nil {
		return
	}

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)

	/* se relizar aquí en el endpoint y no en la rutina e bd porque la rutina de bd no tiene el
	response writer */
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar el tweet, intente nuevamente"+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado insertar el tweet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
