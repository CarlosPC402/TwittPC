package routers

import (
	"encoding/json"
	"net/http"

	"github.com/CarlosPC402/TwittPC/bd"
	"github.com/CarlosPC402/TwittPC/models"
)

/* Registro es la función para crear en BD el registro de usuario */
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	/* El Body de un httpRequest es un stream es decir que solo se puede leer una sola vez y después se destruye */
	/* Después de usar r.Body aquí ya no se puede usar en ningún otro lado */
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}
	/* len devuelve la longitud del string */
	if len(t.Email) == 0 {
		http.Error(w, "El email es requerido ", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "La contraseña debe ser de al menos 6 caracteres ", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)

	if encontrado == true {
		http.Error(w, "Ya existe un usuario registrado con ese email ", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)

	if err != nil {
		http.Error(w, "Ocurrió un error en el registro "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del usuario ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
