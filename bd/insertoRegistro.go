package bd

import (
	"context"
	"time"

	"github.com/CarlosPC402/TwittPC/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* Función devuelve datos, método no */
/* InsertoRegistro es la parada final con la BD para insertar los datos del usuario */
func InsertoRegistro(u models.Usuario) (string, bool, error) {
	/* es el contexto que se viene trayendo desde la conexión con la BD con TODO */
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	/* Se ejecuta al último, aquí se cancela ese contexto de withtime */
	defer cancel()

	db := MongoCN.Database("twittpc")
	col := db.Collection("usuarios")

	u.Password, _ = EncriptarPassword(u.Password)

	/* InsertOne devuelve el id insertado en mongo y devuelve si hubo error */
	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	/* InsertedId devuelve un objectID entonces se convierte  a string aquí */
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
