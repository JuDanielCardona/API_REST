package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"taller_docker/database"
	"taller_docker/models"
	"taller_docker/security"

	"github.com/gorilla/mux"
)

func GetAllUsers_handler(w http.ResponseWriter, r *http.Request) {

	if !(security.IsValidToken(r)) {
		http.Error(w, "Error: Invalid token", http.StatusUnauthorized)
		return
	}

	var users []models.User
	users, err := database.GetUsers()
	if err != nil {
		http.Error(w, "Error: Failed to get users", http.StatusInternalServerError)
		return
	}

	// Formatear la información de los usuarios en un formato más legible
	formattedJSON, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		http.Error(w, "Error: Failed to format JSON", http.StatusInternalServerError)
		return
	}

	// Establecer el encabezado Content-Type y el código de estado
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "\nOK: All users search successfully.\n")
	w.Write(formattedJSON)

}

func GetUserById_handler(w http.ResponseWriter, r *http.Request) {
	if !security.IsValidToken(r) {
		http.Error(w, "Error: Invalid token", http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	var user *models.User
	user, _ = database.GetUserById(params["id"])

	if user == nil {
		http.Error(w, "Error: User not found", http.StatusNotFound)
		return
	}

	// Formatear la información del usuario en un formato más legible
	formattedJSON, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		http.Error(w, "Error: Failed to format JSON", http.StatusInternalServerError)
		return
	}

	// Establecer el encabezado Content-Type y el código de estado
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "\nOK: User search successfully.\n")
	w.Write(formattedJSON)
}

func AddUser_handler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error: Invalid JSON", http.StatusBadRequest)
		return
	}

	createdUser, err := database.AddUser(user)
	if err != nil {
		http.Error(w, "Error: User could not be created", http.StatusBadRequest)
		return
	}

	// Formatear la información del usuario creado en un formato más legible
	formattedJSON, err := json.MarshalIndent(createdUser, "", "  ")
	if err != nil {
		http.Error(w, "Error: Failed to format JSON", http.StatusInternalServerError)
		return
	}

	// Establecer el encabezado Content-Type y el código de estado
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	fmt.Fprintf(w, "\nOK: User created successfully.\n")
	w.Write(formattedJSON)
}

func DeleteUser_handler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error: Problem to covert JSON", http.StatusBadRequest)
		return
	}

	id := user.Id
	if !security.IsValidToken(r) {
		http.Error(w, "Error: Invalid token", http.StatusUnauthorized)
		return
	}

	err = database.DeleteUser(id)
	if err != nil {
		http.Error(w, "Error: User not found to delete", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "\nOK: User with id(%d) was deleted.\n", id)
}

func UpdateUser_handler(w http.ResponseWriter, r *http.Request) {
	// Obtener el ID del cliente de la URL
	vars := mux.Vars(r)
	urlID := vars["id"]
	fmt.Println("ID recibido en la URL:", urlID)

	// Decodificar el JSON del cuerpo de la solicitud
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error: Invalid JSON", http.StatusBadRequest)
		return
	}

	// Verificar si el ID del JSON es diferente al ID de la URL
	if strconv.Itoa(user.Id) != urlID {
		http.Error(w, "Error: ID in JSON does not match ID in URL", http.StatusBadRequest)
		return
	}

	// Convertir el ID a entero
	userID, err := strconv.Atoi(urlID)
	if err != nil {
		http.Error(w, "Error: Invalid user ID", http.StatusBadRequest)
		return
	}

	if !security.IsValidToken(r) {
		http.Error(w, "Error: Invalid token", http.StatusUnauthorized)
		return
	}

	// Llamar a la función de actualización del usuario
	err = database.UpdateUser(userID, user)
	if err != nil {
		http.Error(w, "Error: User not found to update", http.StatusNotFound)
		return
	}

	formattedJSON, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		http.Error(w, "Error: Failed to format JSON", http.StatusInternalServerError)
		return
	}

	// Establecer el encabezado Content-Type y el código de estado
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	fmt.Fprintf(w, "\nOK: User updated successfully.\n")
	w.Write(formattedJSON)
}
