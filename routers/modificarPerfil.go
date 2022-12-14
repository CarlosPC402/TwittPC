package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CarlosPC402/TwittPC/bd"
	"github.com/CarlosPC402/TwittPC/models"
)

/* ModificarPerfil modifica el perfil dl usuario*/
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	/* Body viene vacío o formato json incorrecto por ej */
	if err != nil {
		http.Error(w, "Datos Incorrectos"+err.Error(), 400)
		return
	}

	var status bool
	/* IDUsuario es la variable global grabado en procesoToken*/
	fmt.Println(t)
	fmt.Println(IDUsuario)
	status, err = bd.ModificoRegistro(t, IDUsuario)

	if err != nil {
		http.Error(w, "Ocurrión un error al intentar modificar el registro, reintente nuevamente"+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado modificar el registro", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
