package routers

import (
	"crud-appointments/handlers"

	"github.com/gorilla/mux"
)

func SetAppointmentRoutes(router *mux.Router) {
	router.HandleFunc("/", handlers.Homepage).Methods("GET")
	router.HandleFunc("/appointments", handlers.CreateAppointment).Methods("POST")
	router.HandleFunc("/appointments", handlers.GetAppointments).Methods("GET")
	router.HandleFunc("/appointments/{id}", handlers.GetAppointmentByID).Methods("GET")
	router.HandleFunc("/appointments/{id}", handlers.UpdateAppointment).Methods("PUT")
	router.HandleFunc("/appointments/{id}", handlers.DeleteAppointment).Methods("DELETE")
}
