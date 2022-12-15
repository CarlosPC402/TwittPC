package bd

import (
	"context"
	"time"

	"github.com/CarlosPC402/TwittPC/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* InsertoTweet graba el mensaje en la BD */
func InsertoTweet(t models.GraboTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := MongoCN.Database("twittpc")
	col := db.Collection("tweet")

	registro := bson.M{
		"userid":  t.UserID,
		"mensaje": t.Mensaje,
		"fecha":   t.Fecha,
	}

	/* InsertOne devuleve dos valores, resultado y error(si no hay error devuelve nil) */
	result, err := col.InsertOne(ctx, registro)

	if err != nil {
		/* string("") en caso de que falle con el de abajo poner este */
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)

	/* También se podría hacer con .Hex() */
	return objID.String(), true, nil
}
