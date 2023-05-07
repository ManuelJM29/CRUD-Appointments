package handlers

import (
	"crud-appointments/database"
	"crud-appointments/models"
	"encoding/json"
	"net/http"
)

func CreateAppointment(w http.ResponseWriter, r *http.Request) {
	// Leer el cuerpo de la solicitud y decodificarlo en una estructura Appointment
	var appointment models.Appointment
	err := json.NewDecoder(r.Body).Decode(&appointment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insertar el nuevo appointment en la base de datos
	err = database.NewAppointment(&appointment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Devolver una respuesta con el appointment creado
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(appointment)
}
