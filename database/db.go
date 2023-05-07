package database

import (
	"context"
	"crud-appointments/models"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() (*mongo.Client, error) {

	// Configuracion de las opciones de conexion
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

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

func NewAppointment(appointment *models.Appointment) error {
	client, err := ConnectDB()
	if err != nil {
		return err
	}

	collection := client.Database("crud-appointments-db").Collection("appointments")

	fmt.Printf("Appointment before insert: %+v\n", appointment)

	_, err = collection.InsertOne(context.Background(), appointment)
	if err != nil {
		return err
	}

	return nil
}
