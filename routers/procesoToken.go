package routers

import (
	"errors"
	"strings"

	"github.com/CarlosPC402/TwittPC/bd"
	"github.com/CarlosPC402/TwittPC/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/* Email Valor de email usado en todos los EndPoints */
var Email string

/* IDUsuario es el ID devuelto del modelo que se usará en todos los endpoints */
var IDUsuario string

/* ProcesoToken proceso token para extraer su valores*/
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	/* si una función devuelve error, El error se debe poner al final porque puede dar errores
	en algunas versiones*/
	miClave := []byte("MastersdelDesarrollo_grupodeFacebook")
	/* Tiene que ser un puntero */
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)

		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}

		return claims, encontrado, IDUsuario, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	return claims, false, string(""), err
}
