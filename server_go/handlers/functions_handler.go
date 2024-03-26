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

	if !security.IsValidToken(r, "") {
		http.Error(w, "Error: Invalid token", http.StatusUnauthorized)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	users, err := database.GetUsers(page, pageSize)
	if err != nil {
		http.Error(w, "Error: Failed to get users", http.StatusInternalServerError)
		return
	}

	formattedJSON, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		http.Error(w, "Error: Failed to format JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "\nOK: All users search successfully.\n")
	w.Write(formattedJSON)
}

func GetUserById_handler(w http.ResponseWriter, r *http.Request) {
	if !security.IsValidToken(r, "") {
		http.Error(w, "Error: Invalid token", http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	user, err := database.GetUserById(params["id"])
	if err != nil {
		http.Error(w, "Error: Failed to search user id ("+params["id"]+")", http.StatusInternalServerError)
		return
	}

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

	// Establecer el encabezado Content-Type y enviar la respuesta
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

	if user.Name == "" || user.Password == "" || user.Email == "" {
		http.Error(w, "Error: User, Password & Email info are obligatory", http.StatusBadRequest)
		return
	}

	createdUser, err := database.AddUser(user)
	if err != nil {
		http.Error(w, "Error: This email is already in use", http.StatusBadRequest)
		return
	}

	tokenString := security.GenerateToken(createdUser)
	if tokenString == "" {
		http.Error(w, "Error: Failed to sign token", http.StatusInternalServerError)
		return
	}

	formattedJSON, err := json.MarshalIndent(createdUser, "", "  ")
	if err != nil {
		http.Error(w, "Error: Failed to format JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	fmt.Fprintf(w, "\nOK: User created successfully.\n")
	w.Write(formattedJSON)
	fmt.Fprintf(w, "\nToken:\n")
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, tokenString)
}

func DeleteUser_handler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error: Problem to covert JSON", http.StatusBadRequest)
		return
	}

	if !security.IsValidToken(r, strconv.Itoa(user.Id)) {
		http.Error(w, "Error: Invalid token", http.StatusUnauthorized)
		return
	}

	err = database.DeleteUser(user)
	if err != nil {
		http.Error(w, "Error: Failed to update user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "\nOK: User with id(%d) was deleted.\n", user.Id)
}

func UpdateUser_handler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error: Problem to covert JSON", http.StatusBadRequest)
		return
	}

	if !security.IsValidToken(r, strconv.Itoa(user.Id)) {
		http.Error(w, "Error: Invalid token", http.StatusUnauthorized)
		return
	}

	updatedUser, err := database.UpdateUser(user)
	if err != nil {
		http.Error(w, "Error: Failed to update user", http.StatusInternalServerError)
		return
	}

	formattedJSON, err := json.MarshalIndent(updatedUser, "", "  ")
	if err != nil {
		http.Error(w, "Error: Failed to format JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(formattedJSON)
}

func RecoverPassword_handler(w http.ResponseWriter, r *http.Request) {

	info := mux.Vars(r)
	email := info["email"]

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error: Invalid JSON", http.StatusBadRequest)
		return
	}

	if user.Email != email {
		http.Error(w, "Error: Requester does not match the database (Email)", http.StatusBadRequest)
		return
	}

	if user.Name != "" {
		userFound, err := database.GetUserByEmail(user.Email)
		if err != nil {
			http.Error(w, "Error: Email not found", http.StatusNotFound)
			return
		}
		if userFound.Name != user.Name {
			http.Error(w, "Error: Requester does not match the database (Name)", http.StatusBadRequest)
			return
		}
	}

	password, err := database.RecoverPassword(info["email"])
	if err != nil {
		http.Error(w, "Error: User not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "\nOK: Password was recuperated.\nPassword is: ")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(password))
	fmt.Fprintf(w, "\n")
}

func UpdatePassword_handler(w http.ResponseWriter, r *http.Request) {

	info := mux.Vars(r)
	email := info["email"]

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error: Invalid JSON", http.StatusBadRequest)
		return
	}
	id := user.Name
	if !security.IsValidToken(r, id) {
		http.Error(w, "Error: Invalid token", http.StatusUnauthorized)
		return
	}

	newPassword, err := database.UpdatePassword(user, email)
	if err != nil {
		if err.Error() == "Error: New password must be different from the current one" {
			http.Error(w, "Error: New password must be different from the current one", http.StatusBadRequest)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "\nOK: Password was Updated.\nNew password is: ")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(newPassword))
	fmt.Fprintf(w, "\n")
}
