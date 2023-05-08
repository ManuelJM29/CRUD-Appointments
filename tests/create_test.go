package tests

import (
	"context"
	"crud-appointments/database"

	// "crud-appointments/models"
	"testing"
	// "time"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateAppointment(t *testing.T) {
	// inicializar base de datos y borrar registros previos
	client, err := database.ConnectDB()
	if err != nil {
		t.Fatal(err)
	}

	db := client.Database("crud-appointments-db")
	col := db.Collection("appointments")

	_, err = col.DeleteMany(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}

	// Creacion de una nueva cita)
	// newAppointment := &models.Appointment{
	// 	ID:          primitive.NewObjectID(),
	// 	Patient:     "Manuel Jimenez",
	// 	Description: "depression",
	// 	StartDate:   time.Now(),
	// 	UpdateDate:  time.Now(),
}
