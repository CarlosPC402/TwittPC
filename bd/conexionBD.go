package bd

/* Si inicia en mayúscula es público */

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Pública
/* MongoCN es el objeto de conexión a la BD */
var MongoCN = ConectarBD()

// Privada
var clientOptions = options.Client().ApplyURI("mongodb+srv://dbCursoAngular:pCVIpN8lXC3h0e8v@cluster0.1bsrn.mongodb.net/test?authMechanism=SCRAM-SHA-1")

/* ConectarBD Permite conectar la BD */
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexión exitosa a la BD")
	return client
}

/* ChequeoConnection permite checar la conexión */
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}

	return 1
}
