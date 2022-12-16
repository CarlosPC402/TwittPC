package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/CarlosPC402/TwittPC/bd"
)

/* ObtenerAvatar envia el Avatar al HTTP */
func ObtenerAvatar(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro id!", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Usuario no encontrado!", http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open("uploads/avatars" + perfil.Avatar)
	if err != nil {
		http.Error(w, "Imágen no encontrada!", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Error al copiar la imágen", http.StatusBadRequest)
		/* Dice que aquí no es necesario porque es el final del proceso */
		//return
	}
}
