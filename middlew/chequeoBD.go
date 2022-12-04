package middlew

import (
	"net/http"

	"github.com/CarlosPC402/TwittPC/bd"
)

/* ChequeoBD es el middleware que me permite concer el estado de la BD */
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	/* función anónima, no tiene nombre */
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexión perdida con BD", 500)
			return
		}

		next.ServeHTTP(w, r)
	}
}
