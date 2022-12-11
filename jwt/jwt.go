package jwt

import (
	"time"

	"github.com/CarlosPC402/TwittPC/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/* GeneroJWT genera el token encriptado con JWT */
func GeneroJWT(t models.Usuario) (string, error) {
	miClave := []byte("MastersdelDesarrollo_grupodeFacebook")

	/* No se pasa directamete el modelo del usuario porque en esa se encuentra el token */
	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"ubicacion":        "t.Ubicacion",
		"sitioWeb":         "t.SitioWeb",
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
