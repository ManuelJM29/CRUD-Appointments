package tests

import (
	"bytes"
	"context"
	"crud-appointments/database"
	"crud-appointments/handlers"
	"crud-appointments/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
	newAppointment := &models.Appointment{
		ID:          primitive.NewObjectID(),
		Patient:     "Manuel Jimenez",
		Description: "depresion",
		StartDate:   time.Now(),
		UpdateDate:  time.Now(),
	}

	// Codificar el nuevo appointment como JSON
	newAppointmentJSON, err := json.Marshal(newAppointment)
	if err != nil {
		t.Fatalf("Error al codificar el appointment como JSON: %s", err)
	}

	// Crear una nueva solicitud HTTP POST con el cuerpo del nuevo appointment JSON
	req, err := http.NewRequest("POST", "/appointments", bytes.NewBuffer(newAppointmentJSON))
	if err != nil {
		t.Fatalf("Error al crear la solicitud HTTP: %s", err)
	}

	// Crear un ResponseRecorder (implementa http.ResponseWriter) para capturar la respuesta HTTP
	rr := httptest.NewRecorder()

	// Crear un manejador de HTTP para la ruta "/appointments" y el método HTTP POST
	handler := http.HandlerFunc(handlers.CreateAppointment)

	// Hacer la solicitud HTTP
	handler.ServeHTTP(rr, req)

	// Comprobar que el estado de la respuesta es correcto (StatusCreated)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Estado de respuesta incorrecto: se esperaba %v pero se obtuvo %v", http.StatusCreated, status)
	}

	// Comprobar que el tipo de contenido de la respuesta es JSON
	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("Tipo de contenido incorrecto: se esperaba %v pero se obtuvo %v", expectedContentType, contentType)
	}

	// Decodificar la respuesta como JSON y obtener el appointment creado
	var createdAppointment models.Appointment
	err = json.NewDecoder(rr.Body).Decode(&createdAppointment)
	if err != nil {
		t.Fatalf("Error al decodificar la respuesta como JSON: %s", err)
	}

	// Comprobar que el appointment creado es igual al appointment enviado en la solicitud
	if !reflect.DeepEqual(createdAppointment, newAppointment) {
		t.Errorf("El appointment creado es incorrecto: se esperaba %v pero se obtuvo %v", newAppointment, createdAppointment)
	}
}

// err = CreateAppointment(context.Background(), col, newAppoinment)
// if err != nil {
// 	t.Fatal(err)
// }

// // Comprobar que la cita se haya creado correctamente
// cursor, err := col.Find(context.Background(), nil)
// if err != nil {
// 	t.Fatal(err)
// }
// defer cursor.Close(context.Background())
// if !cursor.Next(context.Background()) {
// 	t.Fatal("No se encontraron registros")
// }
// var result models.Appointment
// err = cursor.Decode(&result)
// if err != nil {
// 	t.Fatal(err)
// }
// if result.ID == "" {
// 	t.Fatal("El ID de la cita no puede ser vacío")
// }
