package bd

import "golang.org/x/crypto/bcrypt"

/* EncriptarPassword es la rutina que permite encriptar la contraseña */
func EncriptarPassword(pass string) (string, error) {
	/* El costo es la cantidad a la que se va elevar el algoritmo en base binario (2), es decir en este
	caso sería 2 elevado a la potencia 8, así que entre más alto la cantidad de pasadas que se le va dar al
	password para encriptar, entre más alto más segura la contraseña pero también se demora más.
	La recomendación es: usuario normales un costo de 6 y para admin 8 */
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
