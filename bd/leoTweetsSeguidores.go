package bd

import (
	"context"
	"time"

	"github.com/CarlosPC402/TwittPC/models"
	"go.mongodb.org/mongo-driver/bson"
)

/* LeoTweetsSeguidores lee el tweet de mis seguidores*/
func LeoTweetsSeguidores(ID string, pagina int) ([]models.DevuelvoTweetsSeguidores, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := MongoCN.Database("twittpc")
	col := db.Collection("relacion")

	skip := (pagina - 1) * 20

	condiciones := make([]bson.M, 0)
	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid": ID}})
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "usuariorelacionid",
			"foreignField": "userid",
			"as":           "tweet",
		}})
	/* el unwind evita que se devuelvan los datos en modo de mastro-detalle */
	condiciones = append(condiciones, bson.M{"$unwind": "$tweet"})

	condiciones = append(condiciones, bson.M{"$sort": bson.M{"tweet.fecha": -1}})
	condiciones = append(condiciones, bson.M{"$skip": skip})
	condiciones = append(condiciones, bson.M{"$limit": 20})

	/* Con el aggregate no se neceita recorrer el cursos */
	cursor, err := col.Aggregate(ctx, condiciones)
	var result []models.DevuelvoTweetsSeguidores
	/* Solo para que no marque error el siguiente if, en el video no lo incluye	 */
	if err != nil {
		return result, false
	}
	err = cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}

	return result, true
}
