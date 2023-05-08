package main

import (
	"crud-appointments/database"
	"crud-appointments/routers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// Conexion a la base de datos
	_, err := database.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}

	// Crear un nuevo enrutador
	router := mux.NewRouter()
	routers.SetAppointmentRoutes(router)

	// Escuchar en el puerto 8080
	log.Fatal(http.ListenAndServe(":3000", router))
	fmt.Println("Servidor iniciado en el puerto 3000")
}
