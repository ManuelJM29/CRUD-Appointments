package main

import (
	"crud-appointments/database"
	"log"
)

func main() {

	//Conexion a la base de datos
	client, err := database.ConncetDB()

	if err != nil {
		log.Fatal(err)
	}
}
