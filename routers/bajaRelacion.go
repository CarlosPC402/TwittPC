package routers

import (
	"net/http"

	"github.com/CarlosPC402/TwittPC/bd"
	"github.com/CarlosPC402/TwittPC/models"
)

/* BajaRelacion realiza el borrado de la relación entre usuarios */
func BajaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.BorroRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrión un error al intentar borrar relación"+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado borrar la relación"+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
