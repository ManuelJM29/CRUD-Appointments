package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConncetDB() (*mongo.Client, error) {

	// Configuracion de las opciones de conexion
	clientOptions := options.Client().ApplyURI("mongodb+srv://manueljimenezm:<kbDC1fdhIGRNH5E9>@crud-appointments-db.wtbadxo.mongodb.net/?retryWrites=true&w=majority")

	//Conexion al servidor MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, clientOptions)
	defer cancel()

	if err != nil {
		return nil, err
	}

	// Comprobacion de la conexxion
	err = client.Ping(context.Background(), nil)

	if err != nil {
		return nil, err
	}

	return client, nil

}
