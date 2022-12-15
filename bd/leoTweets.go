package bd

import (
	"context"
	"log"
	"time"

	"github.com/CarlosPC402/TwittPC/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* LeoTweets  lee los tweets de un perfil*/
func LeoTweets(ID string, pagina int64) ([]*models.DevuelvoTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := MongoCN.Database("twittpc")
	col := db.Collection("tweet")

	var resultados []*models.DevuelvoTweets

	condicion := bson.M{
		"userid": ID,
	}

	opciones := options.Find()
	opciones.SetLimit(20)
	/* Ordenados descendente -1 */
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}})
	/* la primera vez, pagina vale 1, entonces: 1 - 1 = 0 * 20 = 0 -> para la primera vez no se salta
	ningun registro, para la sig pag, salta 20 y así sucesivamente */
	opciones.SetSkip((pagina - 1) * 20)

	/* Un cursor es un puntero donde se graban los resultados y se puede ir recorriendo de a 1 a la vez
	porque por cada registro que se encuentre se debe ir armando el resultado final */
	cursor, err := col.Find(ctx, condicion, opciones)

	if err != nil {
		log.Fatal(err.Error())
		return resultados, false
	}

	/* Se crea un contexto vacío con TODO() */
	for cursor.Next(context.TODO()) {
		var registro models.DevuelvoTweets
		err := cursor.Decode(&registro)
		if err != nil {
			return resultados, false
		}

		resultados = append(resultados, &registro)
	}

	return resultados, true
}
