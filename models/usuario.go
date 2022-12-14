package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* Usuario es el modelo de usuario de la BD mongo */
type Usuario struct {
	/* El primero(bson) corresponde a lo que está en la BD y el segundo(json) corresponde al valor devuelto en el nave
	omitempty indica que lo omita si no se encuentra, para el ID indica que si no se encuentra
	el valor en los datos enviados no genere error y los demás indica que si no se encuentra el
	valor en BD no devuelva nada en el json */
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre          string             `bson:"nombre" json:"nombre,omitempty"`
	Apellidos       string             `bson:"apellidos" json:"apellidos,omitempty"`
	FechaNacimiento time.Time          `bson:"fechaNacimiento" json:"fechaNacimiento,omitempty"`
	Email           string             `bson:"email" json:"email"`
	Password        string             `bson:"password" json:"password,omitempty"`
	Avatar          string             `bson:"avatar" json:"avatar,omitempty"`
	Banner          string             `bson:"banner" json:"banner,omitempty"`
	Biografia       string             `bson:"biografia" json:"biografia,omitempty"`
	Ubicacion       string             `bson:"ubicacion" json:"ubicacion,omitempty"`
	SitioWeb        string             `bson:"sitioWeb" json:"sitioWeb,omitempty"`
}
