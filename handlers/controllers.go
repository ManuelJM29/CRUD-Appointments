package handlers

import (
	"crud-appointments/database"
	"crud-appointments/models"
	"crud-appointments/responses"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}

func CreateAppointment(w http.ResponseWriter, r *http.Request) {
	// Leer el cuerpo de la solicitud y decodificarlo en una estructura Appointment
	var appointment models.Appointment
	err := json.NewDecoder(r.Body).Decode(&appointment)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := responses.Response{
			Msg:    err.Error(),
			Status: http.StatusOK}
		json.NewEncoder(w).Encode(response)
		return
	}
	appointment.ID = primitive.NewObjectID()

	// validate
	validate := validator.New()
	validateErr := validate.Struct(appointment)
	if validateErr != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := responses.Response{
			Msg:    "Campos invalidos",
			Status: http.StatusBadRequest}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Insertar el nuevo appointment en la base de datos
	err = database.NewAppointment(appointment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Devolver una respuesta con el appointment creado
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := responses.Response{
		Data:   appointment,
		Status: http.StatusCreated,
		Msg:    "Cita creada"}
	json.NewEncoder(w).Encode(response)
}

func GetAppointments(w http.ResponseWriter, r *http.Request) {
	// Obtener todas las citas de la base de datos
	appointments, err := database.GetAppointments()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := responses.Response{
			Msg:    err.Error(),
			Status: http.StatusOK}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Devolver una respuesta con las citas obtenidas
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := responses.Response{
		Data:   appointments,
		Status: http.StatusOK,
		Msg:    "Citas encontradas"}
	json.NewEncoder(w).Encode(response)
}

func GetAppointmentByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	appointment, err := database.GetAppointmentByID(params["id"])

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := responses.Response{
			Msg:    "No existe cita",
			Status: http.StatusOK}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := responses.Response{
		Data:   appointment,
		Status: http.StatusOK,
		Msg:    "Cita encontrada"}
	json.NewEncoder(w).Encode(response)
}

func UpdateAppointment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var appointment models.Appointment
	_ = json.NewDecoder(r.Body).Decode(&appointment)
	appointment.ID, _ = primitive.ObjectIDFromHex(id)

	// validate
	validate := validator.New()
	validateErr := validate.Struct(appointment)
	if validateErr != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := responses.Response{
			Msg:    "Campos invalidos",
			Status: http.StatusBadRequest}
		json.NewEncoder(w).Encode(response)
		return
	}

	result, err := database.UpdateAppointment(appointment)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := responses.Response{
			Msg:    err.Error(),
			Status: http.StatusOK}
		json.NewEncoder(w).Encode(response)
		return
	}
	if result.ModifiedCount == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		response := responses.Response{
			Msg:    "Cita no modificada",
			Status: http.StatusNotFound}
		json.NewEncoder(w).Encode(response)
		return
	}

	updatedAppointment, _ := database.GetAppointmentByID(id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := responses.Response{
		Data:   updatedAppointment,
		Status: http.StatusOK,
		Msg:    "Cita actualizada"}
	json.NewEncoder(w).Encode(response)
}

func DeleteAppointment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := database.DeleteAppointmentByID(params["id"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := responses.Response{
			Msg:    "Cita no encontrada",
			Status: http.StatusOK}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := responses.Response{
		Status: http.StatusOK,
		Msg:    "Cita eliminada"}
	json.NewEncoder(w).Encode(response)
}
