package bd

import (
	"context"
	"time"

	"github.com/CarlosPC402/TwittPC/models"
	"go.mongodb.org/mongo-driver/bson"
)

/* ChequeoYaExisteUsuario recibe como parámetro un email y valida su existencia en BD */
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittpc")
	col := db.Collection("usuarios")

	condicion := bson.M{"email": email}

	var resultado models.Usuario

	/* Si hay error lo graba en la variable err y si encuentra resultados lo decodifica y graba en el resultado */
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	/* Se convierte el objectid en string hexadecimal */
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}

	return resultado, true, ID
}
