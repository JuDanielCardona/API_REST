package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"taller_docker/database"
	"taller_docker/models"
	"taller_docker/security"
)

func Login_handler(w http.ResponseWriter, r *http.Request) {

	// Verificar si la solicitud es de tipo POST
	if r.Method != http.MethodPost {
		http.Error(w, "Error: Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//decode user from json
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	_, err := database.SearchUser(&user)

	// Verificar si se proporcionaron usuario y clave
	if user.Name == "" || user.Password == "" {
		http.Error(w, "Error: User & Password info are obligatory", http.StatusBadRequest)
		return
	}

	tokenString := security.GenerateToken(&user)
	if tokenString == "" {
		http.Error(w, "Error: Failed to sign token", http.StatusInternalServerError)
		return
	}

	if err != nil {
		http.Error(w, "Error: User not found", http.StatusNotFound)
		return
	}

	// Responder con el token JWT
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, "Token: \n")
	fmt.Fprint(w, tokenString)
}
