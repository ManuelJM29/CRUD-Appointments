package handlers

import (
	"crud-appointments/database"
	"crud-appointments/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Inicio")
}

func CreateAppointment(w http.ResponseWriter, r *http.Request) {
	// Leer el cuerpo de la solicitud y decodificarlo en una estructura Appointment
	var appointment models.Appointment
	err := json.NewDecoder(r.Body).Decode(&appointment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insertar el nuevo appointment en la base de datos
	appointment.ID = primitive.NewObjectID()
	err = database.NewAppointment(appointment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Devolver una respuesta con el appointment creado
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(appointment)
}

func GetAppointments(w http.ResponseWriter, r *http.Request) {
	// Obtener todas las citas de la base de datos
	appointments, err := database.GetAppointments()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Devolver una respuesta con las citas obtenidas
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(appointments)
}

func GetAppointmentByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	appointment, err := database.GetAppointmentByID(params["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(appointment)
}

func UpdateAppointment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var appointment models.Appointment
	_ = json.NewDecoder(r.Body).Decode(&appointment)
	appointment.ID, _ = primitive.ObjectIDFromHex(id)
	result, err := database.UpdateAppointment(appointment)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if result.ModifiedCount == 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	updatedAppointment, _ := database.GetAppointmentByID(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedAppointment)
}

func DeleteAppointment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	// Eliminamos el appointment con el id correspondiente
	err := database.DeleteAppointmentByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Appointment deleted successfully"})
}
