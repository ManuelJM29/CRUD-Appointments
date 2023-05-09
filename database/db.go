package database

import (
	"context"
	"crud-appointments/models"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// Funcion para devolver una colección de la base de datos
func GetCollection(collectionName string) *mongo.Collection {
	client, err := ConnectDB()
	if err != nil {
		return nil
	}
	collection := client.Database("crud-appointments-db").Collection(collectionName)

	return collection
}

func NewAppointment(appointment models.Appointment) error {
	client, err := ConnectDB()
	if err != nil {
		return err
	}

	collection := client.Database("crud-appointments-db").Collection("appointments")
	_, err = collection.InsertOne(context.Background(), appointment)
	if err != nil {
		return err
	}
	return nil
}

func GetAppointments() ([]models.Appointment, error) {
	// Obtener una conexión a la base de datos
	collection := GetCollection("appointments")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Obtener todas las citas de la base de datos
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	// Declarar una variable para almacenar las citas
	var appointments []models.Appointment

	// Iterar sobre el cursor y agregar cada cita a la lista
	for cursor.Next(context.Background()) {
		var appointment models.Appointment
		if err := cursor.Decode(&appointment); err != nil {
			return nil, err
		}
		appointments = append(appointments, appointment)
	}
	return appointments, nil
}

func GetAppointmentByID(id string) (*models.Appointment, error) {
	objectID, _ := primitive.ObjectIDFromHex(id)
	collection := GetCollection("appointments")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var appointment models.Appointment
	err := collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&appointment)
	if err != nil {
		return nil, err
	}
	return &appointment, nil
}

func UpdateAppointment(appointment models.Appointment) (*mongo.UpdateResult, error) {
	collection := GetCollection("appointments")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": appointment.ID}

	update := bson.M{
		"$set": bson.M{
			"patient":     appointment.Patient,
			"description": appointment.Description,
			"start_date":  appointment.StartDate,
			"update_date": appointment.UpdateDate,
		},
	}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func DeleteAppointmentByID(id string) error {
	objectID, _ := primitive.ObjectIDFromHex(id)
	collection := GetCollection("appointments")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("No se encontró la cita con ID %s", id)
	}
	return nil
}
