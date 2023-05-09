# CRUD Appointments

CRUD de citas con Go y MongoDB.

## Instalación

*Go 
*MongoDB

1. Clonar el repositorio: `git clone https://github.com/ManuelJM29/CRUD-Appointments.git`
2. Instalar las dependencias: `go get ./...`
3. Iniciar el servidor: `go run main.go`
4. Importar la coleccion de Postman que se encuentra dentro de la carpeta `docs`

## Uso

Acceder a la API a través de http://localhost:3000
Base de datos corriendo en local, conexion "mongodb://localhost:27017"

### Endpoints

- `GET /healthcheck`: Verificar si el servidor está en línea
- `POST /appointments`: Crear una nueva cita
- `GET /appointments`: Obtener todas las citas
- `GET /appointments/{id}`: Obtener una cita por ID
- `PUT /appointments/{id}`: Actualizar una cita existente por ID
- `DELETE /appointments/{id}`: Eliminar una cita existente por ID
