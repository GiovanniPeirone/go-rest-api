package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Estructura de ejemplo para enviar datos en la respuesta
type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Passworld string `json:"passworld`
}

func main() {
	// Asignar la ruta y el manejador
	http.HandleFunc("/users", getUsersHandler)

	// Iniciar el servidor
	fmt.Println("Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// Manejador para la ruta GET /users
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Verificar que el método sea GET
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Crear datos de ejemplo
	users := []User{
		{ID: 1, Name: "Juan", Email: "juan@example.com", Passworld: "pepito3"},
	}

	// Configurar el encabezado de respuesta
	w.Header().Set("Content-Type", "application/json")

	// Convertir los datos a JSON y enviarlos como respuesta
	json.NewEncoder(w).Encode(users)
}
